package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

// keeper of the stake store
type Keeper struct {
	storeKey            sdk.StoreKey
	storeTKey           sdk.StoreKey
	cdc                 *codec.Codec
	paramSpace          params.Subspace
	bankKeeper          types.BankKeeper
	stakeKeeper         types.StakeKeeper
	feeCollectionKeeper types.FeeCollectionKeeper

	// codespace
	codespace sdk.CodespaceType
}

func NewKeeper(cdc *codec.Codec, key, tkey sdk.StoreKey, paramSpace params.Subspace, ck types.BankKeeper,
	sk types.StakeKeeper, fck types.FeeCollectionKeeper, codespace sdk.CodespaceType) Keeper {

	keeper := Keeper{
		storeKey:            key,
		storeTKey:           tkey,
		cdc:                 cdc,
		paramSpace:          paramSpace.WithTypeTable(ParamTypeTable()),
		bankKeeper:          ck,
		stakeKeeper:         sk,
		feeCollectionKeeper: fck,
		codespace:           codespace,
	}
	return keeper
}

//______________________________________________________________________

// get the global fee pool distribution info
func (k Keeper) GetFeePool(ctx sdk.Context) (feePool types.FeePool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(FeePoolKey)
	if b == nil {
		panic("Stored fee pool should not have been nil")
	}
	k.cdc.MustUnmarshalBinary(b, &feePool)
	return
}

// set the global fee pool distribution info
func (k Keeper) SetFeePool(ctx sdk.Context, feePool types.FeePool) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinary(feePool)
	store.Set(FeePoolKey, b)
}

//______________________________________________________________________

// set the proposer public key for this block
func (k Keeper) GetProposerConsAddr(ctx sdk.Context) (consAddr sdk.ConsAddress) {
	tstore := ctx.KVStore(k.storeTKey)

	b := tstore.Get(ProposerKey)
	if b == nil {
		panic("Proposer cons address was likely not set in begin block")
	}

	k.cdc.MustUnmarshalBinary(b, &consAddr)
	return
}

// get the proposer public key for this block
func (k Keeper) SetProposerConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) {
	tstore := ctx.KVStore(k.storeTKey)
	b := k.cdc.MustMarshalBinary(consAddr)
	tstore.Set(ProposerKey, b)
}

//______________________________________________________________________

// set the proposer public key for this block
func (k Keeper) GetPercentPrecommitVotes(ctx sdk.Context) (percentPrecommitVotes sdk.Dec) {
	tstore := ctx.KVStore(k.storeTKey)

	b := tstore.Get(PercentPrecommitVotesKey)
	if b == nil {
		panic("Proposer cons address was likely not set in begin block")
	}

	k.cdc.MustUnmarshalBinary(b, &percentPrecommitVotes)
	return
}

// get the proposer public key for this block
func (k Keeper) SetPercentPrecommitVotes(ctx sdk.Context, percentPrecommitVotes sdk.Dec) {
	tstore := ctx.KVStore(k.storeTKey)
	b := k.cdc.MustMarshalBinary(percentPrecommitVotes)
	tstore.Set(PercentPrecommitVotesKey, b)
}

//______________________________________________________________________
// PARAM STORE

// Type declaration for parameters
func ParamTypeTable() params.TypeTable {
	return params.NewTypeTable(
		ParamStoreKeyCommunityTax, sdk.Dec{},
		ParamStoreKeyBaseProposerReward, sdk.Dec{},
		ParamStoreKeyBonusProposerReward, sdk.Dec{},
	)
}

// Returns the current CommunityTax rate from the global param store
// nolint: errcheck
func (k Keeper) GetCommunityTax(ctx sdk.Context) sdk.Dec {
	var percent sdk.Dec
	k.paramSpace.Get(ctx, ParamStoreKeyCommunityTax, &percent)
	return percent
}

// nolint: errcheck
func (k Keeper) SetCommunityTax(ctx sdk.Context, percent sdk.Dec) {
	k.paramSpace.Set(ctx, ParamStoreKeyCommunityTax, &percent)
}

// Returns the current BaseProposerReward rate from the global param store
// nolint: errcheck
func (k Keeper) GetBaseProposerReward(ctx sdk.Context) sdk.Dec {
	var percent sdk.Dec
	k.paramSpace.Get(ctx, ParamStoreKeyBaseProposerReward, &percent)
	return percent
}

// nolint: errcheck
func (k Keeper) SetBaseProposerReward(ctx sdk.Context, percent sdk.Dec) {
	k.paramSpace.Set(ctx, ParamStoreKeyBaseProposerReward, &percent)
}

// Returns the current BaseProposerReward rate from the global param store
// nolint: errcheck
func (k Keeper) GetBonusProposerReward(ctx sdk.Context) sdk.Dec {
	var percent sdk.Dec
	k.paramSpace.Get(ctx, ParamStoreKeyBonusProposerReward, &percent)
	return percent
}

// nolint: errcheck
func (k Keeper) SetBonusProposerReward(ctx sdk.Context, percent sdk.Dec) {
	k.paramSpace.Set(ctx, ParamStoreKeyBonusProposerReward, &percent)
}
