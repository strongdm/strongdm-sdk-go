package v1

import (
    "context"
	"google.golang.org/grpc/status"
    "github.com/strongdm/strongdm-sdk-go/models"
    "github.com/strongdm/strongdm-sdk-go/errors"
    "google.golang.org/grpc/metadata"
)

func CreateResponseMetadataToPorcelain(plumbing *CreateResponseMetadata) models.CreateResponseMetadata {
    porcelain := models.CreateResponseMetadata{}
    return porcelain
}

func CreateResponseMetadataToPlumbing(porcelain models.CreateResponseMetadata) *CreateResponseMetadata {
    plumbing := &CreateResponseMetadata{}
    return plumbing
}

func RepeatedCreateResponseMetadataToPlumbing(porcelains []models.CreateResponseMetadata) []*CreateResponseMetadata {
    var items []*CreateResponseMetadata
    for _, porcelain := range porcelains {
        items = append(items, CreateResponseMetadataToPlumbing(porcelain))
    }
    return items
}

func RepeatedCreateResponseMetadataToPorcelain(plumbings []*CreateResponseMetadata) []models.CreateResponseMetadata {
    var items []models.CreateResponseMetadata
    for _, plumbing := range plumbings {
        items = append(items, CreateResponseMetadataToPorcelain(plumbing))
    }
    return items
}

func GetResponseMetadataToPorcelain(plumbing *GetResponseMetadata) models.GetResponseMetadata {
    porcelain := models.GetResponseMetadata{}
    return porcelain
}

func GetResponseMetadataToPlumbing(porcelain models.GetResponseMetadata) *GetResponseMetadata {
    plumbing := &GetResponseMetadata{}
    return plumbing
}

func RepeatedGetResponseMetadataToPlumbing(porcelains []models.GetResponseMetadata) []*GetResponseMetadata {
    var items []*GetResponseMetadata
    for _, porcelain := range porcelains {
        items = append(items, GetResponseMetadataToPlumbing(porcelain))
    }
    return items
}

func RepeatedGetResponseMetadataToPorcelain(plumbings []*GetResponseMetadata) []models.GetResponseMetadata {
    var items []models.GetResponseMetadata
    for _, plumbing := range plumbings {
        items = append(items, GetResponseMetadataToPorcelain(plumbing))
    }
    return items
}

func UpdateResponseMetadataToPorcelain(plumbing *UpdateResponseMetadata) models.UpdateResponseMetadata {
    porcelain := models.UpdateResponseMetadata{}
    return porcelain
}

func UpdateResponseMetadataToPlumbing(porcelain models.UpdateResponseMetadata) *UpdateResponseMetadata {
    plumbing := &UpdateResponseMetadata{}
    return plumbing
}

func RepeatedUpdateResponseMetadataToPlumbing(porcelains []models.UpdateResponseMetadata) []*UpdateResponseMetadata {
    var items []*UpdateResponseMetadata
    for _, porcelain := range porcelains {
        items = append(items, UpdateResponseMetadataToPlumbing(porcelain))
    }
    return items
}

func RepeatedUpdateResponseMetadataToPorcelain(plumbings []*UpdateResponseMetadata) []models.UpdateResponseMetadata {
    var items []models.UpdateResponseMetadata
    for _, plumbing := range plumbings {
        items = append(items, UpdateResponseMetadataToPorcelain(plumbing))
    }
    return items
}

func DeleteResponseMetadataToPorcelain(plumbing *DeleteResponseMetadata) models.DeleteResponseMetadata {
    porcelain := models.DeleteResponseMetadata{}
    return porcelain
}

func DeleteResponseMetadataToPlumbing(porcelain models.DeleteResponseMetadata) *DeleteResponseMetadata {
    plumbing := &DeleteResponseMetadata{}
    return plumbing
}

