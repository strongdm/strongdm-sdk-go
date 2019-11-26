package sdm

import (
	"context"

	plumbing "github.com/strongdm/strongdm-sdk-go/internal/v1"
	models "github.com/strongdm/strongdm-sdk-go/models"
)




// Nodes are proxies in the strongDM network. They come in two flavors: relays,
// which communicate with resources, and gateways, which communicate with
// clients.
type Nodes struct {
	client   plumbing.NodesClient
	parent   *Client
}

// Create registers a new Node.
func (svc *Nodes) Create(ctx context.Context, node models.Node) (*models.NodeCreateResponse, error) {
	req := &plumbing.NodeCreateRequest{}
	req.Node = plumbing.NodeToPlumbing(node)
	
	plumbingResponse, err := svc.client.Create(svc.parent.wrapContext(ctx), req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.NodeCreateResponse{}
	resp.Meta = plumbing.CreateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = plumbing.NodeToPorcelain(plumbingResponse.Node)
	resp.Token = plumbingResponse.Token
	return resp, nil
}

// Get reads one Node by ID.
func (svc *Nodes) Get(ctx context.Context, id string) (*models.NodeGetResponse, error) {
	req := &plumbing.NodeGetRequest{}
	req.Id = id
	
	plumbingResponse, err := svc.client.Get(svc.parent.wrapContext(ctx), req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.NodeGetResponse{}
	resp.Meta = plumbing.GetResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = plumbing.NodeToPorcelain(plumbingResponse.Node)
	return resp, nil
}

// Update patches a Node by ID.
func (svc *Nodes) Update(ctx context.Context, node models.Node) (*models.NodeUpdateResponse, error) {
	req := &plumbing.NodeUpdateRequest{}
	req.Node = plumbing.NodeToPlumbing(node)
	
	plumbingResponse, err := svc.client.Update(svc.parent.wrapContext(ctx), req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.NodeUpdateResponse{}
	resp.Meta = plumbing.UpdateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = plumbing.NodeToPorcelain(plumbingResponse.Node)
	return resp, nil
}

// Delete removes a Node by ID.
func (svc *Nodes) Delete(ctx context.Context, id string) (*models.NodeDeleteResponse, error) {
	req := &plumbing.NodeDeleteRequest{}
	req.Id = id
	
	plumbingResponse, err := svc.client.Delete(svc.parent.wrapContext(ctx), req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.NodeDeleteResponse{}
	resp.Meta = plumbing.DeleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	return resp, nil
}

// List gets a list of Nodes matching a given set of criteria.
func (svc *Nodes) List(ctx context.Context, filter string) models.NodeIterator {
	req := &plumbing.NodeListRequest{}
	req.Filter = filter
	
	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return plumbing.NewNodeIteratorImpl(
		func() ([]models.Node, bool, error) {
			plumbingResponse, err := svc.client.List(svc.parent.wrapContext(ctx), req)
			if err != nil {
				return nil, false, plumbing.ErrorToPorcelain(err)
			}
			result := plumbing.RepeatedNodeToPorcelain(plumbingResponse.Nodes)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	)
}




// Roles are tools for controlling user access to resources. Each Role holds a
// list of resources which they grant access to. Composite roles are a special
// type of Role which have no resource associations of their own, but instead
// grant access to the combined resources associated with a set of child roles.
// Each user can be a member of one Role or composite role.
type Roles struct {
	client   plumbing.RolesClient
	parent   *Client
}

// Create registers a new Role.
func (svc *Roles) Create(ctx context.Context, role *models.Role) (*models.RoleCreateResponse, error) {
	req := &plumbing.RoleCreateRequest{}
	req.Role = plumbing.RoleToPlumbing(role)
	
	plumbingResponse, err := svc.client.Create(svc.parent.wrapContext(ctx), req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.RoleCreateResponse{}
	resp.Meta = plumbing.CreateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Role = plumbing.RoleToPorcelain(plumbingResponse.Role)
	return resp, nil
}

// Get reads one Role by ID.
func (svc *Roles) Get(ctx context.Context, id string) (*models.RoleGetResponse, error) {
	req := &plumbing.RoleGetRequest{}
	req.Id = id
	
	plumbingResponse, err := svc.client.Get(svc.parent.wrapContext(ctx), req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.RoleGetResponse{}
	resp.Meta = plumbing.GetResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Role = plumbing.RoleToPorcelain(plumbingResponse.Role)
	return resp, nil
}

// Update patches a Role by ID.
func (svc *Roles) Update(ctx context.Context, role *models.Role) (*models.RoleUpdateResponse, error) {
	req := &plumbing.RoleUpdateRequest{}
	req.Role = plumbing.RoleToPlumbing(role)
	
	plumbingResponse, err := svc.client.Update(svc.parent.wrapContext(ctx), req)
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
	req.Id = id
	
	plumbingResponse, err := svc.client.Delete(svc.parent.wrapContext(ctx), req)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(err)
	}
	resp := &models.RoleDeleteResponse{}
	resp.Meta = plumbing.DeleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	return resp, nil
}

// List gets a list of Roles matching a given set of criteria.
func (svc *Roles) List(ctx context.Context, filter string) models.RoleIterator {
	req := &plumbing.RoleListRequest{}
	req.Filter = filter
	
	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return plumbing.NewRoleIteratorImpl(
		func() ([]*models.Role, bool, error) {
			plumbingResponse, err := svc.client.List(svc.parent.wrapContext(ctx), req)
			if err != nil {
				return nil, false, plumbing.ErrorToPorcelain(err)
			}
			result := plumbing.RepeatedRoleToPorcelain(plumbingResponse.Roles)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	)
}

