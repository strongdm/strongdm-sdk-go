package sdm

import (
	"time"
)

// A Resource is a server or service which clients connect to through relays.
type Resource interface {
	// GetID returns the unique identifier of the Resource.
	GetID() string
	isOneOf_Resource()
}

func (*Redis) isOneOf_Resource() {}

// GetID returns the unique identifier of the Redis.
func (m *Redis) GetID() string { return m.ID }

func (*ElasticacheRedis) isOneOf_Resource() {}

// GetID returns the unique identifier of the ElasticacheRedis.
func (m *ElasticacheRedis) GetID() string { return m.ID }

func (*Kubernetes) isOneOf_Resource() {}

// GetID returns the unique identifier of the Kubernetes.
func (m *Kubernetes) GetID() string { return m.ID }

func (*KubernetesBasicAuth) isOneOf_Resource() {}

// GetID returns the unique identifier of the KubernetesBasicAuth.
func (m *KubernetesBasicAuth) GetID() string { return m.ID }

func (*AmazonEKS) isOneOf_Resource() {}

// GetID returns the unique identifier of the AmazonEKS.
func (m *AmazonEKS) GetID() string { return m.ID }

func (*GoogleGKE) isOneOf_Resource() {}

// GetID returns the unique identifier of the GoogleGKE.
func (m *GoogleGKE) GetID() string { return m.ID }

func (*SSH) isOneOf_Resource() {}

// GetID returns the unique identifier of the SSH.
func (m *SSH) GetID() string { return m.ID }

func (*HTTPBasicAuth) isOneOf_Resource() {}

// GetID returns the unique identifier of the HTTPBasicAuth.
func (m *HTTPBasicAuth) GetID() string { return m.ID }

func (*HTTPNoAuth) isOneOf_Resource() {}

// GetID returns the unique identifier of the HTTPNoAuth.
func (m *HTTPNoAuth) GetID() string { return m.ID }

func (*HTTPAuth) isOneOf_Resource() {}

// GetID returns the unique identifier of the HTTPAuth.
func (m *HTTPAuth) GetID() string { return m.ID }

func (*Mysql) isOneOf_Resource() {}

// GetID returns the unique identifier of the Mysql.
func (m *Mysql) GetID() string { return m.ID }

func (*AuroraMysql) isOneOf_Resource() {}

// GetID returns the unique identifier of the AuroraMysql.
func (m *AuroraMysql) GetID() string { return m.ID }

func (*Clustrix) isOneOf_Resource() {}

// GetID returns the unique identifier of the Clustrix.
func (m *Clustrix) GetID() string { return m.ID }

func (*Maria) isOneOf_Resource() {}

// GetID returns the unique identifier of the Maria.
func (m *Maria) GetID() string { return m.ID }

func (*Memsql) isOneOf_Resource() {}

// GetID returns the unique identifier of the Memsql.
func (m *Memsql) GetID() string { return m.ID }

func (*Athena) isOneOf_Resource() {}

// GetID returns the unique identifier of the Athena.
func (m *Athena) GetID() string { return m.ID }

type Redis struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	PortOverride int32

	Password string

	Port int32
}

type ElasticacheRedis struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	PortOverride int32

	Password string

	Port int32

	TlsRequired bool
}

type Kubernetes struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Port int32

	CertificateAuthority string

	ClientCertificate string

	ClientKey string
}

type KubernetesBasicAuth struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Port int32

	Username string

	Password string

	CertificateAuthority string

	ClientCertificate string

	ClientKey string
}

type AmazonEKS struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Endpoint string

	AccessKey string

	SecretAccessKey string

	CertificateAuthority string

	Region string

	ClusterName string
}

type GoogleGKE struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Endpoint string

	CertificateAuthority string

	ServiceAccountKey string
}

type SSH struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Port int32

	PublicKey string
}

type HTTPBasicAuth struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Url string

	HealthcheckPath string

	Username string

	Password string

	HeadersBlacklist string

	DefaultPath string

	Subdomain string
}

type HTTPNoAuth struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Url string

	HealthcheckPath string

	HeadersBlacklist string

	DefaultPath string

	Subdomain string
}

type HTTPAuth struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Url string

	HealthcheckPath string

	AuthHeader string

	HeadersBlacklist string

	DefaultPath string

	Subdomain string
}

type Mysql struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32
}

type AuroraMysql struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32
}

type Clustrix struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32
}

type Maria struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32
}

type Memsql struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32
}

type Athena struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	AccessKey string

	SecretAccessKey string

	Output string

	PortOverride int32

	Region string
}

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