func RepeatedDeleteResponseMetadataToPlumbing(porcelains []models.DeleteResponseMetadata) []*DeleteResponseMetadata {
    var items []*DeleteResponseMetadata
    for _, porcelain := range porcelains {
        items = append(items, DeleteResponseMetadataToPlumbing(porcelain))
    }
    return items
}

func RepeatedDeleteResponseMetadataToPorcelain(plumbings []*DeleteResponseMetadata) []models.DeleteResponseMetadata {
    var items []models.DeleteResponseMetadata
    for _, plumbing := range plumbings {
        items = append(items, DeleteResponseMetadataToPorcelain(plumbing))
    }
    return items
}

func NodeCreateResponseToPorcelain(plumbing *NodeCreateResponse) models.NodeCreateResponse {
    porcelain := models.NodeCreateResponse{}
    porcelain.Meta = CreateResponseMetadataToPorcelain(plumbing.Meta)
    porcelain.Node = NodeToPorcelain(plumbing.Node)
    porcelain.Token = TokenToPorcelain(plumbing.Token)
    return porcelain
}

func NodeCreateResponseToPlumbing(porcelain models.NodeCreateResponse) *NodeCreateResponse {
    plumbing := &NodeCreateResponse{}
    plumbing.Meta = CreateResponseMetadataToPlumbing(porcelain.Meta)
    plumbing.Node = NodeToPlumbing(porcelain.Node)
    plumbing.Token = TokenToPlumbing(porcelain.Token)
    return plumbing
}

func RepeatedNodeCreateResponseToPlumbing(porcelains []models.NodeCreateResponse) []*NodeCreateResponse {
    var items []*NodeCreateResponse
    for _, porcelain := range porcelains {
        items = append(items, NodeCreateResponseToPlumbing(porcelain))
    }
    return items
}

func RepeatedNodeCreateResponseToPorcelain(plumbings []*NodeCreateResponse) []models.NodeCreateResponse {
    var items []models.NodeCreateResponse
    for _, plumbing := range plumbings {
        items = append(items, NodeCreateResponseToPorcelain(plumbing))
    }
    return items
}

func NodeGetResponseToPorcelain(plumbing *NodeGetResponse) models.NodeGetResponse {
    porcelain := models.NodeGetResponse{}
    porcelain.Meta = GetResponseMetadataToPorcelain(plumbing.Meta)
    porcelain.Node = NodeToPorcelain(plumbing.Node)
    return porcelain
}

func NodeGetResponseToPlumbing(porcelain models.NodeGetResponse) *NodeGetResponse {
    plumbing := &NodeGetResponse{}
    plumbing.Meta = GetResponseMetadataToPlumbing(porcelain.Meta)
    plumbing.Node = NodeToPlumbing(porcelain.Node)
    return plumbing
}

func RepeatedNodeGetResponseToPlumbing(porcelains []models.NodeGetResponse) []*NodeGetResponse {
    var items []*NodeGetResponse
    for _, porcelain := range porcelains {
        items = append(items, NodeGetResponseToPlumbing(porcelain))
    }
    return items
}

func RepeatedNodeGetResponseToPorcelain(plumbings []*NodeGetResponse) []models.NodeGetResponse {
    var items []models.NodeGetResponse
    for _, plumbing := range plumbings {
        items = append(items, NodeGetResponseToPorcelain(plumbing))
    }
    return items
}

func NodeUpdateResponseToPorcelain(plumbing *NodeUpdateResponse) models.NodeUpdateResponse {
    porcelain := models.NodeUpdateResponse{}
    porcelain.Meta = UpdateResponseMetadataToPorcelain(plumbing.Meta)
    porcelain.Node = NodeToPorcelain(plumbing.Node)
    return porcelain
}

func NodeUpdateResponseToPlumbing(porcelain models.NodeUpdateResponse) *NodeUpdateResponse {
    plumbing := &NodeUpdateResponse{}
    plumbing.Meta = UpdateResponseMetadataToPlumbing(porcelain.Meta)
    plumbing.Node = NodeToPlumbing(porcelain.Node)
    return plumbing
}

