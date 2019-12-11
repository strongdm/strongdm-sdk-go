package sdm

import (
	"context"

	plumbing "github.com/strongdm/strongdm-sdk-go/internal/v1"
)

type Attachments struct {
	client plumbing.AttachmentsClient
	parent *Client
}

// Create registers a new Attachment.
func (svc *Attachments) Create(ctx context.Context, attachment *Attachment) (*AttachmentCreateResponse, error) {
	req := &plumbing.AttachmentCreateRequest{}
	req.Attachment = attachmentToPlumbing(attachment)

	plumbingResponse, err := svc.client.Create(svc.parent.wrapContext(ctx, req), req)
	if err != nil {
		return nil, errorToPorcelain(err)
	}
	resp := &AttachmentCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Attachment = attachmentToPorcelain(plumbingResponse.Attachment)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one Attachment by ID.
func (svc *Attachments) Get(ctx context.Context, id string) (*AttachmentGetResponse, error) {
	req := &plumbing.AttachmentGetRequest{}
	req.Id = id

	plumbingResponse, err := svc.client.Get(svc.parent.wrapContext(ctx, req), req)
	if err != nil {
		return nil, errorToPorcelain(err)
	}
	resp := &AttachmentGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Attachment = attachmentToPorcelain(plumbingResponse.Attachment)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes an Attachment by ID.
func (svc *Attachments) Delete(ctx context.Context, id string) (*AttachmentDeleteResponse, error) {
	req := &plumbing.AttachmentDeleteRequest{}
	req.Id = id

	plumbingResponse, err := svc.client.Delete(svc.parent.wrapContext(ctx, req), req)
	if err != nil {
		return nil, errorToPorcelain(err)
	}
	resp := &AttachmentDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of Attachments matching a given set of criteria.
func (svc *Attachments) List(ctx context.Context, filter string) AttachmentIterator {
	req := &plumbing.AttachmentListRequest{}
	req.Filter = filter

	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newAttachmentIteratorImpl(
		func() ([]*Attachment, bool, error) {
			plumbingResponse, err := svc.client.List(svc.parent.wrapContext(ctx, req), req)
			if err != nil {
				return nil, false, errorToPorcelain(err)
			}
			result := repeatedAttachmentToPorcelain(plumbingResponse.Attachments)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	)
}

// Nodes are proxies in the strongDM network. They come in two flavors: relays,
// which communicate with resources, and gateways, which communicate with
// clients.
type Nodes struct {
	client plumbing.NodesClient
	parent *Client
}

// Create registers a new Node.
func (svc *Nodes) Create(ctx context.Context, node Node) (*NodeCreateResponse, error) {
	req := &plumbing.NodeCreateRequest{}
	req.Node = nodeToPlumbing(node)

	plumbingResponse, err := svc.client.Create(svc.parent.wrapContext(ctx, req), req)
	if err != nil {
		return nil, errorToPorcelain(err)
	}
	resp := &NodeCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = nodeToPorcelain(plumbingResponse.Node)
	resp.Token = plumbingResponse.Token
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one Node by ID.
func (svc *Nodes) Get(ctx context.Context, id string) (*NodeGetResponse, error) {
	req := &plumbing.NodeGetRequest{}
	req.Id = id

	plumbingResponse, err := svc.client.Get(svc.parent.wrapContext(ctx, req), req)
	if err != nil {
		return nil, errorToPorcelain(err)
	}
	resp := &NodeGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = nodeToPorcelain(plumbingResponse.Node)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Update patches a Node by ID.
func (svc *Nodes) Update(ctx context.Context, node Node) (*NodeUpdateResponse, error) {
	req := &plumbing.NodeUpdateRequest{}
	req.Node = nodeToPlumbing(node)

	plumbingResponse, err := svc.client.Update(svc.parent.wrapContext(ctx, req), req)
	if err != nil {
		return nil, errorToPorcelain(err)
	}
	resp := &NodeUpdateResponse{}
	resp.Meta = updateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = nodeToPorcelain(plumbingResponse.Node)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes a Node by ID.
func (svc *Nodes) Delete(ctx context.Context, id string) (*NodeDeleteResponse, error) {
	req := &plumbing.NodeDeleteRequest{}
	req.Id = id

	plumbingResponse, err := svc.client.Delete(svc.parent.wrapContext(ctx, req), req)
	if err != nil {
		return nil, errorToPorcelain(err)
	}
	resp := &NodeDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of Nodes matching a given set of criteria.
func (svc *Nodes) List(ctx context.Context, filter string) NodeIterator {
	req := &plumbing.NodeListRequest{}
	req.Filter = filter

	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newNodeIteratorImpl(
		func() ([]Node, bool, error) {
			plumbingResponse, err := svc.client.List(svc.parent.wrapContext(ctx, req), req)
			if err != nil {
				return nil, false, errorToPorcelain(err)
			}
			result := repeatedNodeToPorcelain(plumbingResponse.Nodes)
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
	client plumbing.RolesClient
	parent *Client
}

// Create registers a new Role.
func (svc *Roles) Create(ctx context.Context, role *Role) (*RoleCreateResponse, error) {
	req := &plumbing.RoleCreateRequest{}
	req.Role = roleToPlumbing(role)

	plumbingResponse, err := svc.client.Create(svc.parent.wrapContext(ctx, req), req)
	if err != nil {
		return nil, errorToPorcelain(err)
	}
	resp := &RoleCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Role = roleToPorcelain(plumbingResponse.Role)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one Role by ID.
func (svc *Roles) Get(ctx context.Context, id string) (*RoleGetResponse, error) {
	req := &plumbing.RoleGetRequest{}
	req.Id = id

	plumbingResponse, err := svc.client.Get(svc.parent.wrapContext(ctx, req), req)
	if err != nil {
		return nil, errorToPorcelain(err)
	}
	resp := &RoleGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Role = roleToPorcelain(plumbingResponse.Role)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Update patches a Role by ID.
func (svc *Roles) Update(ctx context.Context, role *Role) (*RoleUpdateResponse, error) {
	req := &plumbing.RoleUpdateRequest{}
	req.Role = roleToPlumbing(role)

	plumbingResponse, err := svc.client.Update(svc.parent.wrapContext(ctx, req), req)
	if err != nil {
		return nil, errorToPorcelain(err)
	}
	resp := &RoleUpdateResponse{}
	resp.Meta = updateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Role = roleToPorcelain(plumbingResponse.Role)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes a Role by ID.
func (svc *Roles) Delete(ctx context.Context, id string) (*RoleDeleteResponse, error) {
	req := &plumbing.RoleDeleteRequest{}
	req.Id = id

	plumbingResponse, err := svc.client.Delete(svc.parent.wrapContext(ctx, req), req)
	if err != nil {
		return nil, errorToPorcelain(err)
	}
	resp := &RoleDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of Roles matching a given set of criteria.
func (svc *Roles) List(ctx context.Context, filter string) RoleIterator {
	req := &plumbing.RoleListRequest{}
	req.Filter = filter

	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newRoleIteratorImpl(
		func() ([]*Role, bool, error) {
			plumbingResponse, err := svc.client.List(svc.parent.wrapContext(ctx, req), req)
			if err != nil {
				return nil, false, errorToPorcelain(err)
			}
			result := repeatedRoleToPorcelain(plumbingResponse.Roles)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	)
}
