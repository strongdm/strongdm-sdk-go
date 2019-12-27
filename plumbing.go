package sdm

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	proto "github.com/strongdm/strongdm-sdk-go/internal/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func timestampToPorcelain(t *timestamp.Timestamp) time.Time {
	if t == nil {
		return time.Time{}
	}
	res, _ := ptypes.Timestamp(t)
	return res
}

func timestampToPlumbing(t time.Time) *timestamp.Timestamp {
	res, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil
	}
	return res
}

func driverToPlumbing(porcelain Driver) *proto.Driver {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Driver{}

	switch v := porcelain.(type) {
	case *Mysql:
		plumbing.Driver = &proto.Driver_Mysql{Mysql: mysqlToPlumbing(v)}
	case *AuroraMysql:
		plumbing.Driver = &proto.Driver_AuroraMysql{AuroraMysql: auroraMysqlToPlumbing(v)}
	case *Clustrix:
		plumbing.Driver = &proto.Driver_Clustrix{Clustrix: clustrixToPlumbing(v)}
	case *Maria:
		plumbing.Driver = &proto.Driver_Maria{Maria: mariaToPlumbing(v)}
	case *Memsql:
		plumbing.Driver = &proto.Driver_Memsql{Memsql: memsqlToPlumbing(v)}
	case *Athena:
		plumbing.Driver = &proto.Driver_Athena{Athena: athenaToPlumbing(v)}
	}
	return plumbing
}

func driverToPorcelain(plumbing *proto.Driver) Driver {
	if plumbing.GetMysql() != nil {
		return mysqlToPorcelain(plumbing.GetMysql())
	}
	if plumbing.GetAuroraMysql() != nil {
		return auroraMysqlToPorcelain(plumbing.GetAuroraMysql())
	}
	if plumbing.GetClustrix() != nil {
		return clustrixToPorcelain(plumbing.GetClustrix())
	}
	if plumbing.GetMaria() != nil {
		return mariaToPorcelain(plumbing.GetMaria())
	}
	if plumbing.GetMemsql() != nil {
		return memsqlToPorcelain(plumbing.GetMemsql())
	}
	if plumbing.GetAthena() != nil {
		return athenaToPorcelain(plumbing.GetAthena())
	}
	return nil
}

func repeatedDriverToPlumbing(porcelains []Driver) []*proto.Driver {
	var items []*proto.Driver
	for _, porcelain := range porcelains {
		items = append(items, driverToPlumbing(porcelain))
	}
	return items
}

func repeatedDriverToPorcelain(plumbings []*proto.Driver) []Driver {
	var items []Driver
	for _, plumbing := range plumbings {
		items = append(items, driverToPorcelain(plumbing))
	}
	return items
}

func mysqlToPorcelain(plumbing *proto.Mysql) *Mysql {
	if plumbing == nil {
		return nil
	}
	porcelain := &Mysql{}
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.Port = plumbing.Port
	return porcelain
}

func mysqlToPlumbing(porcelain *Mysql) *proto.Mysql {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Mysql{}
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.Port = porcelain.Port
	return plumbing
}

func repeatedMysqlToPlumbing(porcelains []*Mysql) []*proto.Mysql {
	var items []*proto.Mysql
	for _, porcelain := range porcelains {
		items = append(items, mysqlToPlumbing(porcelain))
	}
	return items
}

func repeatedMysqlToPorcelain(plumbings []*proto.Mysql) []*Mysql {
	var items []*Mysql
	for _, plumbing := range plumbings {
		items = append(items, mysqlToPorcelain(plumbing))
	}
	return items
}

func auroraMysqlToPorcelain(plumbing *proto.AuroraMysql) *AuroraMysql {
	if plumbing == nil {
		return nil
	}
	porcelain := &AuroraMysql{}
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.Port = plumbing.Port
	return porcelain
}

func auroraMysqlToPlumbing(porcelain *AuroraMysql) *proto.AuroraMysql {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.AuroraMysql{}
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.Port = porcelain.Port
	return plumbing
}

func repeatedAuroraMysqlToPlumbing(porcelains []*AuroraMysql) []*proto.AuroraMysql {
	var items []*proto.AuroraMysql
	for _, porcelain := range porcelains {
		items = append(items, auroraMysqlToPlumbing(porcelain))
	}
	return items
}

func repeatedAuroraMysqlToPorcelain(plumbings []*proto.AuroraMysql) []*AuroraMysql {
	var items []*AuroraMysql
	for _, plumbing := range plumbings {
		items = append(items, auroraMysqlToPorcelain(plumbing))
	}
	return items
}

func clustrixToPorcelain(plumbing *proto.Clustrix) *Clustrix {
	if plumbing == nil {
		return nil
	}
	porcelain := &Clustrix{}
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.Port = plumbing.Port
	return porcelain
}

func clustrixToPlumbing(porcelain *Clustrix) *proto.Clustrix {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Clustrix{}
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.Port = porcelain.Port
	return plumbing
}

func repeatedClustrixToPlumbing(porcelains []*Clustrix) []*proto.Clustrix {
	var items []*proto.Clustrix
	for _, porcelain := range porcelains {
		items = append(items, clustrixToPlumbing(porcelain))
	}
	return items
}

