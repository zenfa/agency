package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/zenfa/agency/x/agency/types"
)

type authorizeAgentRequest struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Principal string       `json:"principal"`
	Agent     string       `json:"agent"`
	Scope     string       `json:"scope"`
	Limit     string       `json:"limit"`
}

func authorizeAgentHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req authorizeAgentRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		principal, err := sdk.AccAddressFromBech32(req.Principal)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		agent, err := sdk.AccAddressFromBech32(req.Agent)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		limit, err := sdk.ParseCoins(req.Limit)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgAuthorizeAgent(principal, agent, req.Scope, limit)
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deauthorizeAgentRequest struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Principal string       `json:"principal"`
	Agent     string       `json:"agent"`
}

func deauthorizeAgentHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deauthorizeAgentRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		principal, err := sdk.AccAddressFromBech32(req.Principal)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		agent, err := sdk.AccAddressFromBech32(req.Agent)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgDeauthorizeAgent(principal, agent)
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
