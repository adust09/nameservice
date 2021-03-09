package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Whois struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    Value string `json:"value" yaml:"value"`
    Price string `json:"price" yaml:"price"`
}

func NewWhois(){
	return Whois{
		Price: MinNamePrice,
	}
}
