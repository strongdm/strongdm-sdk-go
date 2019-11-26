package sdm

import (
    "context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
    "google.golang.org/grpc/metadata"
	proto "github.com/strongdm/strongdm-sdk-go/internal/v1"
)

func createResponseMetadataToPorcelain(plumbing *proto.CreateResponseMetadata) *CreateResponseMetadata {
    if plumbing == nil {
        return nil
    }
    porcelain := &CreateResponseMetadata{}
    return porcelain
}

func createResponseMetadataToPlumbing(porcelain *CreateResponseMetadata) *proto.CreateResponseMetadata {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.CreateResponseMetadata{}
    return plumbing
}

func repeatedCreateResponseMetadataToPlumbing(porcelains []*CreateResponseMetadata) []*proto.CreateResponseMetadata {
    var items []*proto.CreateResponseMetadata
    for _, porcelain := range porcelains {
        items = append(items, createResponseMetadataToPlumbing(porcelain))
    }
    return items
}

func repeatedCreateResponseMetadataToPorcelain(plumbings []*proto.CreateResponseMetadata) []*CreateResponseMetadata {
    var items []*CreateResponseMetadata
    for _, plumbing := range plumbings {
        items = append(items, createResponseMetadataToPorcelain(plumbing))
    }
    return items
}

func getResponseMetadataToPorcelain(plumbing *proto.GetResponseMetadata) *GetResponseMetadata {
    if plumbing == nil {
        return nil
    }
    porcelain := &GetResponseMetadata{}
    return porcelain
}

func getResponseMetadataToPlumbing(porcelain *GetResponseMetadata) *proto.GetResponseMetadata {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.GetResponseMetadata{}
    return plumbing
}

func repeatedGetResponseMetadataToPlumbing(porcelains []*GetResponseMetadata) []*proto.GetResponseMetadata {
    var items []*proto.GetResponseMetadata
    for _, porcelain := range porcelains {
        items = append(items, getResponseMetadataToPlumbing(porcelain))
    }
    return items
}

func repeatedGetResponseMetadataToPorcelain(plumbings []*proto.GetResponseMetadata) []*GetResponseMetadata {
    var items []*GetResponseMetadata
    for _, plumbing := range plumbings {
        items = append(items, getResponseMetadataToPorcelain(plumbing))
    }
    return items
}

func updateResponseMetadataToPorcelain(plumbing *proto.UpdateResponseMetadata) *UpdateResponseMetadata {
    if plumbing == nil {
        return nil
    }
    porcelain := &UpdateResponseMetadata{}
    return porcelain
}

func updateResponseMetadataToPlumbing(porcelain *UpdateResponseMetadata) *proto.UpdateResponseMetadata {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.UpdateResponseMetadata{}
    return plumbing
}

func repeatedUpdateResponseMetadataToPlumbing(porcelains []*UpdateResponseMetadata) []*proto.UpdateResponseMetadata {
    var items []*proto.UpdateResponseMetadata
    for _, porcelain := range porcelains {
        items = append(items, updateResponseMetadataToPlumbing(porcelain))
    }
    return items
}

func repeatedUpdateResponseMetadataToPorcelain(plumbings []*proto.UpdateResponseMetadata) []*UpdateResponseMetadata {
    var items []*UpdateResponseMetadata
    for _, plumbing := range plumbings {
        items = append(items, updateResponseMetadataToPorcelain(plumbing))
    }
    return items
}

func deleteResponseMetadataToPorcelain(plumbing *proto.DeleteResponseMetadata) *DeleteResponseMetadata {
    if plumbing == nil {
        return nil
    }
    porcelain := &DeleteResponseMetadata{}
    return porcelain
}

func deleteResponseMetadataToPlumbing(porcelain *DeleteResponseMetadata) *proto.DeleteResponseMetadata {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.DeleteResponseMetadata{}
    return plumbing
}

