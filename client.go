/*
Package sdm implements an API client to strongDM restful API.
*/
package sdm

import (
	"crypto/tls"
	"net"
	"fmt"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	plumbing "github.com/strongdm/strongdm-sdk-go/internal/v1"
)

var (
	host = "app.strongdm.com"
	_ = metadata.Pairs
)

// Client is the strongDM API client implementation.
type Client struct {
	testOptionsMu sync.RWMutex
	testOptions   map[string]interface{}

	grpcConn *grpc.ClientConn
	// Nodes are proxies in strongDM responsible to communicate with servers
	// (relays) and clients (gateways).
	nodes *Nodes
	// Roles are tools for controlling user access to resources. Each role holds a
	// list of resources which they grant access to. Composite roles are a special
	// type of role which have no resource associations of their own, but instead
	// grant access to the combined resources associated with a set of child roles.
	// Each user can be a member of one role or composite role.
	roles *Roles
}

// New creates a new strongDM API client.
func New(host string, key string) (*Client, error) {
	var opts []grpc.DialOption
	
	_, port, err := net.SplitHostPort(host)
	if err != nil {
		return nil, plumbing.ErrorToPorcelain(fmt.Errorf("cannot parse host and port: %w", err))
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
		return nil, plumbing.ErrorToPorcelain(fmt.Errorf("cannot dial API server: %w", err))
	}
	client := &Client{
		grpcConn: cc,
		testOptions: map[string]interface{}{},
	}
	
	client.nodes = &Nodes{
		apiToken: key,
		client: plumbing.NewNodesClient(client.grpcConn),
		parent: client,
	}
	
	client.roles = &Roles{
		apiToken: key,
		client: plumbing.NewRolesClient(client.grpcConn),
		parent: client,
	}
	
	return client, nil
}

func (c *Client) Nodes() *Nodes{
	return c.nodes
}
func (c *Client) Roles() *Roles{
	return c.roles
}


func (c *Client) testOption(key string) interface{} {
	c.testOptionsMu.RLock()
	defer c.testOptionsMu.RUnlock()
	return c.testOptions[key]
}