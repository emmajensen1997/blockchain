package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)


const (
	QueryValidators                    = "validators"
	QueryValidator                     = "validator"
	QueryDelegatorDelegations          = "delegatorDelegations"
	QueryDelegatorUnbondingDelegations = "delegatorUnbondingDelegations"
	QueryRedelegations                 = "redelegations"
	QueryValidatorDelegations          = "validatorDelegations"
	QueryValidatorRedelegations        = "validatorRedelegations"
	QueryValidatorUnbondingDelegations = "validatorUnbondingDelegations"
	QueryDelegation                    = "delegation"
	QueryUnbondingDelegation           = "unbondingDelegation"
	QueryDelegatorValidators           = "delegatorValidators"
	QueryDelegatorValidator            = "delegatorValidator"
	QueryPool                          = "pool"
	QueryParameters                    = "parameters"
	QueryHistoricalInfo                = "historicalInfo"
)





type QueryDelegatorParams struct {
	DelegatorAddr sdk.AccAddress
}

func NewQueryDelegatorParams(delegatorAddr sdk.AccAddress) QueryDelegatorParams {
	return QueryDelegatorParams{
		DelegatorAddr: delegatorAddr,
	}
}





type QueryValidatorParams struct {
	ValidatorAddr sdk.ValAddress
	Page, Limit   int
}

func NewQueryValidatorParams(validatorAddr sdk.ValAddress, page, limit int) QueryValidatorParams {
	return QueryValidatorParams{
		ValidatorAddr: validatorAddr,
		Page:          page,
		Limit:         limit,
	}
}



type QueryRedelegationParams struct {
	DelegatorAddr    sdk.AccAddress
	SrcValidatorAddr sdk.ValAddress
	DstValidatorAddr sdk.ValAddress
}

func NewQueryRedelegationParams(delegatorAddr sdk.AccAddress, srcValidatorAddr, dstValidatorAddr sdk.ValAddress) QueryRedelegationParams {
	return QueryRedelegationParams{
		DelegatorAddr:    delegatorAddr,
		SrcValidatorAddr: srcValidatorAddr,
		DstValidatorAddr: dstValidatorAddr,
	}
}



type QueryValidatorsParams struct {
	Page, Limit int
	Status      string
}

func NewQueryValidatorsParams(page, limit int, status string) QueryValidatorsParams {
	return QueryValidatorsParams{page, limit, status}
}