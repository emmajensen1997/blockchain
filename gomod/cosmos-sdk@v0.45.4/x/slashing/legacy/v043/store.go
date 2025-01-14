package v043

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v043distribution "github.com/cosmos/cosmos-sdk/x/distribution/legacy/v043"
	v040slashing "github.com/cosmos/cosmos-sdk/x/slashing/legacy/v040"
)



//

func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey) error {
	store := ctx.KVStore(storeKey)
	v043distribution.MigratePrefixAddress(store, v040slashing.ValidatorSigningInfoKeyPrefix)
	v043distribution.MigratePrefixAddressBytes(store, v040slashing.ValidatorMissedBlockBitArrayKeyPrefix)
	v043distribution.MigratePrefixAddress(store, v040slashing.AddrPubkeyRelationKeyPrefix)

	return nil
}