func RepeatedNodeUpdateResponseToPlumbing(porcelains []models.NodeUpdateResponse) []*NodeUpdateResponse {
    var items []*NodeUpdateResponse
    for _, porcelain := range porcelains {
        items = append(items, NodeUpdateResponseToPlumbing(porcelain))
    }
    return items
}

func RepeatedNodeUpdateResponseToPorcelain(plumbings []*NodeUpdateResponse) []models.NodeUpdateResponse {
    var items []models.NodeUpdateResponse
    for _, plumbing := range plumbings {
        items = append(items, NodeUpdateResponseToPorcelain(plumbing))
    }
    return items
}

func NodeDeleteResponseToPorcelain(plumbing *NodeDeleteResponse) models.NodeDeleteResponse {
    porcelain := models.NodeDeleteResponse{}
    porcelain.Meta = DeleteResponseMetadataToPorcelain(plumbing.Meta)
    return porcelain
}

func NodeDeleteResponseToPlumbing(porcelain models.NodeDeleteResponse) *NodeDeleteResponse {
    plumbing := &NodeDeleteResponse{}
    plumbing.Meta = DeleteResponseMetadataToPlumbing(porcelain.Meta)
    return plumbing
}

func RepeatedNodeDeleteResponseToPlumbing(porcelains []models.NodeDeleteResponse) []*NodeDeleteResponse {
    var items []*NodeDeleteResponse
    for _, porcelain := range porcelains {
        items = append(items, NodeDeleteResponseToPlumbing(porcelain))
    }
    return items
}

func RepeatedNodeDeleteResponseToPorcelain(plumbings []*NodeDeleteResponse) []models.NodeDeleteResponse {
    var items []models.NodeDeleteResponse
    for _, plumbing := range plumbings {
        items = append(items, NodeDeleteResponseToPorcelain(plumbing))
    }
    return items
}

func NodeToPlumbing(porcelain models.Node) *Node {
    plumbing := &Node{}

    switch v := porcelain.(type) {
    case *models.Relay:
        plumbing.Node = &Node_Relay{ Relay: RelayToPlumbing(*v) }
    case *models.Gateway:
        plumbing.Node = &Node_Gateway{ Gateway: GatewayToPlumbing(*v) }
    }
    return plumbing
}

func NodeToPorcelain(plumbing *Node) models.Node {
    if plumbing.GetRelay() != nil {
        v := RelayToPorcelain(plumbing.GetRelay())
        return &v
    }
    if plumbing.GetGateway() != nil {
        v := GatewayToPorcelain(plumbing.GetGateway())
        return &v
    }
    return nil
}

func RepeatedNodeToPlumbing(porcelains []models.Node) []*Node {
    var items []*Node
    for _, porcelain := range porcelains {
        items = append(items, NodeToPlumbing(porcelain))
    }
    return items
}

func RepeatedNodeToPorcelain(plumbings []*Node) []models.Node {
    var items []models.Node
    for _, plumbing := range plumbings {
        items = append(items, NodeToPorcelain(plumbing))
    }
    return items
}

func RelayToPorcelain(plumbing *Relay) models.Relay {
    porcelain := models.Relay{}
    porcelain.ID = plumbing.Id
    porcelain.Name = plumbing.Name
    return porcelain
}

func RelayToPlumbing(porcelain models.Relay) *Relay {
    plumbing := &Relay{}
    plumbing.Id = porcelain.ID
    plumbing.Name = porcelain.Name
    return plumbing
}

func RepeatedRelayToPlumbing(porcelains []models.Relay) []*Relay {
    var items []*Relay
    for _, porcelain := range porcelains {
        items = append(items, RelayToPlumbing(porcelain))
    }
    return items
}

func RepeatedRelayToPorcelain(plumbings []*Relay) []models.Relay {
    var items []models.Relay
    for _, plumbing := range plumbings {
        items = append(items, RelayToPorcelain(plumbing))
    }
    return items
}

