package tx

import (
	"fmt"

	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
)


type signModeDirectHandler struct{}

var _ signing.SignModeHandler = signModeDirectHandler{}


func (signModeDirectHandler) DefaultMode() signingtypes.SignMode {
	return signingtypes.SignMode_SIGN_MODE_DIRECT
}


func (signModeDirectHandler) Modes() []signingtypes.SignMode {
	return []signingtypes.SignMode{signingtypes.SignMode_SIGN_MODE_DIRECT}
}


func (signModeDirectHandler) GetSignBytes(mode signingtypes.SignMode, data signing.SignerData, tx sdk.Tx) ([]byte, error) {
	if mode != signingtypes.SignMode_SIGN_MODE_DIRECT {
		return nil, fmt.Errorf("expected %s, got %s", signingtypes.SignMode_SIGN_MODE_DIRECT, mode)
	}

	protoTx, ok := tx.(*wrapper)
	if !ok {
		return nil, fmt.Errorf("can only handle a protobuf Tx, got %T", tx)
	}

	bodyBz := protoTx.getBodyBytes()
	authInfoBz := protoTx.getAuthInfoBytes()

	return DirectSignBytes(bodyBz, authInfoBz, data.ChainID, data.AccountNumber)
}



func DirectSignBytes(bodyBytes, authInfoBytes []byte, chainID string, accnum uint64) ([]byte, error) {
	signDoc := types.SignDoc{
		BodyBytes:     bodyBytes,
		AuthInfoBytes: authInfoBytes,
		ChainId:       chainID,
		AccountNumber: accnum,
	}
	return signDoc.Marshal()
}
