package sdm

import (
	"context"

	v1 "github.com/strongdm/strongdm-sdk-go/internal/v1"
)


// Relays service
type Relays struct {
	client v1.RelaysClient
}

// Create adds a new relay, and returns the new relay
func (r *Relays) Create(ctx context.Context, req *v1.RelayCreateRequest) (*v1.RelayCreateResponse, error) {
	resp, err := r.Create(context.Background(), req)
	// TODO: improve error handling
	return resp, err
}

// Update modifies an existing relay
func (r *Relays) Update(ctx context.Context, req *v1.RelayUpdateRequest) (*v1.RelayUpdateResponse, error) {
	resp, err := r.Update(context.Background(), req)
	// TODO: improve error handling
	return resp, err
}

// Delete removes an existing relay
func (r *Relays) Delete(ctx context.Context, req *v1.RelayDeleteRequest) (*v1.RelayDeleteResponse, error) {
	resp, err := r.Delete(context.Background(), req)
	// TODO: improve error handling
	return resp, err
}

// List returns all existing relays
func (r *Relays) List(ctx context.Context, req *v1.RelayListRequest) (*v1.RelayListResponse, error) {
	resp, err := r.List(context.Background(), req)
	// TODO: improve error handling
	return resp, err
}

// Get finds a sandwich by id
func (r *Relays) Get(ctx context.Context, req *v1.RelayGetRequest) (*v1.RelayGetResponse, error) {
	resp, err := r.Get(context.Background(), req)
	// TODO: improve error handling
	return resp, err
}
