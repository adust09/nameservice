package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetWhois{}

type MsgSetWhois struct {
  Owner sdk.AccAddress `json:"owner" yaml:"owner"`
  Value string `json:"value" yaml:"value"`
  Price string `json:"price" yaml:"price"`
}

func NewMsgSetWhois(owner sdk.AccAddress, value string, price string) MsgSetWhois {
  return MsgSetWhois{
		Owner: owner,
    Value: value,
    Price: price,
	}
}

func (msg MsgSetWhois) Route() string {
  return RouterKey
}

func (msg MsgSetWhois) Type() string {
  return "SetWhois"
}

func (msg MsgSetWhois) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg MsgSetWhois) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetWhois) ValidateBasic() error {
  if msg.Owner.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
  }
  return nil
}