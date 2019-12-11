/*
Package sdm implements an API client to strongDM restful API.
*/
package sdm

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
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

	apiToken    string
	apiSecret   []byte
	grpcConn    *grpc.ClientConn
	attachments *Attachments
	nodes       *Nodes
	roles       *Roles
}

// New creates a new strongDM API client.
func New(host, token, secret string) (*Client, error) {
	var opts []grpc.DialOption

	_, port, err := net.SplitHostPort(host)
	if err != nil {
		return nil, errorToPorcelain(fmt.Errorf("cannot parse host and port: %w", err))
	}

	decodedSecret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return nil, errorToPorcelain(fmt.Errorf("invalid secret: %w", err))
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
		apiToken:    token,
		apiSecret:   decodedSecret,
	}

	client.attachments = &Attachments{
		client: plumbing.NewAttachmentsClient(client.grpcConn),
		parent: client,
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

func (c *Client) Attachments() *Attachments {
	return c.attachments
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

// Sign returns the signature for the given byte array
func (c *Client) Sign(message []byte) string {
	// Current UTC date
	y, m, d := time.Now().Date()
	currentDate := time.Date(y, m, d, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)

	signingKey := hmacHelper(c.apiSecret, []byte(currentDate))
	signingKey = hmacHelper(signingKey, []byte("sdm_api_v1"))

	hash := sha256.New()
	hash.Write(message)
	hashedMessage := hash.Sum(nil)

	return base64.StdEncoding.EncodeToString(hmacHelper(signingKey, hashedMessage))
}

func hmacHelper(key, msg []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)
	return mac.Sum(nil)
}

func (c *Client) wrapContext(ctx context.Context, req proto.Message) context.Context {
	msg, _ := proto.Marshal(req)
	return metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		"authorization":   c.apiToken,
		"x-sdm-signature": c.Sign(msg),
	}))
}

func (c *Client) testOption(key string) interface{} {
	c.testOptionsMu.RLock()
	defer c.testOptionsMu.RUnlock()
	return c.testOptions[key]
}