// RateLimitMetadata contains information about remaining requests avaialable
// to the user over some timeframe.
type RateLimitMetadata struct {
	// How many total requests the user/token is authorized to make before being
	// rate limited.
	Limit int64
	// How many remaining requests out of the limit are still avaialable.
	Remaining int64
	// The time when remaining will be reset to limit.
	ResetAt time.Time
	// The bucket this user/token is associated with, which may be shared between
	// multiple users/tokens.
	Bucket string
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
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// NodeGetResponse returns a requested Node.
type NodeGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata
	// The requested Node.
	Node Node
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// NodeUpdateResponse returns the fields of a Node after it has been updated by
// a NodeUpdateRequest.
type NodeUpdateResponse struct {
	// Reserved for future use.
	Meta *UpdateResponseMetadata
	// The updated Node.
	Node Node
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// NodeDeleteResponse returns information about a Node that was deleted.
type NodeDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// A Node is a proxy in the strongDM network. They come in two flavors: relays,
// which communicate with resources, and gateways, which communicate with
// clients.
type Node interface {
	// GetID returns the unique identifier of the Node.
	GetID() string
	isOneOf_Node()
}

func (*Relay) isOneOf_Node() {}

// GetID returns the unique identifier of the Relay.
func (m *Relay) GetID() string { return m.ID }

func (*Gateway) isOneOf_Node() {}

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

// ResourceCreateResponse reports how the Resources were created in the system.
type ResourceCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata
	// The created Resource.
	Resource Resource
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// ResourceGetResponse returns a requested Resource.
type ResourceGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata
	// The requested Resource.
	Resource Resource
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// ResourceUpdateResponse returns the fields of a Resource after it has been updated by
// a ResourceUpdateRequest.
type ResourceUpdateResponse struct {
	// Reserved for future use.
	Meta *UpdateResponseMetadata
	// The updated Resource.
	Resource Resource
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// ResourceDeleteResponse returns information about a Resource that was deleted.
type ResourceDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// RoleAttachmentCreateResponse reports how the RoleAttachments were created in the system.
type RoleAttachmentCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata
	// The created RoleAttachment.
	RoleAttachment *RoleAttachment
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// RoleAttachmentGetResponse returns a requested RoleAttachment.
type RoleAttachmentGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata
	// The requested RoleAttachment.
	RoleAttachment *RoleAttachment
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// RoleAttachmentDeleteResponse returns information about a RoleAttachment that was deleted.
type RoleAttachmentDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// A RoleAttachment connects a composite role to another role, granting members
// of the composite role the permissions granted to the attached role.
type RoleAttachment struct {
	// Unique identifier of the RoleAttachment.
	ID string
	// The id of the composite role of this RoleAttachment.
	CompositeRoleID string
	// The id of the attached role of this RoleAttachment.
	AttachedRoleID string
}

// RoleCreateResponse reports how the Roles were created in the system. It can
// communicate partial successes or failures.
type RoleCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata
	// The created Role.
	Role *Role
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// RoleGetResponse returns a requested Role.
type RoleGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata
	// The requested Role.
	Role *Role
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// RoleUpdateResponse returns the fields of a Role after it has been updated by
// a RoleUpdateRequest.
type RoleUpdateResponse struct {
	// Reserved for future use.
	Meta *UpdateResponseMetadata
	// The updated Role.
	Role *Role
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// RoleDeleteResponse returns information about a Role that was deleted.
type RoleDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata
	// Rate limit information.
	RateLimit *RateLimitMetadata
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
//     for iterator.Next() {
//         node := iterator.Value()
//         // ...
//     }
type NodeIterator interface {
	// Next advances the iterator to the next item in the list. It returns
	// true if an item is available to retrieve via the `Value()` function.
	Next() bool
	// Value returns the current item, if one is available.
	Value() Node
	// Err returns the first error encountered during iteration, if any.
	Err() error
}

// ResourceIterator provides read access to a list of Resource.
// Use it like so:
//     for iterator.Next() {
//         resource := iterator.Value()
//         // ...
//     }
type ResourceIterator interface {
	// Next advances the iterator to the next item in the list. It returns
	// true if an item is available to retrieve via the `Value()` function.
	Next() bool
	// Value returns the current item, if one is available.
	Value() Resource
	// Err returns the first error encountered during iteration, if any.
	Err() error
}

// RoleAttachmentIterator provides read access to a list of RoleAttachment.
// Use it like so:
//     for iterator.Next() {
//         roleAttachment := iterator.Value()
//         // ...
//     }
type RoleAttachmentIterator interface {
	// Next advances the iterator to the next item in the list. It returns
	// true if an item is available to retrieve via the `Value()` function.
	Next() bool
	// Value returns the current item, if one is available.
	Value() *RoleAttachment
	// Err returns the first error encountered during iteration, if any.
	Err() error
}

// RoleIterator provides read access to a list of Role.
// Use it like so:
//     for iterator.Next() {
//         role := iterator.Value()
//         // ...
//     }
type RoleIterator interface {
	// Next advances the iterator to the next item in the list. It returns
	// true if an item is available to retrieve via the `Value()` function.
	Next() bool
	// Value returns the current item, if one is available.
	Value() *Role
	// Err returns the first error encountered during iteration, if any.
	Err() error
}
