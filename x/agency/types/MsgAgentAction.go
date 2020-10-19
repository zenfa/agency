package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAgentAction{}

// MsgAgentAction defines actions taken by an agent
type MsgAgentAction struct {
	Principal sdk.AccAddress `json:"principal" yaml:"principal"` // address of the principal
	Agent     sdk.AccAddress `json:"agent" yaml:"agent"`         // address of the agent
	Action    string         `json:"action" yaml:"action"`       // maybe other actions in future, but now just transfer coins
	Amount    sdk.Coins      `json:"amount" yaml:"amount"`
}

// NewMsgAgentAction creates a new MsgAgentAction instance
func NewMsgAgentAction(principal sdk.AccAddress, agent sdk.AccAddress, action string, amount sdk.Coins) MsgAgentAction {
	return MsgAgentAction{
		Principal: principal,
		Agent:     agent,
		Action:    action,
		Amount:    amount,
	}
}

// nolint
func (msg MsgAgentAction) Route() string { return RouterKey }
func (msg MsgAgentAction) Type() string  { return "AgentAction" }
func (msg MsgAgentAction) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Agent)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgAgentAction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgAgentAction) ValidateBasic() error {
	if msg.Agent.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "agent can't be empty")
	}
	return nil
}