func repeatedClustrixToPorcelain(plumbings []*proto.Clustrix) []*Clustrix {
	var items []*Clustrix
	for _, plumbing := range plumbings {
		items = append(items, clustrixToPorcelain(plumbing))
	}
	return items
}

func mariaToPorcelain(plumbing *proto.Maria) *Maria {
	if plumbing == nil {
		return nil
	}
	porcelain := &Maria{}
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.Port = plumbing.Port
	return porcelain
}

func mariaToPlumbing(porcelain *Maria) *proto.Maria {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Maria{}
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.Port = porcelain.Port
	return plumbing
}

func repeatedMariaToPlumbing(porcelains []*Maria) []*proto.Maria {
	var items []*proto.Maria
	for _, porcelain := range porcelains {
		items = append(items, mariaToPlumbing(porcelain))
	}
	return items
}

func repeatedMariaToPorcelain(plumbings []*proto.Maria) []*Maria {
	var items []*Maria
	for _, plumbing := range plumbings {
		items = append(items, mariaToPorcelain(plumbing))
	}
	return items
}

func memsqlToPorcelain(plumbing *proto.Memsql) *Memsql {
	if plumbing == nil {
		return nil
	}
	porcelain := &Memsql{}
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.Port = plumbing.Port
	return porcelain
}

func memsqlToPlumbing(porcelain *Memsql) *proto.Memsql {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Memsql{}
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.Port = porcelain.Port
	return plumbing
}

func repeatedMemsqlToPlumbing(porcelains []*Memsql) []*proto.Memsql {
	var items []*proto.Memsql
	for _, porcelain := range porcelains {
		items = append(items, memsqlToPlumbing(porcelain))
	}
	return items
}

func repeatedMemsqlToPorcelain(plumbings []*proto.Memsql) []*Memsql {
	var items []*Memsql
	for _, plumbing := range plumbings {
		items = append(items, memsqlToPorcelain(plumbing))
	}
	return items
}

func athenaToPorcelain(plumbing *proto.Athena) *Athena {
	if plumbing == nil {
		return nil
	}
	porcelain := &Athena{}
	porcelain.Hostname = plumbing.Hostname
	porcelain.AccessKey = plumbing.AccessKey
	porcelain.SecretAccessKey = plumbing.SecretAccessKey
	porcelain.Region = plumbing.Region
	porcelain.Output = plumbing.Output
	return porcelain
}

func athenaToPlumbing(porcelain *Athena) *proto.Athena {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Athena{}
	plumbing.Hostname = porcelain.Hostname
	plumbing.AccessKey = porcelain.AccessKey
	plumbing.SecretAccessKey = porcelain.SecretAccessKey
	plumbing.Region = porcelain.Region
	plumbing.Output = porcelain.Output
	return plumbing
}

func repeatedAthenaToPlumbing(porcelains []*Athena) []*proto.Athena {
	var items []*proto.Athena
	for _, porcelain := range porcelains {
		items = append(items, athenaToPlumbing(porcelain))
	}
	return items
}

func repeatedAthenaToPorcelain(plumbings []*proto.Athena) []*Athena {
	var items []*Athena
	for _, plumbing := range plumbings {
		items = append(items, athenaToPorcelain(plumbing))
	}
	return items
}

func createResponseMetadataToPorcelain(plumbing *proto.CreateResponseMetadata) *CreateResponseMetadata {
	if plumbing == nil {
		return nil
	}
	porcelain := &CreateResponseMetadata{}
	return porcelain
}

func createResponseMetadataToPlumbing(porcelain *CreateResponseMetadata) *proto.CreateResponseMetadata {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.CreateResponseMetadata{}
	return plumbing
}

func repeatedCreateResponseMetadataToPlumbing(porcelains []*CreateResponseMetadata) []*proto.CreateResponseMetadata {
	var items []*proto.CreateResponseMetadata
	for _, porcelain := range porcelains {
		items = append(items, createResponseMetadataToPlumbing(porcelain))
	}
	return items
}

func repeatedCreateResponseMetadataToPorcelain(plumbings []*proto.CreateResponseMetadata) []*CreateResponseMetadata {
	var items []*CreateResponseMetadata
	for _, plumbing := range plumbings {
		items = append(items, createResponseMetadataToPorcelain(plumbing))
	}
	return items
}

func getResponseMetadataToPorcelain(plumbing *proto.GetResponseMetadata) *GetResponseMetadata {
	if plumbing == nil {
		return nil
	}
	porcelain := &GetResponseMetadata{}
	return porcelain
}

func getResponseMetadataToPlumbing(porcelain *GetResponseMetadata) *proto.GetResponseMetadata {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.GetResponseMetadata{}
	return plumbing
}

func repeatedGetResponseMetadataToPlumbing(porcelains []*GetResponseMetadata) []*proto.GetResponseMetadata {
	var items []*proto.GetResponseMetadata
	for _, porcelain := range porcelains {
		items = append(items, getResponseMetadataToPlumbing(porcelain))
	}
	return items
}

func repeatedGetResponseMetadataToPorcelain(plumbings []*proto.GetResponseMetadata) []*GetResponseMetadata {
	var items []*GetResponseMetadata
	for _, plumbing := range plumbings {
		items = append(items, getResponseMetadataToPorcelain(plumbing))
	}
	return items
}