func repeatedDeleteResponseMetadataToPlumbing(porcelains []*DeleteResponseMetadata) []*proto.DeleteResponseMetadata {
    var items []*proto.DeleteResponseMetadata
    for _, porcelain := range porcelains {
        items = append(items, deleteResponseMetadataToPlumbing(porcelain))
    }
    return items
}

func repeatedDeleteResponseMetadataToPorcelain(plumbings []*proto.DeleteResponseMetadata) []*DeleteResponseMetadata {
    var items []*DeleteResponseMetadata
    for _, plumbing := range plumbings {
        items = append(items, deleteResponseMetadataToPorcelain(plumbing))
    }
    return items
}

func nodeCreateResponseToPorcelain(plumbing *proto.NodeCreateResponse) *NodeCreateResponse {
    if plumbing == nil {
        return nil
    }
    porcelain := &NodeCreateResponse{}
    porcelain.Meta = createResponseMetadataToPorcelain(plumbing.Meta)
    porcelain.Node = nodeToPorcelain(plumbing.Node)
    porcelain.Token = plumbing.Token
    return porcelain
}

func nodeCreateResponseToPlumbing(porcelain *NodeCreateResponse) *proto.NodeCreateResponse {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.NodeCreateResponse{}
    plumbing.Meta = createResponseMetadataToPlumbing(porcelain.Meta)
    plumbing.Node = nodeToPlumbing(porcelain.Node)
    plumbing.Token = porcelain.Token
    return plumbing
}

func repeatedNodeCreateResponseToPlumbing(porcelains []*NodeCreateResponse) []*proto.NodeCreateResponse {
    var items []*proto.NodeCreateResponse
    for _, porcelain := range porcelains {
        items = append(items, nodeCreateResponseToPlumbing(porcelain))
    }
    return items
}

func repeatedNodeCreateResponseToPorcelain(plumbings []*proto.NodeCreateResponse) []*NodeCreateResponse {
    var items []*NodeCreateResponse
    for _, plumbing := range plumbings {
        items = append(items, nodeCreateResponseToPorcelain(plumbing))
    }
    return items
}

func nodeGetResponseToPorcelain(plumbing *proto.NodeGetResponse) *NodeGetResponse {
    if plumbing == nil {
        return nil
    }
    porcelain := &NodeGetResponse{}
    porcelain.Meta = getResponseMetadataToPorcelain(plumbing.Meta)
    porcelain.Node = nodeToPorcelain(plumbing.Node)
    return porcelain
}

func nodeGetResponseToPlumbing(porcelain *NodeGetResponse) *proto.NodeGetResponse {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.NodeGetResponse{}
    plumbing.Meta = getResponseMetadataToPlumbing(porcelain.Meta)
    plumbing.Node = nodeToPlumbing(porcelain.Node)
    return plumbing
}

func repeatedNodeGetResponseToPlumbing(porcelains []*NodeGetResponse) []*proto.NodeGetResponse {
    var items []*proto.NodeGetResponse
    for _, porcelain := range porcelains {
        items = append(items, nodeGetResponseToPlumbing(porcelain))
    }
    return items
}

func repeatedNodeGetResponseToPorcelain(plumbings []*proto.NodeGetResponse) []*NodeGetResponse {
    var items []*NodeGetResponse
    for _, plumbing := range plumbings {
        items = append(items, nodeGetResponseToPorcelain(plumbing))
    }
    return items
}

func nodeUpdateResponseToPorcelain(plumbing *proto.NodeUpdateResponse) *NodeUpdateResponse {
    if plumbing == nil {
        return nil
    }
    porcelain := &NodeUpdateResponse{}
    porcelain.Meta = updateResponseMetadataToPorcelain(plumbing.Meta)
    porcelain.Node = nodeToPorcelain(plumbing.Node)
    return porcelain
}

