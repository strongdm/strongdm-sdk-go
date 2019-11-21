package models


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
	
// NodeCreateResponse reports how the nodes were created in the system. It can
// communicate partial successes or failures.
type NodeCreateResponse struct {
	Meta CreateResponseMetadata
	Node Node
	Token Token
}
	
// NodeGetResponse returns a requested node.
type NodeGetResponse struct {
	Meta GetResponseMetadata
	Node Node
}
	
// NodeUpdateResponse returns the fields of a node after it has been updated by
// a NodeUpdateRequest.
type NodeUpdateResponse struct {
	Meta UpdateResponseMetadata
	Node Node
}
	
// NodeDeleteResponse returns information about a node that was deleted.
type NodeDeleteResponse struct {
	Meta DeleteResponseMetadata
}
	
type Node interface {
	GetID() string
	isOneOf_Node()
}
func (*Relay) isOneOf_Node(){}
func (m *Relay) GetID() string { return m.ID }

func (*Gateway) isOneOf_Node(){}
func (m *Gateway) GetID() string { return m.ID }

// Relay represents a StrongDM CLI installation running in relay mode.
type Relay struct {
	ID string
	Name string
}
	
// Gateway represents a StrongDM CLI installation running in gateway mode.
type Gateway struct {
	ID string
	Name string
	ListenAddress string
	BindAddress string
}
	
// Token holds the bearer token used to start up nodes.
type Token struct {
	ID string
	Token string
}
	
// RoleCreateResponse reports how the Roles were created in the system. It can
// communicate partial successes or failures.
type RoleCreateResponse struct {
	Meta CreateResponseMetadata
	Role Role
}
	
// RoleGetResponse returns a requested Role.
type RoleGetResponse struct {
	Meta GetResponseMetadata
	Role Role
}
	
// RoleUpdateResponse returns the fields of a Role after it has been updated by
// a RoleUpdateRequest.
type RoleUpdateResponse struct {
	Meta UpdateResponseMetadata
	Role Role
}
	
// RoleDeleteResponse returns information about a Role that was deleted.
type RoleDeleteResponse struct {
	Meta DeleteResponseMetadata
}
	
// A Role grants users access to a set of resources. Composite roles have no
// resource associations of their own, but instead grant access to the combined
// resources of their child roles.
type Role struct {
	ID string
	Name string
	Composite bool
}
	
type NodeIterator interface {
	Next() bool
	Value() Node
	Err() error
}
type RoleIterator interface {
	Next() bool
	Value() Role
	Err() error
}