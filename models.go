package sdm


// CreateResponseMetadata is reserved for future use.
type CreateResponseMetadata struct {
}
	
// GetResponseMetadata is reserved for future use.
type GetResponseMetadata struct {
}
	
// UpdateResponseMetadata is reserved for future use.
type UpdateResponseMetadata struct {
}
	
// DeleteResponseMetadata is reserved for future use.
type DeleteResponseMetadata struct {
}
	
// NodeCreateResponse reports how the Nodes were created in the system.
type NodeCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata
	// The created Node.
	Node Node
	// The auth token generated for the Node. The Node will use this token to
	// authenticate with the strongDM API.
	Token string
}
	
// NodeGetResponse returns a requested Node.
type NodeGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata
	// The requested Node.
	Node Node
}
	
// NodeUpdateResponse returns the fields of a Node after it has been updated by
// a NodeUpdateRequest.
type NodeUpdateResponse struct {
	// Reserved for future use.
	Meta *UpdateResponseMetadata
	// The updated Node.
	Node Node
}
	
// NodeDeleteResponse returns information about a Node that was deleted.
type NodeDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata
}
	
// A Node is a proxy in the strongDM network. They come in two flavors: relays,
// which communicate with resources, and gateways, which communicate with
// clients.
type Node interface {
	// GetID returns the unique identifier of the Node.
	GetID() string
	isOneOf_Node()
}
func (*Relay) isOneOf_Node(){}
// GetID returns the unique identifier of the Relay.
func (m *Relay) GetID() string { return m.ID }

func (*Gateway) isOneOf_Node(){}
// GetID returns the unique identifier of the Gateway.
func (m *Gateway) GetID() string { return m.ID }

// Relay represents a StrongDM CLI installation running in relay mode.
type Relay struct {
	// Unique identifier of the Relay.
	ID string
	// Unique human-readable name of the Relay.
	Name string
	// The current state of the relay. One of: "new", "verifying_restart",
	// "restarting", "started", "stopped", "dead", "unknown",
	State string
}
	
// Gateway represents a StrongDM CLI installation running in gateway mode.
type Gateway struct {
	// Unique identifier of the Relay.
	ID string
	// Unique human-readable name of the Relay.
	Name string
	// The current state of the gateway. One of: "new", "verifying_restart",
	// "restarting", "started", "stopped", "dead", "unknown",
	State string
	// The public hostname/port tuple at which the gateway will be accessible to clients.
	ListenAddress string
	// The hostname/port tuple which the gateway daemon will bind to.
	BindAddress string
}
	
// RoleCreateResponse reports how the Roles were created in the system. It can
// communicate partial successes or failures.
type RoleCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata
	// The created Role.
	Role *Role
}
	
// RoleGetResponse returns a requested Role.
type RoleGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata
	// The requested Role.
	Role *Role
}
	
// RoleUpdateResponse returns the fields of a Role after it has been updated by
// a RoleUpdateRequest.
type RoleUpdateResponse struct {
	// Reserved for future use.
	Meta *UpdateResponseMetadata
	// The updated Role.
	Role *Role
}
	
// RoleDeleteResponse returns information about a Role that was deleted.
type RoleDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata
}
	
// A Role grants users access to a set of resources. Composite roles have no
// resource associations of their own, but instead grant access to the combined
// resources of their child roles.
type Role struct {
	// Unique identifier of the Role.
	ID string
	// Unique human-readable name of the Role.
	Name string
	// True if the Role is a composite role.
	Composite bool
}
	
// NodeIterator provides read access to a list of Node.
// Use it like so:
// ```
// for iterator.Next() {
//     node := iterator.Value()
//     // ...
// }
// ```
type NodeIterator interface {
	// Next advances the iterator to the next item in the list. It returns
	// true if an item is available to retrieve via the `Value()` function.
	Next() bool
	// Value returns the current item, if one is available.
	Value() Node
	// Err returns the first error encountered during iteration, if any.
	Err() error
}
// RoleIterator provides read access to a list of Role.
// Use it like so:
// ```
// for iterator.Next() {
//     role := iterator.Value()
//     // ...
// }
// ```
type RoleIterator interface {
	// Next advances the iterator to the next item in the list. It returns
	// true if an item is available to retrieve via the `Value()` function.
	Next() bool
	// Value returns the current item, if one is available.
	Value() *Role
	// Err returns the first error encountered during iteration, if any.
	Err() error
}