func updateResponseMetadataToPorcelain(plumbing *proto.UpdateResponseMetadata) *UpdateResponseMetadata {
	if plumbing == nil {
		return nil
	}
	porcelain := &UpdateResponseMetadata{}
	return porcelain
}

func updateResponseMetadataToPlumbing(porcelain *UpdateResponseMetadata) *proto.UpdateResponseMetadata {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.UpdateResponseMetadata{}
	return plumbing
}

func repeatedUpdateResponseMetadataToPlumbing(porcelains []*UpdateResponseMetadata) []*proto.UpdateResponseMetadata {
	var items []*proto.UpdateResponseMetadata
	for _, porcelain := range porcelains {
		items = append(items, updateResponseMetadataToPlumbing(porcelain))
	}
	return items
}

func repeatedUpdateResponseMetadataToPorcelain(plumbings []*proto.UpdateResponseMetadata) []*UpdateResponseMetadata {
	var items []*UpdateResponseMetadata
	for _, plumbing := range plumbings {
		items = append(items, updateResponseMetadataToPorcelain(plumbing))
	}
	return items
}

func deleteResponseMetadataToPorcelain(plumbing *proto.DeleteResponseMetadata) *DeleteResponseMetadata {
	if plumbing == nil {
		return nil
	}
	porcelain := &DeleteResponseMetadata{}
	return porcelain
}

func deleteResponseMetadataToPlumbing(porcelain *DeleteResponseMetadata) *proto.DeleteResponseMetadata {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.DeleteResponseMetadata{}
	return plumbing
}

func repeatedDeleteResponseMetadataToPlumbing(porcelains []*DeleteResponseMetadata) []*proto.DeleteResponseMetadata {
	var items []*proto.DeleteResponseMetadata
	for _, porcelain := range porcelains {
		items = append(items, deleteResponseMetadataToPlumbing(porcelain))
	}
	return items
}

func repeatedDeleteResponseMetadataToPorcelain(plumbings []*proto.DeleteResponseMetadata) []*DeleteResponseMetadata {
	var items []*DeleteResponseMetadata
	for _, plumbing := range plumbings {
		items = append(items, deleteResponseMetadataToPorcelain(plumbing))
	}
	return items
}

func rateLimitMetadataToPorcelain(plumbing *proto.RateLimitMetadata) *RateLimitMetadata {
	if plumbing == nil {
		return nil
	}
	porcelain := &RateLimitMetadata{}
	porcelain.Limit = plumbing.Limit
	porcelain.Remaining = plumbing.Remaining
	porcelain.ResetAt = timestampToPorcelain(plumbing.ResetAt)
	porcelain.Bucket = plumbing.Bucket
	return porcelain
}

func rateLimitMetadataToPlumbing(porcelain *RateLimitMetadata) *proto.RateLimitMetadata {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.RateLimitMetadata{}
	plumbing.Limit = porcelain.Limit
	plumbing.Remaining = porcelain.Remaining
	plumbing.ResetAt = timestampToPlumbing(porcelain.ResetAt)
	plumbing.Bucket = porcelain.Bucket
	return plumbing
}

func repeatedRateLimitMetadataToPlumbing(porcelains []*RateLimitMetadata) []*proto.RateLimitMetadata {
	var items []*proto.RateLimitMetadata
	for _, porcelain := range porcelains {
		items = append(items, rateLimitMetadataToPlumbing(porcelain))
	}
	return items
}

func repeatedRateLimitMetadataToPorcelain(plumbings []*proto.RateLimitMetadata) []*RateLimitMetadata {
	var items []*RateLimitMetadata
	for _, plumbing := range plumbings {
		items = append(items, rateLimitMetadataToPorcelain(plumbing))
	}
	return items
}

