package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/user/nameservice/x/nameservice/types"
	"github.com/user/nameservice/x/nameservice/keeper"
)

func handleMsgSetName(ctx sdk.Context, keeper.Keeper, msg MsgSetName) (*sdk.Result, error) {
	if !msg.Creator.Equals(k.GetNameOwner(ctx, msg.Name)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	keeper.SetName(ctx, msg.Name,msg.Value)

	return &sdk.Result{}, nil //return
}
