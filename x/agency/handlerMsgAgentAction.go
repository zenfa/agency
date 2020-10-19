package agency

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zenfa/agency/x/agency/keeper"
	"github.com/zenfa/agency/x/agency/types"
)

func handleMsgAgentAction(ctx sdk.Context, k keeper.Keeper, msg types.MsgAgentAction) (*sdk.Result, error) {
	agency, err := k.GetAgency(ctx, msg.Principal, msg.Agent)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Agent is not allowed to transfer money from principal account")
	}

	if agency.Limit.IsAllGTE(msg.Amount) {
		sdkError := k.CoinKeeper.SendCoins(ctx, msg.Principal, msg.Agent, msg.Amount)
		if sdkError != nil {
			return nil, sdkError
		}
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Exceed limit")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeAgentAction),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Agent.String()),
			sdk.NewAttribute(types.AttributePrincipal, msg.Principal.String()),
			sdk.NewAttribute(types.AttributeAction, msg.Action),
			sdk.NewAttribute(types.AttributeAmount, msg.Amount.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
