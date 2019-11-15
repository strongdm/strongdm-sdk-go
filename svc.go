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
	apiToken string
}

// Create registers a new node.
func (svc *Nodes) Create(ctx context.Context, node models.Node) (*models.NodeCreateResponse, error) {
	req := &plumbing.NodeCreateRequest{}
	
	ctx = plumbing.CreateGRPCContext(ctx, svc.apiToken)
	req.Node = plumbing.NodeToPlumbing(node)
	
	plumbingResponse, err := svc.client.Create(ctx, req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.NodeCreateResponse{}
	resp.Meta = plumbing.CreateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = plumbing.NodeToPorcelain(plumbingResponse.Node)
	resp.Token = plumbing.TokenToPorcelain(plumbingResponse.Token)
	return resp, nil
}

// Get reads one node by ID.
func (svc *Nodes) Get(ctx context.Context, id string) (*models.NodeGetResponse, error) {
	req := &plumbing.NodeGetRequest{}
	
	ctx = plumbing.CreateGRPCContext(ctx, svc.apiToken)
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
func (svc *Nodes) Update(ctx context.Context, node models.Node) (*models.NodeUpdateResponse, error) {
	req := &plumbing.NodeUpdateRequest{}
	
	ctx = plumbing.CreateGRPCContext(ctx, svc.apiToken)
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
	
	ctx = plumbing.CreateGRPCContext(ctx, svc.apiToken)
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
func (svc *Nodes) List(ctx context.Context, filter string) models.NodeIterator {
	req := &plumbing.NodeListRequest{}
	
	ctx = plumbing.CreateGRPCContext(ctx, svc.apiToken)
	req.Filter = filter
	
	req.Meta = &plumbing.ListRequestMetadata{
		Page: 0,
		Limit: 25,
	}
	return plumbing.NewNodeIteratorImpl(
		func() ([]models.Node, bool, error) {
			plumbingResponse, err := svc.client.List(ctx, req)
			if err != nil {
				return nil, false, plumbing.ErrorToPorcelain(err)
			}
			var result []models.Node
			
			result = plumbing.RepeatedNodeToPorcelain(plumbingResponse.Nodes)
			
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	)
}




// Roles are
type Roles struct {
	client plumbing.RolesClient
	apiToken string
}

// Create registers a new role.
func (svc *Roles) Create(ctx context.Context, role models.Role) (*models.RoleCreateResponse, error) {
	req := &plumbing.RoleCreateRequest{}
	
	ctx = plumbing.CreateGRPCContext(ctx, svc.apiToken)
	req.Role = plumbing.RoleToPlumbing(role)
	
	plumbingResponse, err := svc.client.Create(ctx, req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.RoleCreateResponse{}
	resp.Meta = plumbing.CreateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Role = plumbing.RoleToPorcelain(plumbingResponse.Role)
	return resp, nil
}

// Get reads one role by ID.
func (svc *Roles) Get(ctx context.Context, id string) (*models.RoleGetResponse, error) {
	req := &plumbing.RoleGetRequest{}
	
	ctx = plumbing.CreateGRPCContext(ctx, svc.apiToken)
	req.Id = id
	
	plumbingResponse, err := svc.client.Get(ctx, req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.RoleGetResponse{}
	resp.Meta = plumbing.GetResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Role = plumbing.RoleToPorcelain(plumbingResponse.Role)
	return resp, nil
}

// Update patches a Role by ID.
func (svc *Roles) Update(ctx context.Context, role models.Role) (*models.RoleUpdateResponse, error) {
	req := &plumbing.RoleUpdateRequest{}
	
	ctx = plumbing.CreateGRPCContext(ctx, svc.apiToken)
	req.Role = plumbing.RoleToPlumbing(role)
	
	plumbingResponse, err := svc.client.Update(ctx, req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.RoleUpdateResponse{}
	resp.Meta = plumbing.UpdateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Role = plumbing.RoleToPorcelain(plumbingResponse.Role)
	return resp, nil
}

// Delete removes a Role by ID.
func (svc *Roles) Delete(ctx context.Context, id string) (*models.RoleDeleteResponse, error) {
	req := &plumbing.RoleDeleteRequest{}
	
	ctx = plumbing.CreateGRPCContext(ctx, svc.apiToken)
	req.Id = id
	
	plumbingResponse, err := svc.client.Delete(ctx, req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.RoleDeleteResponse{}
	resp.Meta = plumbing.DeleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	return resp, nil
}

// List is a batched Get call.
func (svc *Roles) List(ctx context.Context, filter string) models.RoleIterator {
	req := &plumbing.RoleListRequest{}
	
	ctx = plumbing.CreateGRPCContext(ctx, svc.apiToken)
	req.Filter = filter
	
	req.Meta = &plumbing.ListRequestMetadata{
		Page: 0,
		Limit: 25,
	}
	return plumbing.NewRoleIteratorImpl(
		func() ([]models.Role, bool, error) {
			plumbingResponse, err := svc.client.List(ctx, req)
			if err != nil {
				return nil, false, plumbing.ErrorToPorcelain(err)
			}
			var result []models.Role
			
			result = plumbing.RepeatedRoleToPorcelain(plumbingResponse.Roles)
			
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	)
}