func GatewayToPorcelain(plumbing *Gateway) models.Gateway {
    porcelain := models.Gateway{}
    porcelain.ID = plumbing.Id
    porcelain.Name = plumbing.Name
    porcelain.ListenAddress = plumbing.ListenAddress
    porcelain.BindAddress = plumbing.BindAddress
    return porcelain
}

func GatewayToPlumbing(porcelain models.Gateway) *Gateway {
    plumbing := &Gateway{}
    plumbing.Id = porcelain.ID
    plumbing.Name = porcelain.Name
    plumbing.ListenAddress = porcelain.ListenAddress
    plumbing.BindAddress = porcelain.BindAddress
    return plumbing
}

func RepeatedGatewayToPlumbing(porcelains []models.Gateway) []*Gateway {
    var items []*Gateway
    for _, porcelain := range porcelains {
        items = append(items, GatewayToPlumbing(porcelain))
    }
    return items
}

func RepeatedGatewayToPorcelain(plumbings []*Gateway) []models.Gateway {
    var items []models.Gateway
    for _, plumbing := range plumbings {
        items = append(items, GatewayToPorcelain(plumbing))
    }
    return items
}

func TokenToPorcelain(plumbing *Token) models.Token {
    porcelain := models.Token{}
    porcelain.ID = plumbing.Id
    porcelain.Token = plumbing.Token
    return porcelain
}

func TokenToPlumbing(porcelain models.Token) *Token {
    plumbing := &Token{}
    plumbing.Id = porcelain.ID
    plumbing.Token = porcelain.Token
    return plumbing
}

func RepeatedTokenToPlumbing(porcelains []models.Token) []*Token {
    var items []*Token
    for _, porcelain := range porcelains {
        items = append(items, TokenToPlumbing(porcelain))
    }
    return items
}

func RepeatedTokenToPorcelain(plumbings []*Token) []models.Token {
    var items []models.Token
    for _, plumbing := range plumbings {
        items = append(items, TokenToPorcelain(plumbing))
    }
    return items
}

func RoleCreateResponseToPorcelain(plumbing *RoleCreateResponse) models.RoleCreateResponse {
    porcelain := models.RoleCreateResponse{}
    porcelain.Meta = CreateResponseMetadataToPorcelain(plumbing.Meta)
    porcelain.Role = RoleToPorcelain(plumbing.Role)
    return porcelain
}

func RoleCreateResponseToPlumbing(porcelain models.RoleCreateResponse) *RoleCreateResponse {
    plumbing := &RoleCreateResponse{}
    plumbing.Meta = CreateResponseMetadataToPlumbing(porcelain.Meta)
    plumbing.Role = RoleToPlumbing(porcelain.Role)
    return plumbing
}

func RepeatedRoleCreateResponseToPlumbing(porcelains []models.RoleCreateResponse) []*RoleCreateResponse {
    var items []*RoleCreateResponse
    for _, porcelain := range porcelains {
        items = append(items, RoleCreateResponseToPlumbing(porcelain))
    }
    return items
}

func RepeatedRoleCreateResponseToPorcelain(plumbings []*RoleCreateResponse) []models.RoleCreateResponse {
    var items []models.RoleCreateResponse
    for _, plumbing := range plumbings {
        items = append(items, RoleCreateResponseToPorcelain(plumbing))
    }
    return items
}

func RoleGetResponseToPorcelain(plumbing *RoleGetResponse) models.RoleGetResponse {
    porcelain := models.RoleGetResponse{}
    porcelain.Meta = GetResponseMetadataToPorcelain(plumbing.Meta)
    porcelain.Role = RoleToPorcelain(plumbing.Role)
    return porcelain
}

func RoleGetResponseToPlumbing(porcelain models.RoleGetResponse) *RoleGetResponse {
    plumbing := &RoleGetResponse{}
    plumbing.Meta = GetResponseMetadataToPlumbing(porcelain.Meta)
    plumbing.Role = RoleToPlumbing(porcelain.Role)
    return plumbing
}

