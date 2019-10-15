/*
Package sdm implements an API client to strongDM restful API.
*/
package sdm

import (
	"context"
	"net/http"
	"sync"
)

type ClientOption struct {
	HTTPClient *http.Client
	Token string
}


// Relays service
type Relays struct {
	initOnce sync.Once
	httpClient *http.Client
	reqModifier func(*http.Request)
}

func NewRelays(opts ...ClientOption) *Relays {
	c:= &Relays{}
	c.init()
	for _, opt := range opts {
		if opt.HTTPClient != nil {
			c.httpClient = opt.HTTPClient
		}
		if opt.Token != "" {
			c.reqModifier = addTokenHeader(opt.Token)
		}
	}
	return c
}

func (r *Relays) init() {
	r.initOnce.Do(func(){
		r.httpClient = http.DefaultClient
		r.reqModifier = func(*http.Request){}
	})
}

// Create adds a new relay, and returns the new relay
func (r *Relays) Create(ctx context.Context, req RelayCreateRequest) (RelayCreateResponse, error) {
	// Host: app.strongdm.com
	// BasePath: 
	// Dial to /v1/relays by POST
	// Body: *
	// ResponseBody: 
	panic("not implemented")
}

// Update modifies an existing relay
func (r *Relays) Update(ctx context.Context, req RelayUpdateRequest) (RelayUpdateResponse, error) {
	// Host: app.strongdm.com
	// BasePath: 
	// Dial to /v1/relays/{node.id} by POST
	// Body: *
	// ResponseBody: 
	panic("not implemented")
}

// Delete removes an existing relay
func (r *Relays) Delete(ctx context.Context, req RelayDeleteRequest) (RelayDeleteResponse, error) {
	// Host: app.strongdm.com
	// BasePath: 
	// Dial to /v1/relays/{id} by DELETE
	// Body: 
	// ResponseBody: 
	panic("not implemented")
}

// List returns all existing relays
func (r *Relays) List(ctx context.Context, req RelayListRequest) (RelayListResponse, error) {
	// Host: app.strongdm.com
	// BasePath: 
	// Dial to /v1/relays by GET
	// Body: 
	// ResponseBody: 
	panic("not implemented")
}

// Get finds a sandwich by id
func (r *Relays) Get(ctx context.Context, req RelayGetRequest) (RelayGetResponse, error) {
	// Host: app.strongdm.com
	// BasePath: 
	// Dial to /v1/relays/{id} by GET
	// Body: 
	// ResponseBody: 
	panic("not implemented")
}
// RequestMetadata
type RequestMetadata struct {
	Page string `json:"page"`
}

// ResponseMetadata
type ResponseMetadata struct {
	NextPage string `json:"nextPage"`
	Found int64 `json:"found"`
	Affected int64 `json:"affected"`
}

type RelayCreateRequest struct {
	Metadata RequestMetadata `json:"metadata"`
	Node Node `json:"node"`
}

type RelayCreateResponse struct {
	Metadata ResponseMetadata `json:"metadata"`
	Node Node `json:"node"`
}

type RelayUpdateRequest struct {
	Metadata RequestMetadata `json:"metadata"`
	Node Node `json:"node"`
}

type RelayUpdateResponse struct {
	Metadata ResponseMetadata `json:"metadata"`
	Node Node `json:"node"`
}

type RelayDeleteRequest struct {
	Metadata RequestMetadata `json:"metadata"`
	ID string `json:"id"`
}

type RelayDeleteResponse struct {
	Metadata ResponseMetadata `json:"metadata"`
}

type RelayListRequest struct {
	Metadata RequestMetadata `json:"metadata"`
}

type RelayListResponse struct {
	Metadata ResponseMetadata `json:"metadata"`
	Nodes Node `json:"nodes"`
}

type RelayGetRequest struct {
	Metadata RequestMetadata `json:"metadata"`
	ID string `json:"id"`
}

type RelayGetResponse struct {
	Metadata ResponseMetadata `json:"metadata"`
	Node Node `json:"node"`
}

type ErrorDetail struct {
	Message string `json:"message"`
}

// Node is a domain object.
type Node struct {
	// id is the unique ID for this relay.
	ID string `json:"id"`
	Relay Relay `json:"relay"`
	Gateway Gateway `json:"gateway"`
}

type Relay struct {
	// name is the human readable unique name for this relay.
	Name string `json:"name"`
}

type Gateway struct {
	// name is the human readable unique name for this relay.
	Name string `json:"name"`
	// listen_address represents the network address to which nodes should dial to.
	ListenAddress string `json:"listenAddress"`
}

func addTokenHeader(t string) func(*http.Request) {
	return func(r *http.Request) {
		r.Header.Set("Authorization", "Bearer "+t)
	}
}