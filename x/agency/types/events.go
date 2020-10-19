package types

// agency module event types
const (
	EventTypeAuthorizeAgent   = "AuthorizeAgent"
	EventTypeDeauthorizeAgent = "DeauthorizeAgent"
	EventTypeAgentAction      = "AgentAction"

	AttributePrincipal = "principal"
	AttributeAgent     = "agent"
	AttributeScope     = "scope"
	AttributeLimit     = "limit"
	AttributeAction    = "action"
	AttributeAmount    = "amount"

	AttributeValueCategory = ModuleName
)
