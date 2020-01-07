package sdm

import (
	"context"

	plumbing "github.com/strongdm/strongdm-sdk-go/internal/v1"
)

// Nodes are proxies in the strongDM network. They come in two flavors: relays,
// which communicate with resources, and gateways, which communicate with
// clients.
type Nodes struct {
	client plumbing.NodesClient
	parent *Client
}

// Create registers a new Node.
func (svc *Nodes) Create(
	ctx context.Context,
	node Node,
) (
	*NodeCreateResponse,
	error,
) {
	req := &plumbing.NodeCreateRequest{}

	req.Node = nodeToPlumbing(node)
	var plumbingResponse *plumbing.NodeCreateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Create(svc.parent.wrapContext(ctx, req, "Nodes.Create"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &NodeCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = nodeToPorcelain(plumbingResponse.Node)
	resp.Token = plumbingResponse.Token
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one Node by ID.
func (svc *Nodes) Get(
	ctx context.Context,
	id string,
) (
	*NodeGetResponse,
	error,
) {
	req := &plumbing.NodeGetRequest{}

	req.Id = id
	var plumbingResponse *plumbing.NodeGetResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Get(svc.parent.wrapContext(ctx, req, "Nodes.Get"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &NodeGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = nodeToPorcelain(plumbingResponse.Node)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Update patches a Node by ID.
func (svc *Nodes) Update(
	ctx context.Context,
	node Node,
) (
	*NodeUpdateResponse,
	error,
) {
	req := &plumbing.NodeUpdateRequest{}

	req.Node = nodeToPlumbing(node)
	var plumbingResponse *plumbing.NodeUpdateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Update(svc.parent.wrapContext(ctx, req, "Nodes.Update"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &NodeUpdateResponse{}
	resp.Meta = updateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = nodeToPorcelain(plumbingResponse.Node)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes a Node by ID.
func (svc *Nodes) Delete(
	ctx context.Context,
	id string,
) (
	*NodeDeleteResponse,
	error,
) {
	req := &plumbing.NodeDeleteRequest{}

	req.Id = id
	var plumbingResponse *plumbing.NodeDeleteResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Delete(svc.parent.wrapContext(ctx, req, "Nodes.Delete"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &NodeDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of Nodes matching a given set of criteria.
func (svc *Nodes) List(
	ctx context.Context,
	filter string,
	args ...interface{},
) (
	NodeIterator,
	error,
) {
	req := &plumbing.NodeListRequest{}

	var filterErr error
	req.Filter, filterErr = quoteFilterArgs(filter, args...)
	if filterErr != nil {
		return nil, filterErr
	}
	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newNodeIteratorImpl(
		func() (
			[]Node,
			bool, error) {
			var plumbingResponse *plumbing.NodeListResponse
			var err error
			i := 0
			for {
				plumbingResponse, err = svc.client.List(svc.parent.wrapContext(ctx, req, "Nodes.List"), req)
				if err != nil {
					if !svc.parent.shouldRetry(i, err) {
						return nil, false, errorToPorcelain(err)
					}
					i++
					svc.parent.jitterSleep(i)
					continue
				}
				break
			}
			result := repeatedNodeToPorcelain(plumbingResponse.Nodes)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	), nil
}

type Resources struct {
	client plumbing.ResourcesClient
	parent *Client
}

// Create registers a new Resource.
func (svc *Resources) Create(
	ctx context.Context,
	resource Resource,
) (
	*ResourceCreateResponse,
	error,
) {
	req := &plumbing.ResourceCreateRequest{}

	req.Resource = resourceToPlumbing(resource)
	var plumbingResponse *plumbing.ResourceCreateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Create(svc.parent.wrapContext(ctx, req, "Resources.Create"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &ResourceCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Resource = resourceToPorcelain(plumbingResponse.Resource)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one Resource by ID.
func (svc *Resources) Get(
	ctx context.Context,
	id string,
) (
	*ResourceGetResponse,
	error,
) {
	req := &plumbing.ResourceGetRequest{}

	req.Id = id
	var plumbingResponse *plumbing.ResourceGetResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Get(svc.parent.wrapContext(ctx, req, "Resources.Get"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &ResourceGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Resource = resourceToPorcelain(plumbingResponse.Resource)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Update patches a Resource by ID.
func (svc *Resources) Update(
	ctx context.Context,
	resource Resource,
) (
	*ResourceUpdateResponse,
	error,
) {
	req := &plumbing.ResourceUpdateRequest{}

	req.Resource = resourceToPlumbing(resource)
	var plumbingResponse *plumbing.ResourceUpdateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Update(svc.parent.wrapContext(ctx, req, "Resources.Update"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &ResourceUpdateResponse{}
	resp.Meta = updateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Resource = resourceToPorcelain(plumbingResponse.Resource)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes a Resource by ID.
func (svc *Resources) Delete(
	ctx context.Context,
	id string,
) (
	*ResourceDeleteResponse,
	error,
) {
	req := &plumbing.ResourceDeleteRequest{}

	req.Id = id
	var plumbingResponse *plumbing.ResourceDeleteResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Delete(svc.parent.wrapContext(ctx, req, "Resources.Delete"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &ResourceDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of Resources matching a given set of criteria.
func (svc *Resources) List(
	ctx context.Context,
	filter string,
	args ...interface{},
) (
	ResourceIterator,
	error,
) {
	req := &plumbing.ResourceListRequest{}

	var filterErr error
	req.Filter, filterErr = quoteFilterArgs(filter, args...)
	if filterErr != nil {
		return nil, filterErr
	}
	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newResourceIteratorImpl(
		func() (
			[]Resource,
			bool, error) {
			var plumbingResponse *plumbing.ResourceListResponse
			var err error
			i := 0
			for {
				plumbingResponse, err = svc.client.List(svc.parent.wrapContext(ctx, req, "Resources.List"), req)
				if err != nil {
					if !svc.parent.shouldRetry(i, err) {
						return nil, false, errorToPorcelain(err)
					}
					i++
					svc.parent.jitterSleep(i)
					continue
				}
				break
			}
			result := repeatedResourceToPorcelain(plumbingResponse.Resources)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	), nil
}

// RoleAttachments represent relationships between composite roles and the roles
// that make up those composite roles. When a composite role is attached to another
// role, the permissions granted to members of the composite role are augmented to
// include the permissions granted to members of the attached role.
type RoleAttachments struct {
	client plumbing.RoleAttachmentsClient
	parent *Client
}

// Create registers a new RoleAttachment.
func (svc *RoleAttachments) Create(
	ctx context.Context,
	roleAttachment *RoleAttachment,
) (
	*RoleAttachmentCreateResponse,
	error,
) {
	req := &plumbing.RoleAttachmentCreateRequest{}

	req.RoleAttachment = roleAttachmentToPlumbing(roleAttachment)
	var plumbingResponse *plumbing.RoleAttachmentCreateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Create(svc.parent.wrapContext(ctx, req, "RoleAttachments.Create"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleAttachmentCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RoleAttachment = roleAttachmentToPorcelain(plumbingResponse.RoleAttachment)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one RoleAttachment by ID.
func (svc *RoleAttachments) Get(
	ctx context.Context,
	id string,
) (
	*RoleAttachmentGetResponse,
	error,
) {
	req := &plumbing.RoleAttachmentGetRequest{}

	req.Id = id
	var plumbingResponse *plumbing.RoleAttachmentGetResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Get(svc.parent.wrapContext(ctx, req, "RoleAttachments.Get"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleAttachmentGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RoleAttachment = roleAttachmentToPorcelain(plumbingResponse.RoleAttachment)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes a RoleAttachment by ID.
func (svc *RoleAttachments) Delete(
	ctx context.Context,
	id string,
) (
	*RoleAttachmentDeleteResponse,
	error,
) {
	req := &plumbing.RoleAttachmentDeleteRequest{}

	req.Id = id
	var plumbingResponse *plumbing.RoleAttachmentDeleteResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Delete(svc.parent.wrapContext(ctx, req, "RoleAttachments.Delete"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleAttachmentDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of RoleAttachments matching a given set of criteria.
func (svc *RoleAttachments) List(
	ctx context.Context,
	filter string,
	args ...interface{},
) (
	RoleAttachmentIterator,
	error,
) {
	req := &plumbing.RoleAttachmentListRequest{}

	var filterErr error
	req.Filter, filterErr = quoteFilterArgs(filter, args...)
	if filterErr != nil {
		return nil, filterErr
	}
	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newRoleAttachmentIteratorImpl(
		func() (
			[]*RoleAttachment,
			bool, error) {
			var plumbingResponse *plumbing.RoleAttachmentListResponse
			var err error
			i := 0
			for {
				plumbingResponse, err = svc.client.List(svc.parent.wrapContext(ctx, req, "RoleAttachments.List"), req)
				if err != nil {
					if !svc.parent.shouldRetry(i, err) {
						return nil, false, errorToPorcelain(err)
					}
					i++
					svc.parent.jitterSleep(i)
					continue
				}
				break
			}
			result := repeatedRoleAttachmentToPorcelain(plumbingResponse.RoleAttachments)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	), nil
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
func (svc *Roles) Create(
	ctx context.Context,
	role *Role,
) (
	*RoleCreateResponse,
	error,
) {
	req := &plumbing.RoleCreateRequest{}

	req.Role = roleToPlumbing(role)
	var plumbingResponse *plumbing.RoleCreateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Create(svc.parent.wrapContext(ctx, req, "Roles.Create"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Role = roleToPorcelain(plumbingResponse.Role)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one Role by ID.
func (svc *Roles) Get(
	ctx context.Context,
	id string,
) (
	*RoleGetResponse,
	error,
) {
	req := &plumbing.RoleGetRequest{}

	req.Id = id
	var plumbingResponse *plumbing.RoleGetResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Get(svc.parent.wrapContext(ctx, req, "Roles.Get"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Role = roleToPorcelain(plumbingResponse.Role)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Update patches a Role by ID.
func (svc *Roles) Update(
	ctx context.Context,
	role *Role,
) (
	*RoleUpdateResponse,
	error,
) {
	req := &plumbing.RoleUpdateRequest{}

	req.Role = roleToPlumbing(role)
	var plumbingResponse *plumbing.RoleUpdateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Update(svc.parent.wrapContext(ctx, req, "Roles.Update"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleUpdateResponse{}
	resp.Meta = updateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Role = roleToPorcelain(plumbingResponse.Role)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes a Role by ID.
func (svc *Roles) Delete(
	ctx context.Context,
	id string,
) (
	*RoleDeleteResponse,
	error,
) {
	req := &plumbing.RoleDeleteRequest{}

	req.Id = id
	var plumbingResponse *plumbing.RoleDeleteResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Delete(svc.parent.wrapContext(ctx, req, "Roles.Delete"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of Roles matching a given set of criteria.
func (svc *Roles) List(
	ctx context.Context,
	filter string,
	args ...interface{},
) (
	RoleIterator,
	error,
) {
	req := &plumbing.RoleListRequest{}

	var filterErr error
	req.Filter, filterErr = quoteFilterArgs(filter, args...)
	if filterErr != nil {
		return nil, filterErr
	}
	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newRoleIteratorImpl(
		func() (
			[]*Role,
			bool, error) {
			var plumbingResponse *plumbing.RoleListResponse
			var err error
			i := 0
			for {
				plumbingResponse, err = svc.client.List(svc.parent.wrapContext(ctx, req, "Roles.List"), req)
				if err != nil {
					if !svc.parent.shouldRetry(i, err) {
						return nil, false, errorToPorcelain(err)
					}
					i++
					svc.parent.jitterSleep(i)
					continue
				}
				break
			}
			result := repeatedRoleToPorcelain(plumbingResponse.Roles)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	), nil
}
