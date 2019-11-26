/*
Package sdm implements an API client to strongDM restful API.
*/
package sdm

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"sync"

	plumbing "github.com/strongdm/strongdm-sdk-go/internal/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

var (
	host = "api.strongdm.com"
	_    = metadata.Pairs
)

// Client is the strongDM API client implementation.
type Client struct {
	testOptionsMu sync.RWMutex
	testOptions   map[string]interface{}

	apiToken string
	grpcConn *grpc.ClientConn
	nodes    *Nodes
	roles    *Roles
}

// New creates a new strongDM API client.
func New(host string, key string) (*Client, error) {
	var opts []grpc.DialOption

	_, port, err := net.SplitHostPort(host)
	if err != nil {
		return nil, errorToPorcelain(fmt.Errorf("cannot parse host and port: %w", err))
	}

	if port == "443" {
		tlsOpt := grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			RootCAs:            nil,
			InsecureSkipVerify: false,
		}))
		opts = append(opts, tlsOpt)
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	cc, err := grpc.Dial(
		host,
		opts...,
	)
	if err != nil {
		return nil, errorToPorcelain(fmt.Errorf("cannot dial API server: %w", err))
	}
	client := &Client{
		grpcConn:    cc,
		testOptions: map[string]interface{}{},
		apiToken:    key,
	}

	client.nodes = &Nodes{
		client: plumbing.NewNodesClient(client.grpcConn),
		parent: client,
	}

	client.roles = &Roles{
		client: plumbing.NewRolesClient(client.grpcConn),
		parent: client,
	}

	return client, nil
}

func (c *Client) wrapContext(ctx context.Context) context.Context {
	return createGRPCContext(ctx, c.apiToken)
}

// Nodes are proxies in the strongDM network. They come in two flavors: relays,
// which communicate with resources, and gateways, which communicate with
// clients.
func (c *Client) Nodes() *Nodes {
	return c.nodes
}

// Roles are tools for controlling user access to resources. Each Role holds a
// list of resources which they grant access to. Composite roles are a special
// type of Role which have no resource associations of their own, but instead
// grant access to the combined resources associated with a set of child roles.
// Each user can be a member of one Role or composite role.
func (c *Client) Roles() *Roles {
	return c.roles
}

func (c *Client) testOption(key string) interface{} {
	c.testOptionsMu.RLock()
	defer c.testOptionsMu.RUnlock()
	return c.testOptions[key]
}
