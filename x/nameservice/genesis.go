package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/user/nameservice/x/nameservice/keeper"
	"github.com/user/nameservice/x/nameservice/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper , data types.GenesisState) {
	// TODO: Define logic for when you would like to initalize a new genesis
	for _, record := range data.WhoisRecord{
		keeper.SetWhois(ctx, record.Value,record)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	// TODO: Define logic for exporting state
	var record []types.whois
	iterator := k.GetNameIterator(ctx)
	for ; iterator.Valid(); iterator.Next(){
		name := string(iterator.key())
		whois, _ := k.GetWhois(ctx,name)
		record = append(record,whois)
	}
	return types.GenesisState{WhoisRecords: records}
}
