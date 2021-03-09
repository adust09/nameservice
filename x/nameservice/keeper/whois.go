package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/user/nameservice/x/nameservice/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// GetWhoisCount get the total number of whois
func (k Keeper) GetWhoisCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.WhoisCountPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetWhoisCount set the total number of whois
func (k Keeper) SetWhoisCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.WhoisCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateWhois creates a whois
// func (k Keeper) CreateWhois(ctx sdk.Context, msg types.MsgCreateWhois) {
// 	// Create the whois
// 	count := k.GetWhoisCount(ctx)
//     var whois = types.Whois{
//         Creator: msg.Creator,
//         ID:      strconv.FormatInt(count, 10),
//         Value: msg.Value,
//         Price: msg.Price,
//     }

// 	store := ctx.KVStore(k.storeKey)
// 	key := []byte(types.WhoisPrefix + whois.ID)
// 	value := k.cdc.MustMarshalBinaryLengthPrefixed(whois)
// 	store.Set(key, value)

// 	// Update whois count
//     k.SetWhoisCount(ctx, count+1)
// }

// GetWhois returns the whois information
func (k Keeper) GetWhois(ctx sdk.Context, key string) (types.Whois, error) {
	store := ctx.KVStore(k.storeKey)
	var whois types.Whois
	byteKey := []byte(types.WhoisPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &whois)
	if err != nil {
		return whois, err
	}
	return whois, nil
}

// SetWhois sets a whois
func (k Keeper) SetWhois(ctx sdk.Context, whois types.Whois) {
	whoisKey := whois.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(whois)
	key := []byte(types.WhoisPrefix + whoisKey)
	store.Set(key, bz)
}

// DeleteWhois deletes a whois
func (k Keeper) DeleteWhois(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.WhoisPrefix + key))
}

//
// Functions used by querier
//

func listWhois(ctx sdk.Context, k Keeper) ([]byte, error) {
	var whoisList []types.Whois
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.WhoisPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var whois types.Whois
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &whois)
		whoisList = append(whoisList, whois)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, whoisList)
	return res, nil
}

func getWhois(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	whois, err := k.GetWhois(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, whois)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

//Resolve name, return the value
func resolveName(ctx sdk.Context, path []string,keeper Keeper)([]byte,error){
	value := keeper.resolveName(ctx,path[0])

	if value =""{
	return []byte{}, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,"could not resolve name")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc,types.QueryResresolve{Value: value})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

//get creator of the item
func (k Keeper) GetCreator(ctx sdk.Context,key string) sdk.AccAddress{
	whois, _ := k.GetWhois(ctx,key)
	return whois.Creator
}

//Check if the key exists in the store
func (k Keeper) Exist(ctx sdk.Context,key string)bool{
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.WhoisCountPrefix + key))
}

//Resolvename - returns the steing that the name resolve to
func (k Keeper) Resolvename(ctx, sdk.Context, name string,value string){
	whois, _ := k.GetWhois(ctx, name)
	return whois.Value
}

//Setname - sets the value string tha a name resolves to
func (k Keeper) SetName(ctx sdk.Context, name string, value string){
	whois, _ := k.getWhois(ctx, name)
	whois.Value = value
	k.SetWhois(ctx,name,whois)
}

// HasOwner - returns whether or not the name already has an owner
func (k Keeper) HasCreator(ctx, sdk.Context, name string)bool{
	whois, _ := k.GetWhois(ctx, name)
	return !whois.Creator.Empty()
}

//SetOwner - sets the current owner of a name
func (k Keeper) SetCreator(ctx sdk.Context, name string, creator sdk.AccAddress){
	whois, _ := k.GetWhois(ctx, name)
	whois.Creator = creator
	k.SetWhois(ctx,name,whois)
}

// GetPrice - gets the current price of a name
func (k Keeper) GetPrive(ctx sdk.Context,name string)sdk.Coins{
	whois, _ := k.GetWhois(ctx,name)
	return whois.Price
}

// SetPrice - sets thw current price of a name
func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins{
	whois, _ := k.GetWhois(ctx,name)
	return whois.Price
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.WhoisPrefix))
}

// Get creator of the item
func (k Keeper) GetWhoisOwner(ctx sdk.Context, key string) sdk.AccAddress {
	whois, err := k.GetWhois(ctx, key)
	if err != nil {
		return nil
	}
	return whois.Creator
}


// Check if the key exists in the store
func (k Keeper) WhoisExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.WhoisPrefix + key))
}