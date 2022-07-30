package client

import (
	"context"
	"errors"
	"freemasonry.cc/blockchain/core"
	"freemasonry.cc/blockchain/util"
	"freemasonry.cc/blockchain/x/comm/client/rest"
	"freemasonry.cc/blockchain/x/comm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/sirupsen/logrus"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"strconv"
	"strings"
)

type GatewayClient struct {
	ServerUrl string
	logPrefix string
}


func (this *GatewayClient) StatusInfo() (statusInfo *ctypes.ResultStatus, err error) {
	node, err := clientCtx.GetNode()
	return node.Status(context.Background())
}


func (this *GatewayClient) QueryGateway(gatewayAddress, gatewayNum string) (data *types.Gateway, notFound bool, err error) {
	log := util.BuildLog(util.GetStructFuncName(this), util.LmChainClient).WithFields(logrus.Fields{"gatewayAddress": gatewayAddress})
	params := types.QueryGatewayInfoParams{GatewayAddress: gatewayAddress, GatewayNumIndex: gatewayNum}
	bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		log.WithError(err).Error("MarshalJSON")
		return
	}
	resBytes, _, err := clientCtx.QueryWithData("custom/comm/"+types.QueryGatewayInfo, bz)
	if err != nil {
		
		if strings.Contains(err.Error(), types.ErrGatewayNumNotFound.Error()) {
			notFound = true
			err = nil
		} else {
			log.WithError(err).Error("QueryWithData")
		}
		return
	}
	if resBytes != nil {
		data = new(types.Gateway)
		err = util.Json.Unmarshal(resBytes, data)
		if err != nil {
			return
		}
	}
	return
}


func (this *GatewayClient) QueryGatewayList() (data []types.Gateway, err error) {
	log := util.BuildLog(util.GetStructFuncName(this), util.LmChainClient)
	resBytes, _, err := clientCtx.QueryWithData("custom/comm/"+types.QueryGatewayList, nil)
	if err != nil {
		log.WithError(err).Error("QueryWithData")
		return nil, err
	}
	if resBytes != nil {
		err := util.Json.Unmarshal(resBytes, &data)
		if err != nil {
			return nil, err
		}
	}
	return
}


func (this *GatewayClient) ValidatorInfo() (validatorInfo *types.ValidatorInfor, err error) {
	nodeStatus, err := this.StatusInfo()
	if err != nil {
		return nil, err
	}

	
	consAddress, err := sdk.ConsAddressFromHex(nodeStatus.ValidatorInfo.Address.String())
	if err != nil {
		return nil, err
	}
	validatorInfo = &types.ValidatorInfor{
		ValidatorStatus:   "", 
		ValidatorPubAddr:  nodeStatus.ValidatorInfo.PubKey.Address().String(),
		ValidatorConsAddr: consAddress.String(),
	}

	
	validator, notFound, err := this.FindValidatorByConsAddress(consAddress.String())
	if notFound {
		validatorInfo.ValidatorStatus = "4" 
		return validatorInfo, nil
	}
	if err != nil {
		return nil, err
	}

	validatorInfo.ValidatorStatus = strconv.Itoa(this.GetValidatorStatus(validator.Status, validator.Jailed))
	validatorInfo.ValidatorOperAddr = validator.GetOperator().String()
	accAddre := sdk.AccAddress(validator.GetOperator())
	validatorInfo.AccAddr = accAddre.String()
	return validatorInfo, nil
}


func (this *GatewayClient) FindValidatorByConsAddress(bech32ConsAddr string) (validator *stakingTypes.Validator, notFound bool, err error) {
	log := core.BuildLog(core.GetStructFuncName(this), core.LmChainClient)
	notFound = false
	consAddress, err := sdk.ConsAddressFromBech32(bech32ConsAddr)
	if err != nil {
		log.WithError(err).Error("ConsAddressFromBech32")
		err = errors.New(rest.ParseAccountError)
		return
	}
	params := types.QueryValidatorByConsAddrParams{ValidatorConsAddress: consAddress}
	bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		log.WithError(err).Error("MarshalJSON")
		return nil, notFound, err
	}

	resBytes, _, err := clientCtx.QueryWithData("custom/comm/"+types.QueryValidatorByConsAddress, bz)
	if err != nil {
		
		if strings.Contains(err.Error(), stakingTypes.ErrNoValidatorFound.Error()) {
			notFound = true
			err = nil 
		} else {
			log.WithError(err).Error("QueryWithData")
		}
		return nil, notFound, err
	}
	validator = &stakingTypes.Validator{}

	err = clientCtx.LegacyAmino.UnmarshalJSON(resBytes, validator)
	if err != nil {
		log.WithError(err).Error("UnmarshalJSON")
	}
	return
}


func (this *GatewayClient) GetValidatorStatus(status stakingTypes.BondStatus, jailed bool) int {
	if jailed {
		return 3
	} else {
		switch status.String() {
		case "BOND_STATUS_UNBONDED":
			return 0 
		case "BOND_STATUS_UNBONDING":
			return 1 
		case "BOND_STATUS_BONDED":
			return 2 
		}
	}
	return 0
}
