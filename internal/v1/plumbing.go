package v1

import (
	"github.com/strongdm/strongdm-sdk-go/errors"
	"github.com/strongdm/strongdm-sdk-go/models"
	"google.golang.org/grpc/status"
)

func CreateResponseMetadataToPorcelain(plumbing *CreateResponseMetadata) models.CreateResponseMetadata {
	porcelain := models.CreateResponseMetadata{}
	porcelain.Affected = plumbing.Affected
	return porcelain
}

func CreateResponseMetadataToPlumbing(porcelain models.CreateResponseMetadata) *CreateResponseMetadata {
	plumbing := &CreateResponseMetadata{}
	plumbing.Affected = porcelain.Affected
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
	porcelain.Found = plumbing.Found
	return porcelain
}

func GetResponseMetadataToPlumbing(porcelain models.GetResponseMetadata) *GetResponseMetadata {
	plumbing := &GetResponseMetadata{}
	plumbing.Found = porcelain.Found
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
	porcelain.Affected = plumbing.Affected
	return porcelain
}

func UpdateResponseMetadataToPlumbing(porcelain models.UpdateResponseMetadata) *UpdateResponseMetadata {
	plumbing := &UpdateResponseMetadata{}
	plumbing.Affected = porcelain.Affected
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
	porcelain.Affected = plumbing.Affected
	return porcelain
}

