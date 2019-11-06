package sdm

import (
	"context"

	v1 "github.com/strongdm/strongdm-sdk-go/internal/v1"
	models "github.com/strongdm/strongdm-sdk-go/models"
	errors "github.com/strongdm/strongdm-sdk-go/errors"
)



// Nodes are proxies in strongDM responsible to communicate with servers
// (relays) and clients (gateways).
type Nodes struct {
	client v1.NodesClient
}

// Create registers a new node.
func (svc *Nodes) Create(ctx context.Context, nodes ...models.Node) (*models.NodeCreateResponse, error) {
	req := &v1.NodeCreateRequest{}
	req.Nodes = v1.RepeatedNodeToPlumbing(nodes)
	
	plumbingResponse, err := svc.client.Create(ctx, req)
	if err != nil {
		return nil, errors.New(err)
	}
	resp := &models.NodeCreateResponse{}
	resp.Meta = v1.CreateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Nodes = v1.RepeatedNodeToPorcelain(plumbingResponse.Nodes)
	resp.Tokens = v1.RepeatedTokenToPorcelain(plumbingResponse.Tokens)
	return resp, nil
}

// Get reads one node by ID.
func (svc *Nodes) Get(ctx context.Context, id string) (*models.NodeGetResponse, error) {
	req := &v1.NodeGetRequest{}
	req.Id = id
	
	plumbingResponse, err := svc.client.Get(ctx, req)
	if err != nil {
		return nil, errors.New(err)
	}
	resp := &models.NodeGetResponse{}
	resp.Meta = v1.GetResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = v1.NodeToPorcelain(plumbingResponse.Node)
	return resp, nil
}

// Update patches a node by ID.
func (svc *Nodes) Update(ctx context.Context, id string, node models.Node) (*models.NodeUpdateResponse, error) {
	req := &v1.NodeUpdateRequest{}
	req.Id = id
	req.Node = v1.NodeToPlumbing(node)
	
	plumbingResponse, err := svc.client.Update(ctx, req)
	if err != nil {
		return nil, errors.New(err)
	}
	resp := &models.NodeUpdateResponse{}
	resp.Meta = v1.UpdateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = v1.NodeToPorcelain(plumbingResponse.Node)
	return resp, nil
}

// Delete removes a node by ID.
func (svc *Nodes) Delete(ctx context.Context, id string) (*models.NodeDeleteResponse, error) {
	req := &v1.NodeDeleteRequest{}
	req.Id = id
	
	plumbingResponse, err := svc.client.Delete(ctx, req)
	if err != nil {
		return nil, errors.New(err)
	}
	resp := &models.NodeDeleteResponse{}
	resp.Meta = v1.DeleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	return resp, nil
}

// List is a batched Get call.
func (svc *Nodes) List(ctx context.Context, filter string) (*models.NodeListResponse) {
	req := &v1.NodeListRequest{}
	req.Filter = filter
	
	req.Meta = &v1.ListRequestMetadata{}
	resp := &models.NodeListResponse{}
	iter := v1.NewNodeIteratorImpl(
		func() ([]models.Node, bool, error) {
			plumbingResponse, err := svc.client.List(ctx, req)
			if err != nil {
				return nil, false, errors.New(err)
			}
			var result []models.Node
			
			result = v1.RepeatedNodeToPorcelain(plumbingResponse.Nodes)
			
			req.Meta.Page = plumbingResponse.Meta.NextPage
			return result, req.Meta.Page != "", nil
		},
	)
	resp.Nodes = iter
	return resp
}

// BatchUpdate is a batched Update call.
func (svc *Nodes) BatchUpdate(ctx context.Context, nodes ...models.Node) (*models.NodeBatchUpdateResponse, error) {
	req := &v1.NodeBatchUpdateRequest{}
	req.Nodes = v1.RepeatedNodeToPlumbing(nodes)
	
	plumbingResponse, err := svc.client.BatchUpdate(ctx, req)
	if err != nil {
		return nil, errors.New(err)
	}
	resp := &models.NodeBatchUpdateResponse{}
	resp.Meta = v1.BatchUpdateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Nodes = v1.RepeatedNodeToPorcelain(plumbingResponse.Nodes)
	return resp, nil
}

// BatchDelete is a batched Delete call.
func (svc *Nodes) BatchDelete(ctx context.Context, ids ...string) (*models.NodeBatchDeleteResponse, error) {
	req := &v1.NodeBatchDeleteRequest{}
	req.Ids = append(req.Ids, ids...)
	
	plumbingResponse, err := svc.client.BatchDelete(ctx, req)
	if err != nil {
		return nil, errors.New(err)
	}
	resp := &models.NodeBatchDeleteResponse{}
	resp.Meta = v1.BatchDeleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	return resp, nil
}

