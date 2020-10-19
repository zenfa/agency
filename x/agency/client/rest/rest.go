package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers agency-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding
	r.HandleFunc("/agency/agency", listAgencyHandler(cliCtx, "agency")).Methods("GET")
	r.HandleFunc("/agency/agency/{principal}/{agent}", getAgencyHandler(cliCtx, "agency")).Methods("GET")
	r.HandleFunc("/agency/authorize", authorizeAgentHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/agency/deauthorize", deauthorizeAgentHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/agency/action", agentActionHandler(cliCtx)).Methods("POST")
}
