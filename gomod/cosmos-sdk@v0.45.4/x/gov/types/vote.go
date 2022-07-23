package types

import (
	"encoding/json"
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
)



func NewVote(proposalID uint64, voter sdk.AccAddress, options WeightedVoteOptions) Vote {
	return Vote{ProposalId: proposalID, Voter: voter.String(), Options: options}
}

func (v Vote) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}


type Votes []Vote


func (v Votes) Equal(other Votes) bool {
	if len(v) != len(other) {
		return false
	}

	for i, vote := range v {
		if vote.String() != other[i].String() {
			return false
		}
	}

	return true
}

func (v Votes) String() string {
	if len(v) == 0 {
		return "[]"
	}
	out := fmt.Sprintf("Votes for Proposal %d:", v[0].ProposalId)
	for _, vot := range v {
		out += fmt.Sprintf("\n  %s: %s", vot.Voter, vot.Options)
	}
	return out
}


func (v Vote) Empty() bool {
	return v.String() == Vote{}.String()
}


func NewNonSplitVoteOption(option VoteOption) WeightedVoteOptions {
	return WeightedVoteOptions{{option, sdk.NewDec(1)}}
}

func (v WeightedVoteOption) String() string {
	out, _ := json.Marshal(v)
	return string(out)
}


type WeightedVoteOptions []WeightedVoteOption

func (v WeightedVoteOptions) String() (out string) {
	for _, opt := range v {
		out += opt.String() + "\n"
	}

	return strings.TrimSpace(out)
}


func ValidWeightedVoteOption(option WeightedVoteOption) bool {
	if !option.Weight.IsPositive() || option.Weight.GT(sdk.NewDec(1)) {
		return false
	}
	return ValidVoteOption(option.Option)
}



func VoteOptionFromString(str string) (VoteOption, error) {
	option, ok := VoteOption_value[str]
	if !ok {
		return OptionEmpty, fmt.Errorf("'%s' is not a valid vote option, available options: yes/no/no_with_veto/abstain", str)
	}
	return VoteOption(option), nil
}



func WeightedVoteOptionsFromString(str string) (WeightedVoteOptions, error) {
	options := WeightedVoteOptions{}
	for _, option := range strings.Split(str, ",") {
		fields := strings.Split(option, "=")
		option, err := VoteOptionFromString(fields[0])
		if err != nil {
			return options, err
		}
		if len(fields) < 2 {
			return options, fmt.Errorf("weight field does not exist for %s option", fields[0])
		}
		weight, err := sdk.NewDecFromStr(fields[1])
		if err != nil {
			return options, err
		}
		options = append(options, WeightedVoteOption{option, weight})
	}
	return options, nil
}


func ValidVoteOption(option VoteOption) bool {
	if option == OptionYes ||
		option == OptionAbstain ||
		option == OptionNo ||
		option == OptionNoWithVeto {
		return true
	}
	return false
}


func (vo VoteOption) Marshal() ([]byte, error) {
	return []byte{byte(vo)}, nil
}


func (vo *VoteOption) Unmarshal(data []byte) error {
	*vo = VoteOption(data[0])
	return nil
}


func (vo VoteOption) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(vo.String()))
	default:
		s.Write([]byte(fmt.Sprintf("%v", byte(vo))))
	}
}