func RepeatedRoleGetResponseToPlumbing(porcelains []models.RoleGetResponse) []*RoleGetResponse {
    var items []*RoleGetResponse
    for _, porcelain := range porcelains {
        items = append(items, RoleGetResponseToPlumbing(porcelain))
    }
    return items
}

func RepeatedRoleGetResponseToPorcelain(plumbings []*RoleGetResponse) []models.RoleGetResponse {
    var items []models.RoleGetResponse
    for _, plumbing := range plumbings {
        items = append(items, RoleGetResponseToPorcelain(plumbing))
    }
    return items
}

func RoleUpdateResponseToPorcelain(plumbing *RoleUpdateResponse) models.RoleUpdateResponse {
    porcelain := models.RoleUpdateResponse{}
    porcelain.Meta = UpdateResponseMetadataToPorcelain(plumbing.Meta)
    porcelain.Role = RoleToPorcelain(plumbing.Role)
    return porcelain
}

func RoleUpdateResponseToPlumbing(porcelain models.RoleUpdateResponse) *RoleUpdateResponse {
    plumbing := &RoleUpdateResponse{}
    plumbing.Meta = UpdateResponseMetadataToPlumbing(porcelain.Meta)
    plumbing.Role = RoleToPlumbing(porcelain.Role)
    return plumbing
}

func RepeatedRoleUpdateResponseToPlumbing(porcelains []models.RoleUpdateResponse) []*RoleUpdateResponse {
    var items []*RoleUpdateResponse
    for _, porcelain := range porcelains {
        items = append(items, RoleUpdateResponseToPlumbing(porcelain))
    }
    return items
}

func RepeatedRoleUpdateResponseToPorcelain(plumbings []*RoleUpdateResponse) []models.RoleUpdateResponse {
    var items []models.RoleUpdateResponse
    for _, plumbing := range plumbings {
        items = append(items, RoleUpdateResponseToPorcelain(plumbing))
    }
    return items
}

func RoleDeleteResponseToPorcelain(plumbing *RoleDeleteResponse) models.RoleDeleteResponse {
    porcelain := models.RoleDeleteResponse{}
    porcelain.Meta = DeleteResponseMetadataToPorcelain(plumbing.Meta)
    return porcelain
}

func RoleDeleteResponseToPlumbing(porcelain models.RoleDeleteResponse) *RoleDeleteResponse {
    plumbing := &RoleDeleteResponse{}
    plumbing.Meta = DeleteResponseMetadataToPlumbing(porcelain.Meta)
    return plumbing
}

func RepeatedRoleDeleteResponseToPlumbing(porcelains []models.RoleDeleteResponse) []*RoleDeleteResponse {
    var items []*RoleDeleteResponse
    for _, porcelain := range porcelains {
        items = append(items, RoleDeleteResponseToPlumbing(porcelain))
    }
    return items
}

func RepeatedRoleDeleteResponseToPorcelain(plumbings []*RoleDeleteResponse) []models.RoleDeleteResponse {
    var items []models.RoleDeleteResponse
    for _, plumbing := range plumbings {
        items = append(items, RoleDeleteResponseToPorcelain(plumbing))
    }
    return items
}

func RoleToPorcelain(plumbing *Role) models.Role {
    porcelain := models.Role{}
    porcelain.ID = plumbing.Id
    porcelain.Name = plumbing.Name
    porcelain.Composite = plumbing.Composite
    porcelain.Roles = RepeatedRoleToPorcelain(plumbing.Roles)
    return porcelain
}

func RoleToPlumbing(porcelain models.Role) *Role {
    plumbing := &Role{}
    plumbing.Id = porcelain.ID
    plumbing.Name = porcelain.Name
    plumbing.Composite = porcelain.Composite
    plumbing.Roles = RepeatedRoleToPlumbing(porcelain.Roles)
    return plumbing
}

