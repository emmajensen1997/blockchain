package ante

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var minCommission = sdk.NewDecWithPrec(5, 2)





type ValidatorCommissionDecorator struct {
	cdc codec.BinaryCodec
}


func NewValidatorCommissionDecorator(cdc codec.BinaryCodec) ValidatorCommissionDecorator {
	return ValidatorCommissionDecorator{
		cdc: cdc,
	}
}



func (vcd ValidatorCommissionDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		switch msg := msg.(type) {
		case *authz.MsgExec:

			if err := vcd.validateAuthz(ctx, msg); err != nil {
				return ctx, err
			}

		default:
			if err := vcd.validateMsg(ctx, msg); err != nil {
				return ctx, err
			}
		}
	}

	return next(ctx, tx, simulate)
}


func (vcd ValidatorCommissionDecorator) validateAuthz(ctx sdk.Context, execMsg *authz.MsgExec) error {
	for _, v := range execMsg.Msgs {
		var innerMsg sdk.Msg
		err := vcd.cdc.UnpackAny(v, &innerMsg)
		if err != nil {
			return sdkerrors.Wrap(err, "cannot unmarshal authz exec msgs")
		}

		if err := vcd.validateMsg(ctx, innerMsg); err != nil {
			return err
		}
	}

	return nil
}


func (vcd ValidatorCommissionDecorator) validateMsg(_ sdk.Context, msg sdk.Msg) error {
	switch msg := msg.(type) {
	case *stakingtypes.MsgCreateValidator:
		if msg.Commission.Rate.LT(minCommission) {
			return sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"validator commission %s be lower than minimum of %s", msg.Commission.Rate, minCommission)
		}
	case *stakingtypes.MsgEditValidator:
		if msg.CommissionRate != nil && msg.CommissionRate.LT(minCommission) {
			return sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"validator commission %s be lower than minimum of %s", msg.CommissionRate, minCommission)
		}
	}
	return nil
}
