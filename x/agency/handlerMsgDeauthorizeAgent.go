package agency

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zenfa/agency/x/agency/keeper"
	"github.com/zenfa/agency/x/agency/types"
)

func handleMsgDeauthorizeAgent(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeauthorizeAgent) (*sdk.Result, error) {
	agency, err := k.GetAgency(ctx, msg.Principal, msg.Agent)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Cannot find the Agency")
	}

	k.DeleteAgency(ctx, agency)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeDeauthorizeAgent),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Principal.String()),
			sdk.NewAttribute(types.AttributeAgent, msg.Agent.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
