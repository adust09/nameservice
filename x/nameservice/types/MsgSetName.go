package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

type NewMsgSetName struct {
  Name string `json:"name"`
  Value string `json:"value"`
  Owner sdk.AccAddress `json:"owner"`
}

// NewMsgSetName is a constructor function for MsgName
func NewNewMsgSetName(name string, value string, owner sdk.AccAddress) NewMsgSetName {
  return NewMsgSetName{
    Name: name,
    Value: value,
    Owner: owner,
	}
}

//Route should return the name of the module
func (msg NewMsgSetName) Route() string {return RouterKey}

//Type should return the action
func (msg MsgSetName) Type() string { return "set_name"}

//ValidawteBasic runs stateless checks on the message
func (msg MsgSetName) ValidateBasic() error {
  if msg.Owner.Empty(){
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
  }
  if len(msg.Name) == 0 || len(msg.Value) == 0{
    return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,"Nanme and/or Value cannot be empty")
  }
  return nil
}

//GetSignBytes encode the message for signing
func (msg MsgSetName) GetSignBytes() []byte{
  return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

