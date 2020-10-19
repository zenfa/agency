package types

const (
	// ModuleName is the name of the module
	ModuleName = "agency"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

// AgencyPrefix defines prefix for the module
const (
	AgencyPrefix = "agency-"
)
