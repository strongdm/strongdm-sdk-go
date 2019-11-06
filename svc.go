package sdm

import (
	"context"

	plumbing "github.com/strongdm/strongdm-sdk-go/internal/v1"
	models "github.com/strongdm/strongdm-sdk-go/models"
)



// Nodes are proxies in strongDM responsible to communicate with servers
// (relays) and clients (gateways).
type Nodes struct {
	client plumbing.NodesClient
}

// Create registers a new node.
func (svc *Nodes) Create(ctx context.Context, nodes ...models.Node) (*models.NodeCreateResponse, error) {
	req := &plumbing.NodeCreateRequest{}
	req.Nodes = plumbing.RepeatedNodeToPlumbing(nodes)
	
	plumbingResponse, err := svc.client.Create(ctx, req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.NodeCreateResponse{}
	resp.Meta = plumbing.CreateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Nodes = plumbing.RepeatedNodeToPorcelain(plumbingResponse.Nodes)
	resp.Tokens = plumbing.RepeatedTokenToPorcelain(plumbingResponse.Tokens)
	return resp, nil
}

// Get reads one node by ID.
func (svc *Nodes) Get(ctx context.Context, id string) (*models.NodeGetResponse, error) {
	req := &plumbing.NodeGetRequest{}
	req.Id = id
	
	plumbingResponse, err := svc.client.Get(ctx, req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.NodeGetResponse{}
	resp.Meta = plumbing.GetResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = plumbing.NodeToPorcelain(plumbingResponse.Node)
	return resp, nil
}

// Update patches a node by ID.
func (svc *Nodes) Update(ctx context.Context, id string, node models.Node) (*models.NodeUpdateResponse, error) {
	req := &plumbing.NodeUpdateRequest{}
	req.Id = id
	req.Node = plumbing.NodeToPlumbing(node)
	
	plumbingResponse, err := svc.client.Update(ctx, req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.NodeUpdateResponse{}
	resp.Meta = plumbing.UpdateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = plumbing.NodeToPorcelain(plumbingResponse.Node)
	return resp, nil
}

// Delete removes a node by ID.
func (svc *Nodes) Delete(ctx context.Context, id string) (*models.NodeDeleteResponse, error) {
	req := &plumbing.NodeDeleteRequest{}
	req.Id = id
	
	plumbingResponse, err := svc.client.Delete(ctx, req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.NodeDeleteResponse{}
	resp.Meta = plumbing.DeleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	return resp, nil
}

// List is a batched Get call.
func (svc *Nodes) List(ctx context.Context, filter string) (*models.NodeListResponse) {
	req := &plumbing.NodeListRequest{}
	req.Filter = filter
	
	req.Meta = &plumbing.ListRequestMetadata{}
	resp := &models.NodeListResponse{}
	iter := plumbing.NewNodeIteratorImpl(
		func() ([]models.Node, bool, error) {
			plumbingResponse, err := svc.client.List(ctx, req)
			if err != nil {
				return nil, false, plumbing.ErrorToPorcelain(err)
			}
			var result []models.Node
			
			result = plumbing.RepeatedNodeToPorcelain(plumbingResponse.Nodes)
			
			req.Meta.Page = plumbingResponse.Meta.NextPage
			return result, req.Meta.Page != "", nil
		},
	)
	resp.Nodes = iter
	return resp
}

// BatchUpdate is a batched Update call.
func (svc *Nodes) BatchUpdate(ctx context.Context, nodes ...models.Node) (*models.NodeBatchUpdateResponse, error) {
	req := &plumbing.NodeBatchUpdateRequest{}
	req.Nodes = plumbing.RepeatedNodeToPlumbing(nodes)
	
	plumbingResponse, err := svc.client.BatchUpdate(ctx, req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.NodeBatchUpdateResponse{}
	resp.Meta = plumbing.BatchUpdateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Nodes = plumbing.RepeatedNodeToPorcelain(plumbingResponse.Nodes)
	return resp, nil
}

// BatchDelete is a batched Delete call.
func (svc *Nodes) BatchDelete(ctx context.Context, ids ...string) (*models.NodeBatchDeleteResponse, error) {
	req := &plumbing.NodeBatchDeleteRequest{}
	req.Ids = append(req.Ids, ids...)
	
	plumbingResponse, err := svc.client.BatchDelete(ctx, req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.NodeBatchDeleteResponse{}
	resp.Meta = plumbing.BatchDeleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	return resp, nil
}

