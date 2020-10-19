package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Agency define data structure for Agent authorization
type Agency struct {
	Principal sdk.AccAddress `json:"principal" yaml:"principal"`
	Agent     sdk.AccAddress `json:"agent" yaml:"agent"`
	Scope     string         `json:"scope" yaml:"scope"`
	Limit     sdk.Coins      `json:"limit" yaml:"limit"`
}

func (a Agency) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Principal: %s
	Agent: %s
	Scope: %s
	Limit: %s`,
		a.Principal,
		a.Agent,
		a.Scope,
		a.Limit,
	))
}
