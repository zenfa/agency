package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zenfa/agency/x/agency/types"
)

func getAgencyKey(principal sdk.AccAddress, agent sdk.AccAddress) []byte {
	agencyKey := append(principal, agent...)
	key := append([]byte(types.AgencyPrefix), agencyKey...)

	return key
}

// CreateAgency stores an agency
func (k Keeper) CreateAgency(ctx sdk.Context, agency types.Agency) {
	store := ctx.KVStore(k.storeKey)
	// key := []byte(types.AgencyPrefix + commit.SolutionScavengerHash)
	key := getAgencyKey(agency.Principal, agency.Agent)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(agency)
	store.Set(key, value)
}

// DeleteAgency deletes a agency
func (k Keeper) DeleteAgency(ctx sdk.Context, agency types.Agency) {
	store := ctx.KVStore(k.storeKey)
	key := getAgencyKey(agency.Principal, agency.Agent)
	store.Delete(key)
}

// GetAgency returns an agency
// func (k Keeper) GetAgency(ctx sdk.Context, key string) (types.Agency, error) {
func (k Keeper) GetAgency(ctx sdk.Context, principal sdk.AccAddress, agent sdk.AccAddress) (types.Agency, error) {
	store := ctx.KVStore(k.storeKey)
	var agency types.Agency
	byteKey := getAgencyKey(principal, agent)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &agency)
	if err != nil {
		return agency, err
	}
	return agency, nil
}

func getAgency(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	principal, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}
	agent, err := sdk.AccAddressFromBech32(path[1])
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}
	agency, err := k.GetAgency(ctx, principal, agent)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, agency)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func listAgency(ctx sdk.Context, k Keeper) ([]byte, error) {
	var agencyList []types.Agency
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.AgencyPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var agency types.Agency
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &agency)
		agencyList = append(agencyList, agency)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, agencyList)
	return res, nil
}
