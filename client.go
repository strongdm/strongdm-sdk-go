/*
Package sdm implements an API client to strongDM restful API.
*/
package sdm

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	plumbing "github.com/strongdm/strongdm-sdk-go/internal/v1"
)

var (
	host = "app.strongdm.com"
	_ = metadata.Pairs
)

// Client is the strongDM API client.
type Client struct {
	grpcConn *grpc.ClientConn


	// Nodes are proxies in strongDM responsible to communicate with servers
	// (relays) and clients (gateways).
	Nodes *Nodes

	// Roles are
	Roles *Roles

}

// New creates a new strongDM API client.
func New(host string) (*Client, error) {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	
	cc, err := grpc.Dial(
		host,
		opts...,
	)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(fmt.Errorf("cannot dial API server: %w", err))
	}
	client := &Client{
		grpcConn: cc,
	}
	
	client.Nodes = &Nodes{client: plumbing.NewNodesClient(client.grpcConn),}
	
	client.Roles = &Roles{client: plumbing.NewRolesClient(client.grpcConn),}
	
	return client, nil
}
