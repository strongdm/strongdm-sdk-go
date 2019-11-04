package sdm

import (
	"context"

	v1 "github.com/strongdm/strongdm-sdk-go/internal/v1"
	models "github.com/strongdm/strongdm-sdk-go/models"
)



// Nodes are proxies in strongDM responsible to communicate with servers
// (relays) and clients (gateways).
type Nodes struct {
	client v1.NodesClient
}

// Create registers a new node.
func (*Nodes) Create(ctx context.Context, nodes ...models.Node) (*models.NodeCreateResponse) {
	return nil
}
// Get reads one node by ID.
func (*Nodes) Get(ctx context.Context, id string) (*models.NodeGetResponse) {
	return nil
}
// Update patches a node by ID.
func (*Nodes) Update(ctx context.Context, id string, node models.Node) (*models.NodeUpdateResponse) {
	return nil
}
// Delete removes a node by ID.
func (*Nodes) Delete(ctx context.Context, id string) (*models.NodeDeleteResponse) {
	return nil
}
// List is a batched Get call.
func (*Nodes) List(ctx context.Context, filter string) (*models.NodeListResponse) {
	return nil
}
// BatchUpdate is a batched Update call.
func (*Nodes) BatchUpdate(ctx context.Context, nodes ...models.Node) (*models.NodeBatchUpdateResponse) {
	return nil
}
// BatchDelete is a batched Delete call.
func (*Nodes) BatchDelete(ctx context.Context, ids ...string) (*models.NodeBatchDeleteResponse) {
	return nil
}
