package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateWhois{}

type MsgCreateWhois struct {
  Owner sdk.AccAddress `json:"owner" yaml:"owner"`
  Value string `json:"value" yaml:"value"`
  Price string `json:"price" yaml:"price"`
}

func NewMsgCreateWhois(owner sdk.AccAddress, value string, price string) MsgCreateWhois {
  return MsgCreateWhois{
		Owner: owner,
    Value: value,
    Price: price,
	}
}

func (msg MsgCreateWhois) Route() string {
  return RouterKey
}

func (msg MsgCreateWhois) Type() string {
  return "CreateWhois"
}

func (msg MsgCreateWhois) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg MsgCreateWhois) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateWhois) ValidateBasic() error {
  if msg.Owner.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
  }
  return nil
}