func RepeatedRoleToPlumbing(porcelains []models.Role) []*Role {
    var items []*Role
    for _, porcelain := range porcelains {
        items = append(items, RoleToPlumbing(porcelain))
    }
    return items
}

func RepeatedRoleToPorcelain(plumbings []*Role) []models.Role {
    var items []models.Role
    for _, plumbing := range plumbings {
        items = append(items, RoleToPorcelain(plumbing))
    }
    return items
}

func ErrorToPorcelain(err error) error {
	if s, ok := status.FromError(err); ok {
		for _, details := range s.Details() {
			switch d := details.(type) {

			// AlreadyExistsError is used when an entity already exists in the system
			case *AlreadyExistsError:
				e := &errors.AlreadyExistsError{}
				e.Message = s.Message()
				e.Entities = d.Entities
				return e

			// NotFoundError is used when an entity does not exist in the system
			case *NotFoundError:
				e := &errors.NotFoundError{}
				e.Message = s.Message()
				e.Entities = d.Entities
				return e

			// BadRequestError identifies a bad request sent by the client
			case *BadRequestError:
				e := &errors.BadRequestError{}
				e.Message = s.Message()
				return e

			// AuthenticationError is used to specify an authentication failure condition
			case *AuthenticationError:
				e := &errors.AuthenticationError{}
				e.Message = s.Message()
				return e

			// PermissionError is used to specify a permissions violation
			case *PermissionError:
				e := &errors.PermissionError{}
				e.Message = s.Message()
				e.Permission = d.Permission
				e.Entities = d.Entities
				return e

			// InternalError is used to specify an internal system error
			case *InternalError:
				e := &errors.InternalError{}
				e.Message = s.Message()
				return e

			// RateLimitError is used for rate limit excess condition
			case *RateLimitError:
				e := &errors.RateLimitError{}
				e.Message = s.Message()
				return e

			}
		}
	}
	return &errors.RPCError{Wrapped: err}
}

type NodeIteratorImplFetchFunc func() ([]models.Node, bool, error)
type NodeIteratorImpl struct {
	buffer []models.Node
    index int
    hasNextPage bool
    err error
    fetch NodeIteratorImplFetchFunc
}

func NewNodeIteratorImpl(f NodeIteratorImplFetchFunc) *NodeIteratorImpl {
    return &NodeIteratorImpl{
        hasNextPage: true,
        fetch: f,
    }
}

func (n *NodeIteratorImpl) Next() bool {
    if n.index < len(n.buffer) - 1 {
        n.index++
        return true
    }

    // reached end of buffer
    if !n.hasNextPage {
        return false
    }

    n.index = 0
    n.buffer, n.hasNextPage, n.err = n.fetch()
    return len(n.buffer) > 0
}

func (n *NodeIteratorImpl) Value() models.Node {
    return n.buffer[n.index]
}

func (n *NodeIteratorImpl) Err() error {
    return n.err
}
type RoleIteratorImplFetchFunc func() ([]models.Role, bool, error)
type RoleIteratorImpl struct {
	buffer []models.Role
    index int
    hasNextPage bool
    err error
    fetch RoleIteratorImplFetchFunc
}

func NewRoleIteratorImpl(f RoleIteratorImplFetchFunc) *RoleIteratorImpl {
    return &RoleIteratorImpl{
        hasNextPage: true,
        fetch: f,
    }
}

func (r *RoleIteratorImpl) Next() bool {
    if r.index < len(r.buffer) - 1 {
        r.index++
        return true
    }

    // reached end of buffer
    if !r.hasNextPage {
        return false
    }

    r.index = 0
    r.buffer, r.hasNextPage, r.err = r.fetch()
    return len(r.buffer) > 0
}

func (r *RoleIteratorImpl) Value() models.Role {
    return r.buffer[r.index]
}

func (r *RoleIteratorImpl) Err() error {
    return r.err
}

func CreateGRPCContext(ctx context.Context, token string) context.Context {
    return metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		"authorization": token,
	}))
}