func nodeUpdateResponseToPlumbing(porcelain *NodeUpdateResponse) *proto.NodeUpdateResponse {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.NodeUpdateResponse{}
    plumbing.Meta = updateResponseMetadataToPlumbing(porcelain.Meta)
    plumbing.Node = nodeToPlumbing(porcelain.Node)
    return plumbing
}

func repeatedNodeUpdateResponseToPlumbing(porcelains []*NodeUpdateResponse) []*proto.NodeUpdateResponse {
    var items []*proto.NodeUpdateResponse
    for _, porcelain := range porcelains {
        items = append(items, nodeUpdateResponseToPlumbing(porcelain))
    }
    return items
}

func repeatedNodeUpdateResponseToPorcelain(plumbings []*proto.NodeUpdateResponse) []*NodeUpdateResponse {
    var items []*NodeUpdateResponse
    for _, plumbing := range plumbings {
        items = append(items, nodeUpdateResponseToPorcelain(plumbing))
    }
    return items
}

func nodeDeleteResponseToPorcelain(plumbing *proto.NodeDeleteResponse) *NodeDeleteResponse {
    if plumbing == nil {
        return nil
    }
    porcelain := &NodeDeleteResponse{}
    porcelain.Meta = deleteResponseMetadataToPorcelain(plumbing.Meta)
    return porcelain
}

func nodeDeleteResponseToPlumbing(porcelain *NodeDeleteResponse) *proto.NodeDeleteResponse {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.NodeDeleteResponse{}
    plumbing.Meta = deleteResponseMetadataToPlumbing(porcelain.Meta)
    return plumbing
}

func repeatedNodeDeleteResponseToPlumbing(porcelains []*NodeDeleteResponse) []*proto.NodeDeleteResponse {
    var items []*proto.NodeDeleteResponse
    for _, porcelain := range porcelains {
        items = append(items, nodeDeleteResponseToPlumbing(porcelain))
    }
    return items
}

func repeatedNodeDeleteResponseToPorcelain(plumbings []*proto.NodeDeleteResponse) []*NodeDeleteResponse {
    var items []*NodeDeleteResponse
    for _, plumbing := range plumbings {
        items = append(items, nodeDeleteResponseToPorcelain(plumbing))
    }
    return items
}

func nodeToPlumbing(porcelain Node) *proto.Node {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.Node{}

    switch v := porcelain.(type) {
    case *Relay:
        plumbing.Node = &proto.Node_Relay{ Relay: relayToPlumbing(v) }
    case *Gateway:
        plumbing.Node = &proto.Node_Gateway{ Gateway: gatewayToPlumbing(v) }
    }
    return plumbing
}

func nodeToPorcelain(plumbing *proto.Node) Node {
    if plumbing.GetRelay() != nil {
        return relayToPorcelain(plumbing.GetRelay())
    }
    if plumbing.GetGateway() != nil {
        return gatewayToPorcelain(plumbing.GetGateway())
    }
    return nil
}

func repeatedNodeToPlumbing(porcelains []Node) []*proto.Node {
    var items []*proto.Node
    for _, porcelain := range porcelains {
        items = append(items, nodeToPlumbing(porcelain))
    }
    return items
}

func repeatedNodeToPorcelain(plumbings []*proto.Node) []Node {
    var items []Node
    for _, plumbing := range plumbings {
        items = append(items, nodeToPorcelain(plumbing))
    }
    return items
}

func relayToPorcelain(plumbing *proto.Relay) *Relay {
    if plumbing == nil {
        return nil
    }
    porcelain := &Relay{}
    porcelain.ID = plumbing.Id
    porcelain.Name = plumbing.Name
    porcelain.State = plumbing.State
    return porcelain
}

func relayToPlumbing(porcelain *Relay) *proto.Relay {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.Relay{}
    plumbing.Id = porcelain.ID
    plumbing.Name = porcelain.Name
    plumbing.State = porcelain.State
    return plumbing
}

