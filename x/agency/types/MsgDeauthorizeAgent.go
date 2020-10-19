package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgDeauthorizeAgent
// ------------------------------------------------------------------------------
var _ sdk.Msg = &MsgDeauthorizeAgent{}

// MsgDeauthorizeAgent - users(principals) deauthorize an agent
type MsgDeauthorizeAgent struct {
	Principal sdk.AccAddress `json:"principal" yaml:"principal"` // address of the user(principal)
	Agent     sdk.AccAddress `json:"agent" yaml:"agent"`         // agent to remove
}

// NewMsgDeauthorizeAgent creates a new MsgDeauthorizeAgent instance
func NewMsgDeauthorizeAgent(principal sdk.AccAddress, agent sdk.AccAddress) MsgDeauthorizeAgent {
	return MsgDeauthorizeAgent{
		Principal: principal,
		Agent:     agent,
	}
}

// nolint
func (msg MsgDeauthorizeAgent) Route() string { return RouterKey }
func (msg MsgDeauthorizeAgent) Type() string  { return "DeauthorizeAgent" }
func (msg MsgDeauthorizeAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Principal)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgDeauthorizeAgent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgDeauthorizeAgent) ValidateBasic() error {
	if msg.Principal.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "principal can't be empty")
	}
	if msg.Principal.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "agent can't be empty")
	}
	return nil
}
