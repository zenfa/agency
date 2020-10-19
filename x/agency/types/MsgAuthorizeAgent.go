package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAuthorizeAgent{}

// MsgAuthorizeAgent - users(principals) can authorize an agent to access their resources
type MsgAuthorizeAgent struct {
	Principal sdk.AccAddress `json:"principal" yaml:"principal"`
	Agent     sdk.AccAddress `json:"agent" yaml:"agent"`
	// Description  string         `json:"description" yaml:"description"`
	Scope string `json:"scope" yaml:"scope"`
	// bank for this demo
	// maybe other things in future, like documents to collabarate, or user profile
	Limit sdk.Coins `json:"limit" yaml:"limit"`
	// Reward       sdk.Coins      `json:"reward" yaml:"reward"`
}

// NewMsgAuthorizeAgent creates a new NewMsgAuthorizeAgent instance
func NewMsgAuthorizeAgent(principal sdk.AccAddress, agent sdk.AccAddress, scope string, limit sdk.Coins) MsgAuthorizeAgent {
	return MsgAuthorizeAgent{
		Principal: principal,
		Agent:     agent,
		Scope:     scope,
		Limit:     limit,
	}
}

// nolint
func (msg MsgAuthorizeAgent) Route() string { return RouterKey }
func (msg MsgAuthorizeAgent) Type() string  { return "AuthorizeAgent" }
func (msg MsgAuthorizeAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Principal)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgAuthorizeAgent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgAuthorizeAgent) ValidateBasic() error {
	if msg.Principal.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Principal can't be empty")
	}
	if msg.Scope == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Scope can't be empty")
	}
	return nil
}