func DeleteResponseMetadataToPlumbing(porcelain models.DeleteResponseMetadata) *DeleteResponseMetadata {
	plumbing := &DeleteResponseMetadata{}
	plumbing.Affected = porcelain.Affected
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

func BatchUpdateResponseMetadataToPorcelain(plumbing *BatchUpdateResponseMetadata) models.BatchUpdateResponseMetadata {
	porcelain := models.BatchUpdateResponseMetadata{}
	porcelain.Found = plumbing.Found
	porcelain.Affected = plumbing.Affected
	return porcelain
}

func BatchUpdateResponseMetadataToPlumbing(porcelain models.BatchUpdateResponseMetadata) *BatchUpdateResponseMetadata {
	plumbing := &BatchUpdateResponseMetadata{}
	plumbing.Found = porcelain.Found
	plumbing.Affected = porcelain.Affected
	return plumbing
}

func RepeatedBatchUpdateResponseMetadataToPlumbing(porcelains []models.BatchUpdateResponseMetadata) []*BatchUpdateResponseMetadata {
	var items []*BatchUpdateResponseMetadata
	for _, porcelain := range porcelains {
		items = append(items, BatchUpdateResponseMetadataToPlumbing(porcelain))
	}
	return items
}

func RepeatedBatchUpdateResponseMetadataToPorcelain(plumbings []*BatchUpdateResponseMetadata) []models.BatchUpdateResponseMetadata {
	var items []models.BatchUpdateResponseMetadata
	for _, plumbing := range plumbings {
		items = append(items, BatchUpdateResponseMetadataToPorcelain(plumbing))
	}
	return items
}

func BatchDeleteResponseMetadataToPorcelain(plumbing *BatchDeleteResponseMetadata) models.BatchDeleteResponseMetadata {
	porcelain := models.BatchDeleteResponseMetadata{}
	porcelain.Found = plumbing.Found
	porcelain.Affected = plumbing.Affected
	return porcelain
}

func BatchDeleteResponseMetadataToPlumbing(porcelain models.BatchDeleteResponseMetadata) *BatchDeleteResponseMetadata {
	plumbing := &BatchDeleteResponseMetadata{}
	plumbing.Found = porcelain.Found
	plumbing.Affected = porcelain.Affected
	return plumbing
}

func RepeatedBatchDeleteResponseMetadataToPlumbing(porcelains []models.BatchDeleteResponseMetadata) []*BatchDeleteResponseMetadata {
	var items []*BatchDeleteResponseMetadata
	for _, porcelain := range porcelains {
		items = append(items, BatchDeleteResponseMetadataToPlumbing(porcelain))
	}
	return items
}

func RepeatedBatchDeleteResponseMetadataToPorcelain(plumbings []*BatchDeleteResponseMetadata) []models.BatchDeleteResponseMetadata {
	var items []models.BatchDeleteResponseMetadata
	for _, plumbing := range plumbings {
		items = append(items, BatchDeleteResponseMetadataToPorcelain(plumbing))
	}
	return items
}

func NodeCreateResponseToPorcelain(plumbing *NodeCreateResponse) models.NodeCreateResponse {
	porcelain := models.NodeCreateResponse{}
	porcelain.Meta = CreateResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.Nodes = RepeatedNodeToPorcelain(plumbing.Nodes)
	porcelain.Tokens = RepeatedTokenToPorcelain(plumbing.Tokens)
	return porcelain
}

func NodeCreateResponseToPlumbing(porcelain models.NodeCreateResponse) *NodeCreateResponse {
	plumbing := &NodeCreateResponse{}
	plumbing.Meta = CreateResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.Nodes = RepeatedNodeToPlumbing(porcelain.Nodes)
	plumbing.Tokens = RepeatedTokenToPlumbing(porcelain.Tokens)
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

func NodeListResponseToPorcelain(plumbing *NodeListResponse) models.NodeListResponse {
	porcelain := models.NodeListResponse{}
	return porcelain
}

func NodeListResponseToPlumbing(porcelain models.NodeListResponse) *NodeListResponse {
	plumbing := &NodeListResponse{}
	return plumbing
}

func RepeatedNodeListResponseToPlumbing(porcelains []models.NodeListResponse) []*NodeListResponse {
	var items []*NodeListResponse
	for _, porcelain := range porcelains {
		items = append(items, NodeListResponseToPlumbing(porcelain))
	}
	return items
}

func RepeatedNodeListResponseToPorcelain(plumbings []*NodeListResponse) []models.NodeListResponse {
	var items []models.NodeListResponse
	for _, plumbing := range plumbings {
		items = append(items, NodeListResponseToPorcelain(plumbing))
	}
	return items
}

func NodeBatchUpdateResponseToPorcelain(plumbing *NodeBatchUpdateResponse) models.NodeBatchUpdateResponse {
	porcelain := models.NodeBatchUpdateResponse{}
	porcelain.Meta = BatchUpdateResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.Nodes = RepeatedNodeToPorcelain(plumbing.Nodes)
	return porcelain
}

func NodeBatchUpdateResponseToPlumbing(porcelain models.NodeBatchUpdateResponse) *NodeBatchUpdateResponse {
	plumbing := &NodeBatchUpdateResponse{}
	plumbing.Meta = BatchUpdateResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.Nodes = RepeatedNodeToPlumbing(porcelain.Nodes)
	return plumbing
}

func RepeatedNodeBatchUpdateResponseToPlumbing(porcelains []models.NodeBatchUpdateResponse) []*NodeBatchUpdateResponse {
	var items []*NodeBatchUpdateResponse
	for _, porcelain := range porcelains {
		items = append(items, NodeBatchUpdateResponseToPlumbing(porcelain))
	}
	return items
}

func RepeatedNodeBatchUpdateResponseToPorcelain(plumbings []*NodeBatchUpdateResponse) []models.NodeBatchUpdateResponse {
	var items []models.NodeBatchUpdateResponse
	for _, plumbing := range plumbings {
		items = append(items, NodeBatchUpdateResponseToPorcelain(plumbing))
	}
	return items
}

func NodeBatchDeleteResponseToPorcelain(plumbing *NodeBatchDeleteResponse) models.NodeBatchDeleteResponse {
	porcelain := models.NodeBatchDeleteResponse{}
	porcelain.Meta = BatchDeleteResponseMetadataToPorcelain(plumbing.Meta)
	return porcelain
}

func NodeBatchDeleteResponseToPlumbing(porcelain models.NodeBatchDeleteResponse) *NodeBatchDeleteResponse {
	plumbing := &NodeBatchDeleteResponse{}
	plumbing.Meta = BatchDeleteResponseMetadataToPlumbing(porcelain.Meta)
	return plumbing
}

func RepeatedNodeBatchDeleteResponseToPlumbing(porcelains []models.NodeBatchDeleteResponse) []*NodeBatchDeleteResponse {
	var items []*NodeBatchDeleteResponse
	for _, porcelain := range porcelains {
		items = append(items, NodeBatchDeleteResponseToPlumbing(porcelain))
	}
	return items
}

func RepeatedNodeBatchDeleteResponseToPorcelain(plumbings []*NodeBatchDeleteResponse) []models.NodeBatchDeleteResponse {
	var items []models.NodeBatchDeleteResponse
	for _, plumbing := range plumbings {
		items = append(items, NodeBatchDeleteResponseToPorcelain(plumbing))
	}
	return items
}

func NodeToPlumbing(porcelain models.Node) *Node {
	plumbing := &Node{}

	switch v := porcelain.(type) {
	case *models.Relay:
		plumbing.Node = &Node_Relay{Relay: RelayToPlumbing(*v)}
	case *models.Gateway:
		plumbing.Node = &Node_Gateway{Gateway: GatewayToPlumbing(*v)}
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
				e := &errors.KindInternalError{}
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
	buffer      []models.Node
	index       int
	hasNextPage bool
	err         error
	fetch       NodeIteratorImplFetchFunc
}

func NewNodeIteratorImpl(f NodeIteratorImplFetchFunc) *NodeIteratorImpl {
	return &NodeIteratorImpl{
		hasNextPage: true,
		fetch:       f,
	}
}

func (n *NodeIteratorImpl) Next() bool {
	if n.index < len(n.buffer)-1 {
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
