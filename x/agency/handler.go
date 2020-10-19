package agency

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zenfa/agency/x/agency/keeper"
	"github.com/zenfa/agency/x/agency/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// this line is used by starport scaffolding
		case types.MsgAuthorizeAgent:
			return handleMsgAuthorizeAgent(ctx, k, msg)
		case types.MsgDeauthorizeAgent:
			return handleMsgDeauthorizeAgent(ctx, k, msg)
		case types.MsgAgentAction:
			return handleMsgAgentAction(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