func repeatedRelayToPlumbing(porcelains []*Relay) []*proto.Relay {
    var items []*proto.Relay
    for _, porcelain := range porcelains {
        items = append(items, relayToPlumbing(porcelain))
    }
    return items
}

func repeatedRelayToPorcelain(plumbings []*proto.Relay) []*Relay {
    var items []*Relay
    for _, plumbing := range plumbings {
        items = append(items, relayToPorcelain(plumbing))
    }
    return items
}

func gatewayToPorcelain(plumbing *proto.Gateway) *Gateway {
    if plumbing == nil {
        return nil
    }
    porcelain := &Gateway{}
    porcelain.ID = plumbing.Id
    porcelain.Name = plumbing.Name
    porcelain.State = plumbing.State
    porcelain.ListenAddress = plumbing.ListenAddress
    porcelain.BindAddress = plumbing.BindAddress
    return porcelain
}

func gatewayToPlumbing(porcelain *Gateway) *proto.Gateway {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.Gateway{}
    plumbing.Id = porcelain.ID
    plumbing.Name = porcelain.Name
    plumbing.State = porcelain.State
    plumbing.ListenAddress = porcelain.ListenAddress
    plumbing.BindAddress = porcelain.BindAddress
    return plumbing
}

func repeatedGatewayToPlumbing(porcelains []*Gateway) []*proto.Gateway {
    var items []*proto.Gateway
    for _, porcelain := range porcelains {
        items = append(items, gatewayToPlumbing(porcelain))
    }
    return items
}

func repeatedGatewayToPorcelain(plumbings []*proto.Gateway) []*Gateway {
    var items []*Gateway
    for _, plumbing := range plumbings {
        items = append(items, gatewayToPorcelain(plumbing))
    }
    return items
}

func roleCreateResponseToPorcelain(plumbing *proto.RoleCreateResponse) *RoleCreateResponse {
    if plumbing == nil {
        return nil
    }
    porcelain := &RoleCreateResponse{}
    porcelain.Meta = createResponseMetadataToPorcelain(plumbing.Meta)
    porcelain.Role = roleToPorcelain(plumbing.Role)
    return porcelain
}

func roleCreateResponseToPlumbing(porcelain *RoleCreateResponse) *proto.RoleCreateResponse {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.RoleCreateResponse{}
    plumbing.Meta = createResponseMetadataToPlumbing(porcelain.Meta)
    plumbing.Role = roleToPlumbing(porcelain.Role)
    return plumbing
}

func repeatedRoleCreateResponseToPlumbing(porcelains []*RoleCreateResponse) []*proto.RoleCreateResponse {
    var items []*proto.RoleCreateResponse
    for _, porcelain := range porcelains {
        items = append(items, roleCreateResponseToPlumbing(porcelain))
    }
    return items
}

func repeatedRoleCreateResponseToPorcelain(plumbings []*proto.RoleCreateResponse) []*RoleCreateResponse {
    var items []*RoleCreateResponse
    for _, plumbing := range plumbings {
        items = append(items, roleCreateResponseToPorcelain(plumbing))
    }
    return items
}

func roleGetResponseToPorcelain(plumbing *proto.RoleGetResponse) *RoleGetResponse {
    if plumbing == nil {
        return nil
    }
    porcelain := &RoleGetResponse{}
    porcelain.Meta = getResponseMetadataToPorcelain(plumbing.Meta)
    porcelain.Role = roleToPorcelain(plumbing.Role)
    return porcelain
}

func roleGetResponseToPlumbing(porcelain *RoleGetResponse) *proto.RoleGetResponse {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.RoleGetResponse{}
    plumbing.Meta = getResponseMetadataToPlumbing(porcelain.Meta)
    plumbing.Role = roleToPlumbing(porcelain.Role)
    return plumbing
}