func nodeCreateResponseToPorcelain(plumbing *proto.NodeCreateResponse) *NodeCreateResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &NodeCreateResponse{}
	porcelain.Meta = createResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.Node = nodeToPorcelain(plumbing.Node)
	porcelain.Token = plumbing.Token
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func nodeCreateResponseToPlumbing(porcelain *NodeCreateResponse) *proto.NodeCreateResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.NodeCreateResponse{}
	plumbing.Meta = createResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.Node = nodeToPlumbing(porcelain.Node)
	plumbing.Token = porcelain.Token
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedNodeCreateResponseToPlumbing(porcelains []*NodeCreateResponse) []*proto.NodeCreateResponse {
	var items []*proto.NodeCreateResponse
	for _, porcelain := range porcelains {
		items = append(items, nodeCreateResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedNodeCreateResponseToPorcelain(plumbings []*proto.NodeCreateResponse) []*NodeCreateResponse {
	var items []*NodeCreateResponse
	for _, plumbing := range plumbings {
		items = append(items, nodeCreateResponseToPorcelain(plumbing))
	}
	return items
}

func nodeGetResponseToPorcelain(plumbing *proto.NodeGetResponse) *NodeGetResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &NodeGetResponse{}
	porcelain.Meta = getResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.Node = nodeToPorcelain(plumbing.Node)
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func nodeGetResponseToPlumbing(porcelain *NodeGetResponse) *proto.NodeGetResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.NodeGetResponse{}
	plumbing.Meta = getResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.Node = nodeToPlumbing(porcelain.Node)
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedNodeGetResponseToPlumbing(porcelains []*NodeGetResponse) []*proto.NodeGetResponse {
	var items []*proto.NodeGetResponse
	for _, porcelain := range porcelains {
		items = append(items, nodeGetResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedNodeGetResponseToPorcelain(plumbings []*proto.NodeGetResponse) []*NodeGetResponse {
	var items []*NodeGetResponse
	for _, plumbing := range plumbings {
		items = append(items, nodeGetResponseToPorcelain(plumbing))
	}
	return items
}

func nodeUpdateResponseToPorcelain(plumbing *proto.NodeUpdateResponse) *NodeUpdateResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &NodeUpdateResponse{}
	porcelain.Meta = updateResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.Node = nodeToPorcelain(plumbing.Node)
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func nodeUpdateResponseToPlumbing(porcelain *NodeUpdateResponse) *proto.NodeUpdateResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.NodeUpdateResponse{}
	plumbing.Meta = updateResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.Node = nodeToPlumbing(porcelain.Node)
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedNodeUpdateResponseToPlumbing(porcelains []*NodeUpdateResponse) []*proto.NodeUpdateResponse {
	var items []*proto.NodeUpdateResponse
	for _, porcelain := range porcelains {
		items = append(items, nodeUpdateResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedNodeUpdateResponseToPorcelain(plumbings []*proto.NodeUpdateResponse) []*NodeUpdateResponse {
	var items []*NodeUpdateResponse
	for _, plumbing := range plumbings {
		items = append(items, nodeUpdateResponseToPorcelain(plumbing))
	}
	return items
}

func nodeDeleteResponseToPorcelain(plumbing *proto.NodeDeleteResponse) *NodeDeleteResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &NodeDeleteResponse{}
	porcelain.Meta = deleteResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func nodeDeleteResponseToPlumbing(porcelain *NodeDeleteResponse) *proto.NodeDeleteResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.NodeDeleteResponse{}
	plumbing.Meta = deleteResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedNodeDeleteResponseToPlumbing(porcelains []*NodeDeleteResponse) []*proto.NodeDeleteResponse {
	var items []*proto.NodeDeleteResponse
	for _, porcelain := range porcelains {
		items = append(items, nodeDeleteResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedNodeDeleteResponseToPorcelain(plumbings []*proto.NodeDeleteResponse) []*NodeDeleteResponse {
	var items []*NodeDeleteResponse
	for _, plumbing := range plumbings {
		items = append(items, nodeDeleteResponseToPorcelain(plumbing))
	}
	return items
}

func nodeToPlumbing(porcelain Node) *proto.Node {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Node{}

	switch v := porcelain.(type) {
	case *Relay:
		plumbing.Node = &proto.Node_Relay{Relay: relayToPlumbing(v)}
	case *Gateway:
		plumbing.Node = &proto.Node_Gateway{Gateway: gatewayToPlumbing(v)}
	}
	return plumbing
}

func nodeToPorcelain(plumbing *proto.Node) Node {
	if plumbing.GetRelay() != nil {
		return relayToPorcelain(plumbing.GetRelay())
	}
	if plumbing.GetGateway() != nil {
		return gatewayToPorcelain(plumbing.GetGateway())
	}
	return nil
}

func repeatedNodeToPlumbing(porcelains []Node) []*proto.Node {
	var items []*proto.Node
	for _, porcelain := range porcelains {
		items = append(items, nodeToPlumbing(porcelain))
	}
	return items
}

func repeatedNodeToPorcelain(plumbings []*proto.Node) []Node {
	var items []Node
	for _, plumbing := range plumbings {
		items = append(items, nodeToPorcelain(plumbing))
	}
	return items
}

func relayToPorcelain(plumbing *proto.Relay) *Relay {
	if plumbing == nil {
		return nil
	}
	porcelain := &Relay{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.State = plumbing.State
	return porcelain
}

func relayToPlumbing(porcelain *Relay) *proto.Relay {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Relay{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.State = porcelain.State
	return plumbing
}

func repeatedRelayToPlumbing(porcelains []*Relay) []*proto.Relay {
	var items []*proto.Relay
	for _, porcelain := range porcelains {
		items = append(items, relayToPlumbing(porcelain))
	}
	return items
}

func repeatedRelayToPorcelain(plumbings []*proto.Relay) []*Relay {
	var items []*Relay
	for _, plumbing := range plumbings {
		items = append(items, relayToPorcelain(plumbing))
	}
	return items
}

func gatewayToPorcelain(plumbing *proto.Gateway) *Gateway {
	if plumbing == nil {
		return nil
	}
	porcelain := &Gateway{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.State = plumbing.State
	porcelain.ListenAddress = plumbing.ListenAddress
	porcelain.BindAddress = plumbing.BindAddress
	return porcelain
}

func gatewayToPlumbing(porcelain *Gateway) *proto.Gateway {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Gateway{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.State = porcelain.State
	plumbing.ListenAddress = porcelain.ListenAddress
	plumbing.BindAddress = porcelain.BindAddress
	return plumbing
}

func repeatedGatewayToPlumbing(porcelains []*Gateway) []*proto.Gateway {
	var items []*proto.Gateway
	for _, porcelain := range porcelains {
		items = append(items, gatewayToPlumbing(porcelain))
	}
	return items
}

func repeatedGatewayToPorcelain(plumbings []*proto.Gateway) []*Gateway {
	var items []*Gateway
	for _, plumbing := range plumbings {
		items = append(items, gatewayToPorcelain(plumbing))
	}
	return items
}

func resourceCreateResponseToPorcelain(plumbing *proto.ResourceCreateResponse) *ResourceCreateResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &ResourceCreateResponse{}
	porcelain.Meta = createResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.Resource = resourceToPorcelain(plumbing.Resource)
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func resourceCreateResponseToPlumbing(porcelain *ResourceCreateResponse) *proto.ResourceCreateResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.ResourceCreateResponse{}
	plumbing.Meta = createResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.Resource = resourceToPlumbing(porcelain.Resource)
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedResourceCreateResponseToPlumbing(porcelains []*ResourceCreateResponse) []*proto.ResourceCreateResponse {
	var items []*proto.ResourceCreateResponse
	for _, porcelain := range porcelains {
		items = append(items, resourceCreateResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedResourceCreateResponseToPorcelain(plumbings []*proto.ResourceCreateResponse) []*ResourceCreateResponse {
	var items []*ResourceCreateResponse
	for _, plumbing := range plumbings {
		items = append(items, resourceCreateResponseToPorcelain(plumbing))
	}
	return items
}

func resourceGetResponseToPorcelain(plumbing *proto.ResourceGetResponse) *ResourceGetResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &ResourceGetResponse{}
	porcelain.Meta = getResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.Resource = resourceToPorcelain(plumbing.Resource)
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func resourceGetResponseToPlumbing(porcelain *ResourceGetResponse) *proto.ResourceGetResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.ResourceGetResponse{}
	plumbing.Meta = getResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.Resource = resourceToPlumbing(porcelain.Resource)
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedResourceGetResponseToPlumbing(porcelains []*ResourceGetResponse) []*proto.ResourceGetResponse {
	var items []*proto.ResourceGetResponse
	for _, porcelain := range porcelains {
		items = append(items, resourceGetResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedResourceGetResponseToPorcelain(plumbings []*proto.ResourceGetResponse) []*ResourceGetResponse {
	var items []*ResourceGetResponse
	for _, plumbing := range plumbings {
		items = append(items, resourceGetResponseToPorcelain(plumbing))
	}
	return items
}

func resourceUpdateResponseToPorcelain(plumbing *proto.ResourceUpdateResponse) *ResourceUpdateResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &ResourceUpdateResponse{}
	porcelain.Meta = updateResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.Resource = resourceToPorcelain(plumbing.Resource)
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func resourceUpdateResponseToPlumbing(porcelain *ResourceUpdateResponse) *proto.ResourceUpdateResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.ResourceUpdateResponse{}
	plumbing.Meta = updateResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.Resource = resourceToPlumbing(porcelain.Resource)
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedResourceUpdateResponseToPlumbing(porcelains []*ResourceUpdateResponse) []*proto.ResourceUpdateResponse {
	var items []*proto.ResourceUpdateResponse
	for _, porcelain := range porcelains {
		items = append(items, resourceUpdateResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedResourceUpdateResponseToPorcelain(plumbings []*proto.ResourceUpdateResponse) []*ResourceUpdateResponse {
	var items []*ResourceUpdateResponse
	for _, plumbing := range plumbings {
		items = append(items, resourceUpdateResponseToPorcelain(plumbing))
	}
	return items
}

func resourceDeleteResponseToPorcelain(plumbing *proto.ResourceDeleteResponse) *ResourceDeleteResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &ResourceDeleteResponse{}
	porcelain.Meta = deleteResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func resourceDeleteResponseToPlumbing(porcelain *ResourceDeleteResponse) *proto.ResourceDeleteResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.ResourceDeleteResponse{}
	plumbing.Meta = deleteResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedResourceDeleteResponseToPlumbing(porcelains []*ResourceDeleteResponse) []*proto.ResourceDeleteResponse {
	var items []*proto.ResourceDeleteResponse
	for _, porcelain := range porcelains {
		items = append(items, resourceDeleteResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedResourceDeleteResponseToPorcelain(plumbings []*proto.ResourceDeleteResponse) []*ResourceDeleteResponse {
	var items []*ResourceDeleteResponse
	for _, plumbing := range plumbings {
		items = append(items, resourceDeleteResponseToPorcelain(plumbing))
	}
	return items
}

func resourceToPorcelain(plumbing *proto.Resource) *Resource {
	if plumbing == nil {
		return nil
	}
	porcelain := &Resource{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Healthy = plumbing.Healthy
	porcelain.Driver = driverToPorcelain(plumbing.Driver)
	return porcelain
}

func resourceToPlumbing(porcelain *Resource) *proto.Resource {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Resource{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Healthy = porcelain.Healthy
	plumbing.Driver = driverToPlumbing(porcelain.Driver)
	return plumbing
}

func repeatedResourceToPlumbing(porcelains []*Resource) []*proto.Resource {
	var items []*proto.Resource
	for _, porcelain := range porcelains {
		items = append(items, resourceToPlumbing(porcelain))
	}
	return items
}

func repeatedResourceToPorcelain(plumbings []*proto.Resource) []*Resource {
	var items []*Resource
	for _, plumbing := range plumbings {
		items = append(items, resourceToPorcelain(plumbing))
	}
	return items
}

func roleAttachmentCreateResponseToPorcelain(plumbing *proto.RoleAttachmentCreateResponse) *RoleAttachmentCreateResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &RoleAttachmentCreateResponse{}
	porcelain.Meta = createResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.RoleAttachment = roleAttachmentToPorcelain(plumbing.RoleAttachment)
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func roleAttachmentCreateResponseToPlumbing(porcelain *RoleAttachmentCreateResponse) *proto.RoleAttachmentCreateResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.RoleAttachmentCreateResponse{}
	plumbing.Meta = createResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.RoleAttachment = roleAttachmentToPlumbing(porcelain.RoleAttachment)
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedRoleAttachmentCreateResponseToPlumbing(porcelains []*RoleAttachmentCreateResponse) []*proto.RoleAttachmentCreateResponse {
	var items []*proto.RoleAttachmentCreateResponse
	for _, porcelain := range porcelains {
		items = append(items, roleAttachmentCreateResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedRoleAttachmentCreateResponseToPorcelain(plumbings []*proto.RoleAttachmentCreateResponse) []*RoleAttachmentCreateResponse {
	var items []*RoleAttachmentCreateResponse
	for _, plumbing := range plumbings {
		items = append(items, roleAttachmentCreateResponseToPorcelain(plumbing))
	}
	return items
}

func roleAttachmentGetResponseToPorcelain(plumbing *proto.RoleAttachmentGetResponse) *RoleAttachmentGetResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &RoleAttachmentGetResponse{}
	porcelain.Meta = getResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.RoleAttachment = roleAttachmentToPorcelain(plumbing.RoleAttachment)
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func roleAttachmentGetResponseToPlumbing(porcelain *RoleAttachmentGetResponse) *proto.RoleAttachmentGetResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.RoleAttachmentGetResponse{}
	plumbing.Meta = getResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.RoleAttachment = roleAttachmentToPlumbing(porcelain.RoleAttachment)
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedRoleAttachmentGetResponseToPlumbing(porcelains []*RoleAttachmentGetResponse) []*proto.RoleAttachmentGetResponse {
	var items []*proto.RoleAttachmentGetResponse
	for _, porcelain := range porcelains {
		items = append(items, roleAttachmentGetResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedRoleAttachmentGetResponseToPorcelain(plumbings []*proto.RoleAttachmentGetResponse) []*RoleAttachmentGetResponse {
	var items []*RoleAttachmentGetResponse
	for _, plumbing := range plumbings {
		items = append(items, roleAttachmentGetResponseToPorcelain(plumbing))
	}
	return items
}

func roleAttachmentDeleteResponseToPorcelain(plumbing *proto.RoleAttachmentDeleteResponse) *RoleAttachmentDeleteResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &RoleAttachmentDeleteResponse{}
	porcelain.Meta = deleteResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func roleAttachmentDeleteResponseToPlumbing(porcelain *RoleAttachmentDeleteResponse) *proto.RoleAttachmentDeleteResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.RoleAttachmentDeleteResponse{}
	plumbing.Meta = deleteResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedRoleAttachmentDeleteResponseToPlumbing(porcelains []*RoleAttachmentDeleteResponse) []*proto.RoleAttachmentDeleteResponse {
	var items []*proto.RoleAttachmentDeleteResponse
	for _, porcelain := range porcelains {
		items = append(items, roleAttachmentDeleteResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedRoleAttachmentDeleteResponseToPorcelain(plumbings []*proto.RoleAttachmentDeleteResponse) []*RoleAttachmentDeleteResponse {
	var items []*RoleAttachmentDeleteResponse
	for _, plumbing := range plumbings {
		items = append(items, roleAttachmentDeleteResponseToPorcelain(plumbing))
	}
	return items
}

func roleAttachmentToPorcelain(plumbing *proto.RoleAttachment) *RoleAttachment {
	if plumbing == nil {
		return nil
	}
	porcelain := &RoleAttachment{}
	porcelain.ID = plumbing.Id
	porcelain.CompositeRoleID = plumbing.CompositeRoleId
	porcelain.AttachedRoleID = plumbing.AttachedRoleId
	return porcelain
}

func roleAttachmentToPlumbing(porcelain *RoleAttachment) *proto.RoleAttachment {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.RoleAttachment{}
	plumbing.Id = porcelain.ID
	plumbing.CompositeRoleId = porcelain.CompositeRoleID
	plumbing.AttachedRoleId = porcelain.AttachedRoleID
	return plumbing
}

func repeatedRoleAttachmentToPlumbing(porcelains []*RoleAttachment) []*proto.RoleAttachment {
	var items []*proto.RoleAttachment
	for _, porcelain := range porcelains {
		items = append(items, roleAttachmentToPlumbing(porcelain))
	}
	return items
}

func repeatedRoleAttachmentToPorcelain(plumbings []*proto.RoleAttachment) []*RoleAttachment {
	var items []*RoleAttachment
	for _, plumbing := range plumbings {
		items = append(items, roleAttachmentToPorcelain(plumbing))
	}
	return items
}

func roleCreateResponseToPorcelain(plumbing *proto.RoleCreateResponse) *RoleCreateResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &RoleCreateResponse{}
	porcelain.Meta = createResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.Role = roleToPorcelain(plumbing.Role)
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func roleCreateResponseToPlumbing(porcelain *RoleCreateResponse) *proto.RoleCreateResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.RoleCreateResponse{}
	plumbing.Meta = createResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.Role = roleToPlumbing(porcelain.Role)
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedRoleCreateResponseToPlumbing(porcelains []*RoleCreateResponse) []*proto.RoleCreateResponse {
	var items []*proto.RoleCreateResponse
	for _, porcelain := range porcelains {
		items = append(items, roleCreateResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedRoleCreateResponseToPorcelain(plumbings []*proto.RoleCreateResponse) []*RoleCreateResponse {
	var items []*RoleCreateResponse
	for _, plumbing := range plumbings {
		items = append(items, roleCreateResponseToPorcelain(plumbing))
	}
	return items
}

func roleGetResponseToPorcelain(plumbing *proto.RoleGetResponse) *RoleGetResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &RoleGetResponse{}
	porcelain.Meta = getResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.Role = roleToPorcelain(plumbing.Role)
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func roleGetResponseToPlumbing(porcelain *RoleGetResponse) *proto.RoleGetResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.RoleGetResponse{}
	plumbing.Meta = getResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.Role = roleToPlumbing(porcelain.Role)
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedRoleGetResponseToPlumbing(porcelains []*RoleGetResponse) []*proto.RoleGetResponse {
	var items []*proto.RoleGetResponse
	for _, porcelain := range porcelains {
		items = append(items, roleGetResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedRoleGetResponseToPorcelain(plumbings []*proto.RoleGetResponse) []*RoleGetResponse {
	var items []*RoleGetResponse
	for _, plumbing := range plumbings {
		items = append(items, roleGetResponseToPorcelain(plumbing))
	}
	return items
}

func roleUpdateResponseToPorcelain(plumbing *proto.RoleUpdateResponse) *RoleUpdateResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &RoleUpdateResponse{}
	porcelain.Meta = updateResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.Role = roleToPorcelain(plumbing.Role)
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func roleUpdateResponseToPlumbing(porcelain *RoleUpdateResponse) *proto.RoleUpdateResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.RoleUpdateResponse{}
	plumbing.Meta = updateResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.Role = roleToPlumbing(porcelain.Role)
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedRoleUpdateResponseToPlumbing(porcelains []*RoleUpdateResponse) []*proto.RoleUpdateResponse {
	var items []*proto.RoleUpdateResponse
	for _, porcelain := range porcelains {
		items = append(items, roleUpdateResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedRoleUpdateResponseToPorcelain(plumbings []*proto.RoleUpdateResponse) []*RoleUpdateResponse {
	var items []*RoleUpdateResponse
	for _, plumbing := range plumbings {
		items = append(items, roleUpdateResponseToPorcelain(plumbing))
	}
	return items
}

func roleDeleteResponseToPorcelain(plumbing *proto.RoleDeleteResponse) *RoleDeleteResponse {
	if plumbing == nil {
		return nil
	}
	porcelain := &RoleDeleteResponse{}
	porcelain.Meta = deleteResponseMetadataToPorcelain(plumbing.Meta)
	porcelain.RateLimit = rateLimitMetadataToPorcelain(plumbing.RateLimit)
	return porcelain
}

func roleDeleteResponseToPlumbing(porcelain *RoleDeleteResponse) *proto.RoleDeleteResponse {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.RoleDeleteResponse{}
	plumbing.Meta = deleteResponseMetadataToPlumbing(porcelain.Meta)
	plumbing.RateLimit = rateLimitMetadataToPlumbing(porcelain.RateLimit)
	return plumbing
}

func repeatedRoleDeleteResponseToPlumbing(porcelains []*RoleDeleteResponse) []*proto.RoleDeleteResponse {
	var items []*proto.RoleDeleteResponse
	for _, porcelain := range porcelains {
		items = append(items, roleDeleteResponseToPlumbing(porcelain))
	}
	return items
}

func repeatedRoleDeleteResponseToPorcelain(plumbings []*proto.RoleDeleteResponse) []*RoleDeleteResponse {
	var items []*RoleDeleteResponse
	for _, plumbing := range plumbings {
		items = append(items, roleDeleteResponseToPorcelain(plumbing))
	}
	return items
}

func roleToPorcelain(plumbing *proto.Role) *Role {
	if plumbing == nil {
		return nil
	}
	porcelain := &Role{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Composite = plumbing.Composite
	return porcelain
}

func roleToPlumbing(porcelain *Role) *proto.Role {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Role{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Composite = porcelain.Composite
	return plumbing
}

func repeatedRoleToPlumbing(porcelains []*Role) []*proto.Role {
	var items []*proto.Role
	for _, porcelain := range porcelains {
		items = append(items, roleToPlumbing(porcelain))
	}
	return items
}

func repeatedRoleToPorcelain(plumbings []*proto.Role) []*Role {
	var items []*Role
	for _, plumbing := range plumbings {
		items = append(items, roleToPorcelain(plumbing))
	}
	return items
}

type rpcError struct {
	wrapped error
	code    int
}

func (e *rpcError) Error() string {
	return e.wrapped.Error()
}

func (e *rpcError) Unwrap() error {
	return e.wrapped
}

func (e *rpcError) Code() int {
	return e.code
}

func errorToPorcelain(err error) error {
	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.Canceled:
			return &ContextCanceledError{Wrapped: err}
		case codes.DeadlineExceeded:
			return &DeadlineExceededError{Wrapped: err}
		case codes.AlreadyExists:
			return &AlreadyExistsError{Message: s.Message()}
		case codes.NotFound:
			return &NotFoundError{Message: s.Message()}
		case codes.InvalidArgument:
			return &BadRequestError{Message: s.Message()}
		case codes.Unauthenticated:
			return &AuthenticationError{Message: s.Message()}
		case codes.PermissionDenied:
			return &PermissionError{Message: s.Message()}
		case codes.Internal:
			return &InternalError{Message: s.Message()}
		case codes.ResourceExhausted:
			for _, d := range s.Details() {
				if d, ok := d.(*proto.RateLimitMetadata); ok {
					return &RateLimitError{Message: s.Message(), RateLimit: rateLimitMetadataToPorcelain(d)}
				}
			}
		}
		return &rpcError{wrapped: err, code: int(s.Code())}
	}
	return &UnknownError{Wrapped: err}
}

type nodeIteratorImplFetchFunc func() ([]Node, bool, error)
type nodeIteratorImpl struct {
	buffer      []Node
	index       int
	hasNextPage bool
	err         error
	fetch       nodeIteratorImplFetchFunc
}

func newNodeIteratorImpl(f nodeIteratorImplFetchFunc) *nodeIteratorImpl {
	return &nodeIteratorImpl{
		hasNextPage: true,
		fetch:       f,
	}
}

func (n *nodeIteratorImpl) Next() bool {
	if n.index < len(n.buffer)-1 {
		n.index++
		return true
	}

	// reached end of buffer
	if !n.hasNextPage {
		return false
	}

	n.index = 0
	n.buffer, n.hasNextPage, n.err = n.fetch()
	return len(n.buffer) > 0
}

func (n *nodeIteratorImpl) Value() Node {
	if n.index >= len(n.buffer) {
		return nil
	}
	return n.buffer[n.index]
}

func (n *nodeIteratorImpl) Err() error {
	return n.err
}

type resourceIteratorImplFetchFunc func() ([]*Resource, bool, error)
type resourceIteratorImpl struct {
	buffer      []*Resource
	index       int
	hasNextPage bool
	err         error
	fetch       resourceIteratorImplFetchFunc
}

func newResourceIteratorImpl(f resourceIteratorImplFetchFunc) *resourceIteratorImpl {
	return &resourceIteratorImpl{
		hasNextPage: true,
		fetch:       f,
	}
}

func (r *resourceIteratorImpl) Next() bool {
	if r.index < len(r.buffer)-1 {
		r.index++
		return true
	}

	// reached end of buffer
	if !r.hasNextPage {
		return false
	}

	r.index = 0
	r.buffer, r.hasNextPage, r.err = r.fetch()
	return len(r.buffer) > 0
}

func (r *resourceIteratorImpl) Value() *Resource {
	if r.index >= len(r.buffer) {
		return nil
	}
	return r.buffer[r.index]
}

func (r *resourceIteratorImpl) Err() error {
	return r.err
}

type roleAttachmentIteratorImplFetchFunc func() ([]*RoleAttachment, bool, error)
type roleAttachmentIteratorImpl struct {
	buffer      []*RoleAttachment
	index       int
	hasNextPage bool
	err         error
	fetch       roleAttachmentIteratorImplFetchFunc
}

func newRoleAttachmentIteratorImpl(f roleAttachmentIteratorImplFetchFunc) *roleAttachmentIteratorImpl {
	return &roleAttachmentIteratorImpl{
		hasNextPage: true,
		fetch:       f,
	}
}

func (r *roleAttachmentIteratorImpl) Next() bool {
	if r.index < len(r.buffer)-1 {
		r.index++
		return true
	}

	// reached end of buffer
	if !r.hasNextPage {
		return false
	}

	r.index = 0
	r.buffer, r.hasNextPage, r.err = r.fetch()
	return len(r.buffer) > 0
}

func (r *roleAttachmentIteratorImpl) Value() *RoleAttachment {
	if r.index >= len(r.buffer) {
		return nil
	}
	return r.buffer[r.index]
}

func (r *roleAttachmentIteratorImpl) Err() error {
	return r.err
}

type roleIteratorImplFetchFunc func() ([]*Role, bool, error)
type roleIteratorImpl struct {
	buffer      []*Role
	index       int
	hasNextPage bool
	err         error
	fetch       roleIteratorImplFetchFunc
}

func newRoleIteratorImpl(f roleIteratorImplFetchFunc) *roleIteratorImpl {
	return &roleIteratorImpl{
		hasNextPage: true,
		fetch:       f,
	}
}

func (r *roleIteratorImpl) Next() bool {
	if r.index < len(r.buffer)-1 {
		r.index++
		return true
	}

	// reached end of buffer
	if !r.hasNextPage {
		return false
	}

	r.index = 0
	r.buffer, r.hasNextPage, r.err = r.fetch()
	return len(r.buffer) > 0
}

func (r *roleIteratorImpl) Value() *Role {
	if r.index >= len(r.buffer) {
		return nil
	}
	return r.buffer[r.index]
}

func (r *roleIteratorImpl) Err() error {
	return r.err
}
