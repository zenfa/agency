package agency

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zenfa/agency/x/agency/keeper"
	"github.com/zenfa/agency/x/agency/types"
)

func handleMsgAuthorizeAgent(ctx sdk.Context, k keeper.Keeper, msg types.MsgAuthorizeAgent) (*sdk.Result, error) {
	var agency = types.Agency{
		Principal: msg.Principal,
		Agent:     msg.Agent,
		Scope:     msg.Scope,
		Limit:     msg.Limit,
	}
	_, err := k.GetAgency(ctx, agency.Principal, agency.Agent)
	// should produce an error when agency is found
	if err == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Agency with that Principal and Agent already exists")
	}
	k.CreateAgency(ctx, agency)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeAuthorizeAgent),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Principal.String()),
			sdk.NewAttribute(types.AttributeAgent, msg.Agent.String()),
			sdk.NewAttribute(types.AttributeScope, msg.Scope),
			sdk.NewAttribute(types.AttributeLimit, msg.Limit.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