func repeatedRoleGetResponseToPlumbing(porcelains []*RoleGetResponse) []*proto.RoleGetResponse {
    var items []*proto.RoleGetResponse
    for _, porcelain := range porcelains {
        items = append(items, roleGetResponseToPlumbing(porcelain))
    }
    return items
}

func repeatedRoleGetResponseToPorcelain(plumbings []*proto.RoleGetResponse) []*RoleGetResponse {
    var items []*RoleGetResponse
    for _, plumbing := range plumbings {
        items = append(items, roleGetResponseToPorcelain(plumbing))
    }
    return items
}

func roleUpdateResponseToPorcelain(plumbing *proto.RoleUpdateResponse) *RoleUpdateResponse {
    if plumbing == nil {
        return nil
    }
    porcelain := &RoleUpdateResponse{}
    porcelain.Meta = updateResponseMetadataToPorcelain(plumbing.Meta)
    porcelain.Role = roleToPorcelain(plumbing.Role)
    return porcelain
}

func roleUpdateResponseToPlumbing(porcelain *RoleUpdateResponse) *proto.RoleUpdateResponse {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.RoleUpdateResponse{}
    plumbing.Meta = updateResponseMetadataToPlumbing(porcelain.Meta)
    plumbing.Role = roleToPlumbing(porcelain.Role)
    return plumbing
}

func repeatedRoleUpdateResponseToPlumbing(porcelains []*RoleUpdateResponse) []*proto.RoleUpdateResponse {
    var items []*proto.RoleUpdateResponse
    for _, porcelain := range porcelains {
        items = append(items, roleUpdateResponseToPlumbing(porcelain))
    }
    return items
}

func repeatedRoleUpdateResponseToPorcelain(plumbings []*proto.RoleUpdateResponse) []*RoleUpdateResponse {
    var items []*RoleUpdateResponse
    for _, plumbing := range plumbings {
        items = append(items, roleUpdateResponseToPorcelain(plumbing))
    }
    return items
}

func roleDeleteResponseToPorcelain(plumbing *proto.RoleDeleteResponse) *RoleDeleteResponse {
    if plumbing == nil {
        return nil
    }
    porcelain := &RoleDeleteResponse{}
    porcelain.Meta = deleteResponseMetadataToPorcelain(plumbing.Meta)
    return porcelain
}

func roleDeleteResponseToPlumbing(porcelain *RoleDeleteResponse) *proto.RoleDeleteResponse {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.RoleDeleteResponse{}
    plumbing.Meta = deleteResponseMetadataToPlumbing(porcelain.Meta)
    return plumbing
}

func repeatedRoleDeleteResponseToPlumbing(porcelains []*RoleDeleteResponse) []*proto.RoleDeleteResponse {
    var items []*proto.RoleDeleteResponse
    for _, porcelain := range porcelains {
        items = append(items, roleDeleteResponseToPlumbing(porcelain))
    }
    return items
}

func repeatedRoleDeleteResponseToPorcelain(plumbings []*proto.RoleDeleteResponse) []*RoleDeleteResponse {
    var items []*RoleDeleteResponse
    for _, plumbing := range plumbings {
        items = append(items, roleDeleteResponseToPorcelain(plumbing))
    }
    return items
}

func roleToPorcelain(plumbing *proto.Role) *Role {
    if plumbing == nil {
        return nil
    }
    porcelain := &Role{}
    porcelain.ID = plumbing.Id
    porcelain.Name = plumbing.Name
    porcelain.Composite = plumbing.Composite
    return porcelain
}

func roleToPlumbing(porcelain *Role) *proto.Role {
    if porcelain == nil {
        return nil
    }
    plumbing := &proto.Role{}
    plumbing.Id = porcelain.ID
    plumbing.Name = porcelain.Name
    plumbing.Composite = porcelain.Composite
    return plumbing
}

