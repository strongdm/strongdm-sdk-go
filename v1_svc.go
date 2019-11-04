package sdm

import (
	"context"

	v1 "github.com/strongdm/strongdm-sdk-go/internal/v1"
)


// Nodes are proxies in strongDM responsible to communicate with servers
// (relays) and clients (gateways).
type Nodes struct {
	client v1.NodesClient
}

// Create registers a new node.
func (n *Nodes) Create(ctx context.Context, req *v1.NodeCreateRequest) (*v1.NodeCreateResponse, error) {
	resp, err := n.Create(context.Background(), req)
	// TODO: improve error handling
	return resp, err
}

// Get reads one node by ID.
func (n *Nodes) Get(ctx context.Context, req *v1.NodeGetRequest) (*v1.NodeGetResponse, error) {
	resp, err := n.Get(context.Background(), req)
	// TODO: improve error handling
	return resp, err
}

// Update patches a node by ID.
func (n *Nodes) Update(ctx context.Context, req *v1.NodeUpdateRequest) (*v1.NodeUpdateResponse, error) {
	resp, err := n.Update(context.Background(), req)
	// TODO: improve error handling
	return resp, err
}

// Delete removes a node by ID.
func (n *Nodes) Delete(ctx context.Context, req *v1.NodeDeleteRequest) (*v1.NodeDeleteResponse, error) {
	resp, err := n.Delete(context.Background(), req)
	// TODO: improve error handling
	return resp, err
}

// List is a batched Get call.
func (n *Nodes) List(ctx context.Context, req *v1.NodeListRequest) (*v1.NodeListResponse, error) {
	resp, err := n.List(context.Background(), req)
	// TODO: improve error handling
	return resp, err
}

// BatchUpdate is a batched Update call.
func (n *Nodes) BatchUpdate(ctx context.Context, req *v1.NodeBatchUpdateRequest) (*v1.NodeBatchUpdateResponse, error) {
	resp, err := n.BatchUpdate(context.Background(), req)
	// TODO: improve error handling
	return resp, err
}

// BatchDelete is a batched Delete call.
func (n *Nodes) BatchDelete(ctx context.Context, req *v1.NodeBatchDeleteRequest) (*v1.NodeBatchDeleteResponse, error) {
	resp, err := n.BatchDelete(context.Background(), req)
	// TODO: improve error handling
	return resp, err
}