func repeatedRoleToPlumbing(porcelains []*Role) []*proto.Role {
    var items []*proto.Role
    for _, porcelain := range porcelains {
        items = append(items, roleToPlumbing(porcelain))
    }
    return items
}

func repeatedRoleToPorcelain(plumbings []*proto.Role) []*Role {
    var items []*Role
    for _, plumbing := range plumbings {
        items = append(items, roleToPorcelain(plumbing))
    }
    return items
}


type rpcError struct {
	wrapped error
	code    int
}

func (e *rpcError) Error() string {
	return e.wrapped.Error()
}

func (e *rpcError) Unwrap() error {
	return e.wrapped
}

func (e *rpcError) Code() int {
	return e.code
}

func errorToPorcelain(err error) error {
	if s, ok := status.FromError(err); ok {
		for _, details := range s.Details() {
			switch d := details.(type) {

			// AlreadyExistsError is used when an entity already exists in the system
			case *proto.AlreadyExistsError:
				e := &AlreadyExistsError{}
				e.Message = s.Message()
				e.Entity = d.Entity
				return e

			// NotFoundError is used when an entity does not exist in the system
			case *proto.NotFoundError:
				e := &NotFoundError{}
				e.Message = s.Message()
				e.Entity = d.Entity
				return e

			// BadRequestError identifies a bad request sent by the client
			case *proto.BadRequestError:
				e := &BadRequestError{}
				e.Message = s.Message()
				return e

			// AuthenticationError is used to specify an authentication failure condition
			case *proto.AuthenticationError:
				e := &AuthenticationError{}
				e.Message = s.Message()
				return e

			// PermissionError is used to specify a permissions violation
			case *proto.PermissionError:
				e := &PermissionError{}
				e.Message = s.Message()
				return e

			// InternalError is used to specify an internal system error
			case *proto.InternalError:
				e := &InternalError{}
				e.Message = s.Message()
				return e

			// RateLimitError is used for rate limit excess condition
			case *proto.RateLimitError:
				e := &RateLimitError{}
				e.Message = s.Message()
				return e

			}
		}
        if s.Code() == codes.Canceled {
            return &ContextCanceledError{Wrapped: err}
        }
        if s.Code() == codes.DeadlineExceeded {
            return &DeadlineExceededError{Wrapped: err}
        }
        return &rpcError{wrapped: err, code: int(s.Code())}
	}
	return &UnknownError{Wrapped: err}
}

type nodeIteratorImplFetchFunc func() ([]Node, bool, error)
type nodeIteratorImpl struct {
	buffer []Node
    index int
    hasNextPage bool
    err error
    fetch nodeIteratorImplFetchFunc
}

func newNodeIteratorImpl(f nodeIteratorImplFetchFunc) *nodeIteratorImpl {
    return &nodeIteratorImpl{
        hasNextPage: true,
        fetch: f,
    }
}

func (n *nodeIteratorImpl) Next() bool {
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

func (n *nodeIteratorImpl) Value() Node {
    if n.index >= len(n.buffer) {
        return nil
    }
    return n.buffer[n.index]
}

func (n *nodeIteratorImpl) Err() error {
    return n.err
}
type roleIteratorImplFetchFunc func() ([]*Role, bool, error)
type roleIteratorImpl struct {
	buffer []*Role
    index int
    hasNextPage bool
    err error
    fetch roleIteratorImplFetchFunc
}

func newRoleIteratorImpl(f roleIteratorImplFetchFunc) *roleIteratorImpl {
    return &roleIteratorImpl{
        hasNextPage: true,
        fetch: f,
    }
}

func (r *roleIteratorImpl) Next() bool {
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

func (r *roleIteratorImpl) Value() *Role {
    if r.index >= len(r.buffer) {
        return nil
    }
    return r.buffer[r.index]
}

func (r *roleIteratorImpl) Err() error {
    return r.err
}

func createGRPCContext(ctx context.Context, token string) context.Context {
    return metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		"authorization": token,
	}))
}