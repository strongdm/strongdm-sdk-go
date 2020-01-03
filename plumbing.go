package sdm

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	proto "github.com/strongdm/strongdm-sdk-go/internal/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

func quoteFilterArgs(filter string, args ...interface{}) (string, error) {
	parts := strings.Split(filter, "?")
	if len(parts) != len(args)+1 {
		return "", &BadRequestError{Message: "incorrect number of replacements"}
	}
	var b strings.Builder
	for i, v := range parts {
		b.WriteString(v)
		if i < len(args) {
			s := fmt.Sprint(args[i])
			j, err := json.Marshal(s)
			if err != nil {
				return "", &BadRequestError{Message: "unable to marshal string to JSON"}
			}
			b.Write(j)
		}
	}
	return b.String(), nil
}

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

func resourceToPlumbing(porcelain Resource) *proto.Resource {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Resource{}

	switch v := porcelain.(type) {
	case *Sybase:
		plumbing.Resource = &proto.Resource_Sybase{Sybase: sybaseToPlumbing(v)}
	case *Presto:
		plumbing.Resource = &proto.Resource_Presto{Presto: prestoToPlumbing(v)}
	case *Teradata:
		plumbing.Resource = &proto.Resource_Teradata{Teradata: teradataToPlumbing(v)}
	case *AmazonES:
		plumbing.Resource = &proto.Resource_AmazonEs{AmazonEs: amazonEsToPlumbing(v)}
	case *Elastic:
		plumbing.Resource = &proto.Resource_Elastic{Elastic: elasticToPlumbing(v)}
	case *Redis:
		plumbing.Resource = &proto.Resource_Redis{Redis: redisToPlumbing(v)}
	case *ElasticacheRedis:
		plumbing.Resource = &proto.Resource_ElasticacheRedis{ElasticacheRedis: elasticacheRedisToPlumbing(v)}
	case *Kubernetes:
		plumbing.Resource = &proto.Resource_Kubernetes{Kubernetes: kubernetesToPlumbing(v)}
	case *KubernetesBasicAuth:
		plumbing.Resource = &proto.Resource_KubernetesBasicAuth{KubernetesBasicAuth: kubernetesBasicAuthToPlumbing(v)}
	case *AmazonEKS:
		plumbing.Resource = &proto.Resource_AmazonEks{AmazonEks: amazonEksToPlumbing(v)}
	case *GoogleGKE:
		plumbing.Resource = &proto.Resource_GoogleGke{GoogleGke: googleGkeToPlumbing(v)}
	case *Oracle:
		plumbing.Resource = &proto.Resource_Oracle{Oracle: oracleToPlumbing(v)}
	case *DynamoDB:
		plumbing.Resource = &proto.Resource_DynamoDb{DynamoDb: dynamoDbToPlumbing(v)}
	case *RDP:
		plumbing.Resource = &proto.Resource_Rdp{Rdp: rdpToPlumbing(v)}
	case *BigQuery:
		plumbing.Resource = &proto.Resource_BigQuery{BigQuery: bigQueryToPlumbing(v)}
	case *Snowflake:
		plumbing.Resource = &proto.Resource_Snowflake{Snowflake: snowflakeToPlumbing(v)}
	case *Memcached:
		plumbing.Resource = &proto.Resource_Memcached{Memcached: memcachedToPlumbing(v)}
	case *Postgres:
		plumbing.Resource = &proto.Resource_Postgres{Postgres: postgresToPlumbing(v)}
	case *AuroraPostgres:
		plumbing.Resource = &proto.Resource_AuroraPostgres{AuroraPostgres: auroraPostgresToPlumbing(v)}
	case *Greenplum:
		plumbing.Resource = &proto.Resource_Greenplum{Greenplum: greenplumToPlumbing(v)}
	case *Cockroach:
		plumbing.Resource = &proto.Resource_Cockroach{Cockroach: cockroachToPlumbing(v)}
	case *Redshift:
		plumbing.Resource = &proto.Resource_Redshift{Redshift: redshiftToPlumbing(v)}
	case *SSH:
		plumbing.Resource = &proto.Resource_Ssh{Ssh: sshToPlumbing(v)}
	case *HTTPBasicAuth:
		plumbing.Resource = &proto.Resource_HttpBasicAuth{HttpBasicAuth: httpBasicAuthToPlumbing(v)}
	case *HTTPNoAuth:
		plumbing.Resource = &proto.Resource_HttpNoAuth{HttpNoAuth: httpNoAuthToPlumbing(v)}
	case *HTTPAuth:
		plumbing.Resource = &proto.Resource_HttpAuth{HttpAuth: httpAuthToPlumbing(v)}
	case *Cassandra:
		plumbing.Resource = &proto.Resource_Cassandra{Cassandra: cassandraToPlumbing(v)}
	case *Mysql:
		plumbing.Resource = &proto.Resource_Mysql{Mysql: mysqlToPlumbing(v)}
	case *AuroraMysql:
		plumbing.Resource = &proto.Resource_AuroraMysql{AuroraMysql: auroraMysqlToPlumbing(v)}
	case *Clustrix:
		plumbing.Resource = &proto.Resource_Clustrix{Clustrix: clustrixToPlumbing(v)}
	case *Maria:
		plumbing.Resource = &proto.Resource_Maria{Maria: mariaToPlumbing(v)}
	case *Memsql:
		plumbing.Resource = &proto.Resource_Memsql{Memsql: memsqlToPlumbing(v)}
	case *Druid:
		plumbing.Resource = &proto.Resource_Druid{Druid: druidToPlumbing(v)}
	case *SQLServer:
		plumbing.Resource = &proto.Resource_SqlServer{SqlServer: sqlServerToPlumbing(v)}
	case *MongoHybrid:
		plumbing.Resource = &proto.Resource_MongoHybrid{MongoHybrid: mongoHybridToPlumbing(v)}
	case *MongoHost:
		plumbing.Resource = &proto.Resource_MongoHost{MongoHost: mongoHostToPlumbing(v)}
	case *MongoReplicaSet:
		plumbing.Resource = &proto.Resource_MongoReplicaSet{MongoReplicaSet: mongoReplicaSetToPlumbing(v)}
	case *Athena:
		plumbing.Resource = &proto.Resource_Athena{Athena: athenaToPlumbing(v)}
	}
	return plumbing
}

func resourceToPorcelain(plumbing *proto.Resource) Resource {
	if plumbing.GetSybase() != nil {
		return sybaseToPorcelain(plumbing.GetSybase())
	}
	if plumbing.GetPresto() != nil {
		return prestoToPorcelain(plumbing.GetPresto())
	}
	if plumbing.GetTeradata() != nil {
		return teradataToPorcelain(plumbing.GetTeradata())
	}
	if plumbing.GetAmazonEs() != nil {
		return amazonEsToPorcelain(plumbing.GetAmazonEs())
	}
	if plumbing.GetElastic() != nil {
		return elasticToPorcelain(plumbing.GetElastic())
	}
	if plumbing.GetRedis() != nil {
		return redisToPorcelain(plumbing.GetRedis())
	}
	if plumbing.GetElasticacheRedis() != nil {
		return elasticacheRedisToPorcelain(plumbing.GetElasticacheRedis())
	}
	if plumbing.GetKubernetes() != nil {
		return kubernetesToPorcelain(plumbing.GetKubernetes())
	}
	if plumbing.GetKubernetesBasicAuth() != nil {
		return kubernetesBasicAuthToPorcelain(plumbing.GetKubernetesBasicAuth())
	}
	if plumbing.GetAmazonEks() != nil {
		return amazonEksToPorcelain(plumbing.GetAmazonEks())
	}
	if plumbing.GetGoogleGke() != nil {
		return googleGkeToPorcelain(plumbing.GetGoogleGke())
	}
	if plumbing.GetOracle() != nil {
		return oracleToPorcelain(plumbing.GetOracle())
	}
	if plumbing.GetDynamoDb() != nil {
		return dynamoDbToPorcelain(plumbing.GetDynamoDb())
	}
	if plumbing.GetRdp() != nil {
		return rdpToPorcelain(plumbing.GetRdp())
	}
	if plumbing.GetBigQuery() != nil {
		return bigQueryToPorcelain(plumbing.GetBigQuery())
	}
	if plumbing.GetSnowflake() != nil {
		return snowflakeToPorcelain(plumbing.GetSnowflake())
	}
	if plumbing.GetMemcached() != nil {
		return memcachedToPorcelain(plumbing.GetMemcached())
	}
	if plumbing.GetPostgres() != nil {
		return postgresToPorcelain(plumbing.GetPostgres())
	}
	if plumbing.GetAuroraPostgres() != nil {
		return auroraPostgresToPorcelain(plumbing.GetAuroraPostgres())
	}
	if plumbing.GetGreenplum() != nil {
		return greenplumToPorcelain(plumbing.GetGreenplum())
	}
	if plumbing.GetCockroach() != nil {
		return cockroachToPorcelain(plumbing.GetCockroach())
	}
	if plumbing.GetRedshift() != nil {
		return redshiftToPorcelain(plumbing.GetRedshift())
	}
	if plumbing.GetSsh() != nil {
		return sshToPorcelain(plumbing.GetSsh())
	}
	if plumbing.GetHttpBasicAuth() != nil {
		return httpBasicAuthToPorcelain(plumbing.GetHttpBasicAuth())
	}
	if plumbing.GetHttpNoAuth() != nil {
		return httpNoAuthToPorcelain(plumbing.GetHttpNoAuth())
	}
	if plumbing.GetHttpAuth() != nil {
		return httpAuthToPorcelain(plumbing.GetHttpAuth())
	}
	if plumbing.GetCassandra() != nil {
		return cassandraToPorcelain(plumbing.GetCassandra())
	}
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
	if plumbing.GetDruid() != nil {
		return druidToPorcelain(plumbing.GetDruid())
	}
	if plumbing.GetSqlServer() != nil {
		return sqlServerToPorcelain(plumbing.GetSqlServer())
	}
	if plumbing.GetMongoHybrid() != nil {
		return mongoHybridToPorcelain(plumbing.GetMongoHybrid())
	}
	if plumbing.GetMongoHost() != nil {
		return mongoHostToPorcelain(plumbing.GetMongoHost())
	}
	if plumbing.GetMongoReplicaSet() != nil {
		return mongoReplicaSetToPorcelain(plumbing.GetMongoReplicaSet())
	}
	if plumbing.GetAthena() != nil {
		return athenaToPorcelain(plumbing.GetAthena())
	}
	return nil
}

func repeatedResourceToPlumbing(porcelains []Resource) []*proto.Resource {
	var items []*proto.Resource
	for _, porcelain := range porcelains {
		items = append(items, resourceToPlumbing(porcelain))
	}
	return items
}

func repeatedResourceToPorcelain(plumbings []*proto.Resource) []Resource {
	var items []Resource
	for _, plumbing := range plumbings {
		items = append(items, resourceToPorcelain(plumbing))
	}
	return items
}

func sybaseToPorcelain(plumbing *proto.Sybase) *Sybase {
	if plumbing == nil {
		return nil
	}
	porcelain := &Sybase{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	porcelain.Password = plumbing.Password
	return porcelain
}

func sybaseToPlumbing(porcelain *Sybase) *proto.Sybase {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Sybase{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Port = porcelain.Port
	plumbing.Password = porcelain.Password
	return plumbing
}

func repeatedSybaseToPlumbing(porcelains []*Sybase) []*proto.Sybase {
	var items []*proto.Sybase
	for _, porcelain := range porcelains {
		items = append(items, sybaseToPlumbing(porcelain))
	}
	return items
}

func repeatedSybaseToPorcelain(plumbings []*proto.Sybase) []*Sybase {
	var items []*Sybase
	for _, plumbing := range plumbings {
		items = append(items, sybaseToPorcelain(plumbing))
	}
	return items
}

func prestoToPorcelain(plumbing *proto.Presto) *Presto {
	if plumbing == nil {
		return nil
	}
	porcelain := &Presto{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	porcelain.Username = plumbing.Username
	porcelain.TlsRequired = plumbing.TlsRequired
	return porcelain
}

func prestoToPlumbing(porcelain *Presto) *proto.Presto {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Presto{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Port = porcelain.Port
	plumbing.Username = porcelain.Username
	plumbing.TlsRequired = porcelain.TlsRequired
	return plumbing
}

func repeatedPrestoToPlumbing(porcelains []*Presto) []*proto.Presto {
	var items []*proto.Presto
	for _, porcelain := range porcelains {
		items = append(items, prestoToPlumbing(porcelain))
	}
	return items
}

func repeatedPrestoToPorcelain(plumbings []*proto.Presto) []*Presto {
	var items []*Presto
	for _, plumbing := range plumbings {
		items = append(items, prestoToPorcelain(plumbing))
	}
	return items
}

func teradataToPorcelain(plumbing *proto.Teradata) *Teradata {
	if plumbing == nil {
		return nil
	}
	porcelain := &Teradata{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	return porcelain
}

func teradataToPlumbing(porcelain *Teradata) *proto.Teradata {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Teradata{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Port = porcelain.Port
	return plumbing
}

func repeatedTeradataToPlumbing(porcelains []*Teradata) []*proto.Teradata {
	var items []*proto.Teradata
	for _, porcelain := range porcelains {
		items = append(items, teradataToPlumbing(porcelain))
	}
	return items
}

func repeatedTeradataToPorcelain(plumbings []*proto.Teradata) []*Teradata {
	var items []*Teradata
	for _, plumbing := range plumbings {
		items = append(items, teradataToPorcelain(plumbing))
	}
	return items
}

func amazonEsToPorcelain(plumbing *proto.AmazonES) *AmazonES {
	if plumbing == nil {
		return nil
	}
	porcelain := &AmazonES{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Endpoint = plumbing.Endpoint
	porcelain.AccessKey = plumbing.AccessKey
	porcelain.SecretAccessKey = plumbing.SecretAccessKey
	porcelain.Region = plumbing.Region
	porcelain.PortOverride = plumbing.PortOverride
	return porcelain
}

func amazonEsToPlumbing(porcelain *AmazonES) *proto.AmazonES {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.AmazonES{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Endpoint = porcelain.Endpoint
	plumbing.AccessKey = porcelain.AccessKey
	plumbing.SecretAccessKey = porcelain.SecretAccessKey
	plumbing.Region = porcelain.Region
	plumbing.PortOverride = porcelain.PortOverride
	return plumbing
}

func repeatedAmazonESToPlumbing(porcelains []*AmazonES) []*proto.AmazonES {
	var items []*proto.AmazonES
	for _, porcelain := range porcelains {
		items = append(items, amazonEsToPlumbing(porcelain))
	}
	return items
}

func repeatedAmazonESToPorcelain(plumbings []*proto.AmazonES) []*AmazonES {
	var items []*AmazonES
	for _, plumbing := range plumbings {
		items = append(items, amazonEsToPorcelain(plumbing))
	}
	return items
}

func elasticToPorcelain(plumbing *proto.Elastic) *Elastic {
	if plumbing == nil {
		return nil
	}
	porcelain := &Elastic{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	porcelain.TlsRequired = plumbing.TlsRequired
	return porcelain
}

func elasticToPlumbing(porcelain *Elastic) *proto.Elastic {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Elastic{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Port = porcelain.Port
	plumbing.TlsRequired = porcelain.TlsRequired
	return plumbing
}

func repeatedElasticToPlumbing(porcelains []*Elastic) []*proto.Elastic {
	var items []*proto.Elastic
	for _, porcelain := range porcelains {
		items = append(items, elasticToPlumbing(porcelain))
	}
	return items
}

func repeatedElasticToPorcelain(plumbings []*proto.Elastic) []*Elastic {
	var items []*Elastic
	for _, plumbing := range plumbings {
		items = append(items, elasticToPorcelain(plumbing))
	}
	return items
}

func redisToPorcelain(plumbing *proto.Redis) *Redis {
	if plumbing == nil {
		return nil
	}
	porcelain := &Redis{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Password = plumbing.Password
	porcelain.Port = plumbing.Port
	return porcelain
}

func redisToPlumbing(porcelain *Redis) *proto.Redis {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Redis{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Password = porcelain.Password
	plumbing.Port = porcelain.Port
	return plumbing
}

func repeatedRedisToPlumbing(porcelains []*Redis) []*proto.Redis {
	var items []*proto.Redis
	for _, porcelain := range porcelains {
		items = append(items, redisToPlumbing(porcelain))
	}
	return items
}

func repeatedRedisToPorcelain(plumbings []*proto.Redis) []*Redis {
	var items []*Redis
	for _, plumbing := range plumbings {
		items = append(items, redisToPorcelain(plumbing))
	}
	return items
}

func elasticacheRedisToPorcelain(plumbing *proto.ElasticacheRedis) *ElasticacheRedis {
	if plumbing == nil {
		return nil
	}
	porcelain := &ElasticacheRedis{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Password = plumbing.Password
	porcelain.Port = plumbing.Port
	porcelain.TlsRequired = plumbing.TlsRequired
	return porcelain
}

func elasticacheRedisToPlumbing(porcelain *ElasticacheRedis) *proto.ElasticacheRedis {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.ElasticacheRedis{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Password = porcelain.Password
	plumbing.Port = porcelain.Port
	plumbing.TlsRequired = porcelain.TlsRequired
	return plumbing
}

func repeatedElasticacheRedisToPlumbing(porcelains []*ElasticacheRedis) []*proto.ElasticacheRedis {
	var items []*proto.ElasticacheRedis
	for _, porcelain := range porcelains {
		items = append(items, elasticacheRedisToPlumbing(porcelain))
	}
	return items
}

func repeatedElasticacheRedisToPorcelain(plumbings []*proto.ElasticacheRedis) []*ElasticacheRedis {
	var items []*ElasticacheRedis
	for _, plumbing := range plumbings {
		items = append(items, elasticacheRedisToPorcelain(plumbing))
	}
	return items
}

func kubernetesToPorcelain(plumbing *proto.Kubernetes) *Kubernetes {
	if plumbing == nil {
		return nil
	}
	porcelain := &Kubernetes{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Port = plumbing.Port
	porcelain.CertificateAuthority = plumbing.CertificateAuthority
	porcelain.CertificateAuthorityFilename = plumbing.CertificateAuthorityFilename
	porcelain.ClientCertificate = plumbing.ClientCertificate
	porcelain.ClientCertificateFilename = plumbing.ClientCertificateFilename
	porcelain.ClientKey = plumbing.ClientKey
	porcelain.ClientKeyFilename = plumbing.ClientKeyFilename
	return porcelain
}

func kubernetesToPlumbing(porcelain *Kubernetes) *proto.Kubernetes {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Kubernetes{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Port = porcelain.Port
	plumbing.CertificateAuthority = porcelain.CertificateAuthority
	plumbing.CertificateAuthorityFilename = porcelain.CertificateAuthorityFilename
	plumbing.ClientCertificate = porcelain.ClientCertificate
	plumbing.ClientCertificateFilename = porcelain.ClientCertificateFilename
	plumbing.ClientKey = porcelain.ClientKey
	plumbing.ClientKeyFilename = porcelain.ClientKeyFilename
	return plumbing
}

func repeatedKubernetesToPlumbing(porcelains []*Kubernetes) []*proto.Kubernetes {
	var items []*proto.Kubernetes
	for _, porcelain := range porcelains {
		items = append(items, kubernetesToPlumbing(porcelain))
	}
	return items
}

func repeatedKubernetesToPorcelain(plumbings []*proto.Kubernetes) []*Kubernetes {
	var items []*Kubernetes
	for _, plumbing := range plumbings {
		items = append(items, kubernetesToPorcelain(plumbing))
	}
	return items
}

func kubernetesBasicAuthToPorcelain(plumbing *proto.KubernetesBasicAuth) *KubernetesBasicAuth {
	if plumbing == nil {
		return nil
	}
	porcelain := &KubernetesBasicAuth{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Port = plumbing.Port
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.CertificateAuthority = plumbing.CertificateAuthority
	porcelain.CertificateAuthorityFilename = plumbing.CertificateAuthorityFilename
	porcelain.ClientCertificate = plumbing.ClientCertificate
	porcelain.ClientCertificateFilename = plumbing.ClientCertificateFilename
	porcelain.ClientKey = plumbing.ClientKey
	porcelain.ClientKeyFilename = plumbing.ClientKeyFilename
	return porcelain
}

func kubernetesBasicAuthToPlumbing(porcelain *KubernetesBasicAuth) *proto.KubernetesBasicAuth {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.KubernetesBasicAuth{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Port = porcelain.Port
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.CertificateAuthority = porcelain.CertificateAuthority
	plumbing.CertificateAuthorityFilename = porcelain.CertificateAuthorityFilename
	plumbing.ClientCertificate = porcelain.ClientCertificate
	plumbing.ClientCertificateFilename = porcelain.ClientCertificateFilename
	plumbing.ClientKey = porcelain.ClientKey
	plumbing.ClientKeyFilename = porcelain.ClientKeyFilename
	return plumbing
}

func repeatedKubernetesBasicAuthToPlumbing(porcelains []*KubernetesBasicAuth) []*proto.KubernetesBasicAuth {
	var items []*proto.KubernetesBasicAuth
	for _, porcelain := range porcelains {
		items = append(items, kubernetesBasicAuthToPlumbing(porcelain))
	}
	return items
}

func repeatedKubernetesBasicAuthToPorcelain(plumbings []*proto.KubernetesBasicAuth) []*KubernetesBasicAuth {
	var items []*KubernetesBasicAuth
	for _, plumbing := range plumbings {
		items = append(items, kubernetesBasicAuthToPorcelain(plumbing))
	}
	return items
}

func amazonEksToPorcelain(plumbing *proto.AmazonEKS) *AmazonEKS {
	if plumbing == nil {
		return nil
	}
	porcelain := &AmazonEKS{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Endpoint = plumbing.Endpoint
	porcelain.AccessKey = plumbing.AccessKey
	porcelain.SecretAccessKey = plumbing.SecretAccessKey
	porcelain.CertificateAuthority = plumbing.CertificateAuthority
	porcelain.CertificateAuthorityFilename = plumbing.CertificateAuthorityFilename
	porcelain.Region = plumbing.Region
	porcelain.ClusterName = plumbing.ClusterName
	return porcelain
}

func amazonEksToPlumbing(porcelain *AmazonEKS) *proto.AmazonEKS {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.AmazonEKS{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Endpoint = porcelain.Endpoint
	plumbing.AccessKey = porcelain.AccessKey
	plumbing.SecretAccessKey = porcelain.SecretAccessKey
	plumbing.CertificateAuthority = porcelain.CertificateAuthority
	plumbing.CertificateAuthorityFilename = porcelain.CertificateAuthorityFilename
	plumbing.Region = porcelain.Region
	plumbing.ClusterName = porcelain.ClusterName
	return plumbing
}

func repeatedAmazonEKSToPlumbing(porcelains []*AmazonEKS) []*proto.AmazonEKS {
	var items []*proto.AmazonEKS
	for _, porcelain := range porcelains {
		items = append(items, amazonEksToPlumbing(porcelain))
	}
	return items
}

func repeatedAmazonEKSToPorcelain(plumbings []*proto.AmazonEKS) []*AmazonEKS {
	var items []*AmazonEKS
	for _, plumbing := range plumbings {
		items = append(items, amazonEksToPorcelain(plumbing))
	}
	return items
}

func googleGkeToPorcelain(plumbing *proto.GoogleGKE) *GoogleGKE {
	if plumbing == nil {
		return nil
	}
	porcelain := &GoogleGKE{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Endpoint = plumbing.Endpoint
	porcelain.CertificateAuthority = plumbing.CertificateAuthority
	porcelain.CertificateAuthorityFilename = plumbing.CertificateAuthorityFilename
	porcelain.ServiceAccountKey = plumbing.ServiceAccountKey
	porcelain.ServiceAccountKeyFilename = plumbing.ServiceAccountKeyFilename
	return porcelain
}

func googleGkeToPlumbing(porcelain *GoogleGKE) *proto.GoogleGKE {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.GoogleGKE{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Endpoint = porcelain.Endpoint
	plumbing.CertificateAuthority = porcelain.CertificateAuthority
	plumbing.CertificateAuthorityFilename = porcelain.CertificateAuthorityFilename
	plumbing.ServiceAccountKey = porcelain.ServiceAccountKey
	plumbing.ServiceAccountKeyFilename = porcelain.ServiceAccountKeyFilename
	return plumbing
}

func repeatedGoogleGKEToPlumbing(porcelains []*GoogleGKE) []*proto.GoogleGKE {
	var items []*proto.GoogleGKE
	for _, porcelain := range porcelains {
		items = append(items, googleGkeToPlumbing(porcelain))
	}
	return items
}

func repeatedGoogleGKEToPorcelain(plumbings []*proto.GoogleGKE) []*GoogleGKE {
	var items []*GoogleGKE
	for _, plumbing := range plumbings {
		items = append(items, googleGkeToPorcelain(plumbing))
	}
	return items
}

func oracleToPorcelain(plumbing *proto.Oracle) *Oracle {
	if plumbing == nil {
		return nil
	}
	porcelain := &Oracle{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.Port = plumbing.Port
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.TlsRequired = plumbing.TlsRequired
	return porcelain
}

func oracleToPlumbing(porcelain *Oracle) *proto.Oracle {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Oracle{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.Port = porcelain.Port
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.TlsRequired = porcelain.TlsRequired
	return plumbing
}

func repeatedOracleToPlumbing(porcelains []*Oracle) []*proto.Oracle {
	var items []*proto.Oracle
	for _, porcelain := range porcelains {
		items = append(items, oracleToPlumbing(porcelain))
	}
	return items
}

func repeatedOracleToPorcelain(plumbings []*proto.Oracle) []*Oracle {
	var items []*Oracle
	for _, plumbing := range plumbings {
		items = append(items, oracleToPorcelain(plumbing))
	}
	return items
}

func dynamoDbToPorcelain(plumbing *proto.DynamoDB) *DynamoDB {
	if plumbing == nil {
		return nil
	}
	porcelain := &DynamoDB{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Endpoint = plumbing.Endpoint
	porcelain.AccessKey = plumbing.AccessKey
	porcelain.SecretAccessKey = plumbing.SecretAccessKey
	porcelain.Region = plumbing.Region
	porcelain.PortOverride = plumbing.PortOverride
	return porcelain
}

func dynamoDbToPlumbing(porcelain *DynamoDB) *proto.DynamoDB {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.DynamoDB{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Endpoint = porcelain.Endpoint
	plumbing.AccessKey = porcelain.AccessKey
	plumbing.SecretAccessKey = porcelain.SecretAccessKey
	plumbing.Region = porcelain.Region
	plumbing.PortOverride = porcelain.PortOverride
	return plumbing
}

func repeatedDynamoDBToPlumbing(porcelains []*DynamoDB) []*proto.DynamoDB {
	var items []*proto.DynamoDB
	for _, porcelain := range porcelains {
		items = append(items, dynamoDbToPlumbing(porcelain))
	}
	return items
}

func repeatedDynamoDBToPorcelain(plumbings []*proto.DynamoDB) []*DynamoDB {
	var items []*DynamoDB
	for _, plumbing := range plumbings {
		items = append(items, dynamoDbToPorcelain(plumbing))
	}
	return items
}

func rdpToPorcelain(plumbing *proto.RDP) *RDP {
	if plumbing == nil {
		return nil
	}
	porcelain := &RDP{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	return porcelain
}

func rdpToPlumbing(porcelain *RDP) *proto.RDP {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.RDP{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Port = porcelain.Port
	return plumbing
}

func repeatedRDPToPlumbing(porcelains []*RDP) []*proto.RDP {
	var items []*proto.RDP
	for _, porcelain := range porcelains {
		items = append(items, rdpToPlumbing(porcelain))
	}
	return items
}

func repeatedRDPToPorcelain(plumbings []*proto.RDP) []*RDP {
	var items []*RDP
	for _, plumbing := range plumbings {
		items = append(items, rdpToPorcelain(plumbing))
	}
	return items
}

func bigQueryToPorcelain(plumbing *proto.BigQuery) *BigQuery {
	if plumbing == nil {
		return nil
	}
	porcelain := &BigQuery{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Endpoint = plumbing.Endpoint
	porcelain.PrivateKey = plumbing.PrivateKey
	porcelain.Project = plumbing.Project
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Username = plumbing.Username
	return porcelain
}

func bigQueryToPlumbing(porcelain *BigQuery) *proto.BigQuery {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.BigQuery{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Endpoint = porcelain.Endpoint
	plumbing.PrivateKey = porcelain.PrivateKey
	plumbing.Project = porcelain.Project
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Username = porcelain.Username
	return plumbing
}

func repeatedBigQueryToPlumbing(porcelains []*BigQuery) []*proto.BigQuery {
	var items []*proto.BigQuery
	for _, porcelain := range porcelains {
		items = append(items, bigQueryToPlumbing(porcelain))
	}
	return items
}

func repeatedBigQueryToPorcelain(plumbings []*proto.BigQuery) []*BigQuery {
	var items []*BigQuery
	for _, plumbing := range plumbings {
		items = append(items, bigQueryToPorcelain(plumbing))
	}
	return items
}

func snowflakeToPorcelain(plumbing *proto.Snowflake) *Snowflake {
	if plumbing == nil {
		return nil
	}
	porcelain := &Snowflake{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.Schema = plumbing.Schema
	porcelain.PortOverride = plumbing.PortOverride
	return porcelain
}

func snowflakeToPlumbing(porcelain *Snowflake) *proto.Snowflake {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Snowflake{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.Schema = porcelain.Schema
	plumbing.PortOverride = porcelain.PortOverride
	return plumbing
}

func repeatedSnowflakeToPlumbing(porcelains []*Snowflake) []*proto.Snowflake {
	var items []*proto.Snowflake
	for _, porcelain := range porcelains {
		items = append(items, snowflakeToPlumbing(porcelain))
	}
	return items
}

func repeatedSnowflakeToPorcelain(plumbings []*proto.Snowflake) []*Snowflake {
	var items []*Snowflake
	for _, plumbing := range plumbings {
		items = append(items, snowflakeToPorcelain(plumbing))
	}
	return items
}

func memcachedToPorcelain(plumbing *proto.Memcached) *Memcached {
	if plumbing == nil {
		return nil
	}
	porcelain := &Memcached{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	return porcelain
}

func memcachedToPlumbing(porcelain *Memcached) *proto.Memcached {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Memcached{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Port = porcelain.Port
	return plumbing
}

func repeatedMemcachedToPlumbing(porcelains []*Memcached) []*proto.Memcached {
	var items []*proto.Memcached
	for _, porcelain := range porcelains {
		items = append(items, memcachedToPlumbing(porcelain))
	}
	return items
}

func repeatedMemcachedToPorcelain(plumbings []*proto.Memcached) []*Memcached {
	var items []*Memcached
	for _, plumbing := range plumbings {
		items = append(items, memcachedToPorcelain(plumbing))
	}
	return items
}

func postgresToPorcelain(plumbing *proto.Postgres) *Postgres {
	if plumbing == nil {
		return nil
	}
	porcelain := &Postgres{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	porcelain.OverrideDatabase = plumbing.OverrideDatabase
	return porcelain
}

func postgresToPlumbing(porcelain *Postgres) *proto.Postgres {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Postgres{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Port = porcelain.Port
	plumbing.OverrideDatabase = porcelain.OverrideDatabase
	return plumbing
}

func repeatedPostgresToPlumbing(porcelains []*Postgres) []*proto.Postgres {
	var items []*proto.Postgres
	for _, porcelain := range porcelains {
		items = append(items, postgresToPlumbing(porcelain))
	}
	return items
}

func repeatedPostgresToPorcelain(plumbings []*proto.Postgres) []*Postgres {
	var items []*Postgres
	for _, plumbing := range plumbings {
		items = append(items, postgresToPorcelain(plumbing))
	}
	return items
}

func auroraPostgresToPorcelain(plumbing *proto.AuroraPostgres) *AuroraPostgres {
	if plumbing == nil {
		return nil
	}
	porcelain := &AuroraPostgres{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	porcelain.OverrideDatabase = plumbing.OverrideDatabase
	return porcelain
}

func auroraPostgresToPlumbing(porcelain *AuroraPostgres) *proto.AuroraPostgres {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.AuroraPostgres{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Port = porcelain.Port
	plumbing.OverrideDatabase = porcelain.OverrideDatabase
	return plumbing
}

func repeatedAuroraPostgresToPlumbing(porcelains []*AuroraPostgres) []*proto.AuroraPostgres {
	var items []*proto.AuroraPostgres
	for _, porcelain := range porcelains {
		items = append(items, auroraPostgresToPlumbing(porcelain))
	}
	return items
}

func repeatedAuroraPostgresToPorcelain(plumbings []*proto.AuroraPostgres) []*AuroraPostgres {
	var items []*AuroraPostgres
	for _, plumbing := range plumbings {
		items = append(items, auroraPostgresToPorcelain(plumbing))
	}
	return items
}

func greenplumToPorcelain(plumbing *proto.Greenplum) *Greenplum {
	if plumbing == nil {
		return nil
	}
	porcelain := &Greenplum{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	porcelain.OverrideDatabase = plumbing.OverrideDatabase
	return porcelain
}

func greenplumToPlumbing(porcelain *Greenplum) *proto.Greenplum {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Greenplum{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Port = porcelain.Port
	plumbing.OverrideDatabase = porcelain.OverrideDatabase
	return plumbing
}

func repeatedGreenplumToPlumbing(porcelains []*Greenplum) []*proto.Greenplum {
	var items []*proto.Greenplum
	for _, porcelain := range porcelains {
		items = append(items, greenplumToPlumbing(porcelain))
	}
	return items
}

func repeatedGreenplumToPorcelain(plumbings []*proto.Greenplum) []*Greenplum {
	var items []*Greenplum
	for _, plumbing := range plumbings {
		items = append(items, greenplumToPorcelain(plumbing))
	}
	return items
}

func cockroachToPorcelain(plumbing *proto.Cockroach) *Cockroach {
	if plumbing == nil {
		return nil
	}
	porcelain := &Cockroach{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	porcelain.OverrideDatabase = plumbing.OverrideDatabase
	return porcelain
}

func cockroachToPlumbing(porcelain *Cockroach) *proto.Cockroach {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Cockroach{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Port = porcelain.Port
	plumbing.OverrideDatabase = porcelain.OverrideDatabase
	return plumbing
}

func repeatedCockroachToPlumbing(porcelains []*Cockroach) []*proto.Cockroach {
	var items []*proto.Cockroach
	for _, porcelain := range porcelains {
		items = append(items, cockroachToPlumbing(porcelain))
	}
	return items
}

func repeatedCockroachToPorcelain(plumbings []*proto.Cockroach) []*Cockroach {
	var items []*Cockroach
	for _, plumbing := range plumbings {
		items = append(items, cockroachToPorcelain(plumbing))
	}
	return items
}

func redshiftToPorcelain(plumbing *proto.Redshift) *Redshift {
	if plumbing == nil {
		return nil
	}
	porcelain := &Redshift{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	porcelain.OverrideDatabase = plumbing.OverrideDatabase
	return porcelain
}

func redshiftToPlumbing(porcelain *Redshift) *proto.Redshift {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Redshift{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Port = porcelain.Port
	plumbing.OverrideDatabase = porcelain.OverrideDatabase
	return plumbing
}

func repeatedRedshiftToPlumbing(porcelains []*Redshift) []*proto.Redshift {
	var items []*proto.Redshift
	for _, porcelain := range porcelains {
		items = append(items, redshiftToPlumbing(porcelain))
	}
	return items
}

func repeatedRedshiftToPorcelain(plumbings []*proto.Redshift) []*Redshift {
	var items []*Redshift
	for _, plumbing := range plumbings {
		items = append(items, redshiftToPorcelain(plumbing))
	}
	return items
}

func sshToPorcelain(plumbing *proto.SSH) *SSH {
	if plumbing == nil {
		return nil
	}
	porcelain := &SSH{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Port = plumbing.Port
	porcelain.PublicKey = plumbing.PublicKey
	return porcelain
}

func sshToPlumbing(porcelain *SSH) *proto.SSH {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.SSH{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Port = porcelain.Port
	plumbing.PublicKey = porcelain.PublicKey
	return plumbing
}

func repeatedSSHToPlumbing(porcelains []*SSH) []*proto.SSH {
	var items []*proto.SSH
	for _, porcelain := range porcelains {
		items = append(items, sshToPlumbing(porcelain))
	}
	return items
}

func repeatedSSHToPorcelain(plumbings []*proto.SSH) []*SSH {
	var items []*SSH
	for _, plumbing := range plumbings {
		items = append(items, sshToPorcelain(plumbing))
	}
	return items
}

func httpBasicAuthToPorcelain(plumbing *proto.HTTPBasicAuth) *HTTPBasicAuth {
	if plumbing == nil {
		return nil
	}
	porcelain := &HTTPBasicAuth{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Url = plumbing.Url
	porcelain.HealthcheckPath = plumbing.HealthcheckPath
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.HeadersBlacklist = plumbing.HeadersBlacklist
	porcelain.DefaultPath = plumbing.DefaultPath
	porcelain.Subdomain = plumbing.Subdomain
	return porcelain
}

func httpBasicAuthToPlumbing(porcelain *HTTPBasicAuth) *proto.HTTPBasicAuth {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.HTTPBasicAuth{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Url = porcelain.Url
	plumbing.HealthcheckPath = porcelain.HealthcheckPath
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.HeadersBlacklist = porcelain.HeadersBlacklist
	plumbing.DefaultPath = porcelain.DefaultPath
	plumbing.Subdomain = porcelain.Subdomain
	return plumbing
}

func repeatedHTTPBasicAuthToPlumbing(porcelains []*HTTPBasicAuth) []*proto.HTTPBasicAuth {
	var items []*proto.HTTPBasicAuth
	for _, porcelain := range porcelains {
		items = append(items, httpBasicAuthToPlumbing(porcelain))
	}
	return items
}

func repeatedHTTPBasicAuthToPorcelain(plumbings []*proto.HTTPBasicAuth) []*HTTPBasicAuth {
	var items []*HTTPBasicAuth
	for _, plumbing := range plumbings {
		items = append(items, httpBasicAuthToPorcelain(plumbing))
	}
	return items
}

func httpNoAuthToPorcelain(plumbing *proto.HTTPNoAuth) *HTTPNoAuth {
	if plumbing == nil {
		return nil
	}
	porcelain := &HTTPNoAuth{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Url = plumbing.Url
	porcelain.HealthcheckPath = plumbing.HealthcheckPath
	porcelain.HeadersBlacklist = plumbing.HeadersBlacklist
	porcelain.DefaultPath = plumbing.DefaultPath
	porcelain.Subdomain = plumbing.Subdomain
	return porcelain
}

func httpNoAuthToPlumbing(porcelain *HTTPNoAuth) *proto.HTTPNoAuth {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.HTTPNoAuth{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Url = porcelain.Url
	plumbing.HealthcheckPath = porcelain.HealthcheckPath
	plumbing.HeadersBlacklist = porcelain.HeadersBlacklist
	plumbing.DefaultPath = porcelain.DefaultPath
	plumbing.Subdomain = porcelain.Subdomain
	return plumbing
}

func repeatedHTTPNoAuthToPlumbing(porcelains []*HTTPNoAuth) []*proto.HTTPNoAuth {
	var items []*proto.HTTPNoAuth
	for _, porcelain := range porcelains {
		items = append(items, httpNoAuthToPlumbing(porcelain))
	}
	return items
}

func repeatedHTTPNoAuthToPorcelain(plumbings []*proto.HTTPNoAuth) []*HTTPNoAuth {
	var items []*HTTPNoAuth
	for _, plumbing := range plumbings {
		items = append(items, httpNoAuthToPorcelain(plumbing))
	}
	return items
}

func httpAuthToPorcelain(plumbing *proto.HTTPAuth) *HTTPAuth {
	if plumbing == nil {
		return nil
	}
	porcelain := &HTTPAuth{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Url = plumbing.Url
	porcelain.HealthcheckPath = plumbing.HealthcheckPath
	porcelain.AuthHeader = plumbing.AuthHeader
	porcelain.HeadersBlacklist = plumbing.HeadersBlacklist
	porcelain.DefaultPath = plumbing.DefaultPath
	porcelain.Subdomain = plumbing.Subdomain
	return porcelain
}

func httpAuthToPlumbing(porcelain *HTTPAuth) *proto.HTTPAuth {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.HTTPAuth{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Url = porcelain.Url
	plumbing.HealthcheckPath = porcelain.HealthcheckPath
	plumbing.AuthHeader = porcelain.AuthHeader
	plumbing.HeadersBlacklist = porcelain.HeadersBlacklist
	plumbing.DefaultPath = porcelain.DefaultPath
	plumbing.Subdomain = porcelain.Subdomain
	return plumbing
}

func repeatedHTTPAuthToPlumbing(porcelains []*HTTPAuth) []*proto.HTTPAuth {
	var items []*proto.HTTPAuth
	for _, porcelain := range porcelains {
		items = append(items, httpAuthToPlumbing(porcelain))
	}
	return items
}

func repeatedHTTPAuthToPorcelain(plumbings []*proto.HTTPAuth) []*HTTPAuth {
	var items []*HTTPAuth
	for _, plumbing := range plumbings {
		items = append(items, httpAuthToPorcelain(plumbing))
	}
	return items
}

func cassandraToPorcelain(plumbing *proto.Cassandra) *Cassandra {
	if plumbing == nil {
		return nil
	}
	porcelain := &Cassandra{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	porcelain.TlsRequired = plumbing.TlsRequired
	return porcelain
}

func cassandraToPlumbing(porcelain *Cassandra) *proto.Cassandra {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Cassandra{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Port = porcelain.Port
	plumbing.TlsRequired = porcelain.TlsRequired
	return plumbing
}

func repeatedCassandraToPlumbing(porcelains []*Cassandra) []*proto.Cassandra {
	var items []*proto.Cassandra
	for _, porcelain := range porcelains {
		items = append(items, cassandraToPlumbing(porcelain))
	}
	return items
}

func repeatedCassandraToPorcelain(plumbings []*proto.Cassandra) []*Cassandra {
	var items []*Cassandra
	for _, plumbing := range plumbings {
		items = append(items, cassandraToPorcelain(plumbing))
	}
	return items
}

func mysqlToPorcelain(plumbing *proto.Mysql) *Mysql {
	if plumbing == nil {
		return nil
	}
	porcelain := &Mysql{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	return porcelain
}

func mysqlToPlumbing(porcelain *Mysql) *proto.Mysql {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Mysql{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.PortOverride = porcelain.PortOverride
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
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	return porcelain
}

func auroraMysqlToPlumbing(porcelain *AuroraMysql) *proto.AuroraMysql {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.AuroraMysql{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.PortOverride = porcelain.PortOverride
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
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	return porcelain
}

func clustrixToPlumbing(porcelain *Clustrix) *proto.Clustrix {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Clustrix{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.PortOverride = porcelain.PortOverride
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
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	return porcelain
}

func mariaToPlumbing(porcelain *Maria) *proto.Maria {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Maria{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.PortOverride = porcelain.PortOverride
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
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	return porcelain
}

func memsqlToPlumbing(porcelain *Memsql) *proto.Memsql {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Memsql{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.PortOverride = porcelain.PortOverride
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

func druidToPorcelain(plumbing *proto.Druid) *Druid {
	if plumbing == nil {
		return nil
	}
	porcelain := &Druid{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Port = plumbing.Port
	return porcelain
}

func druidToPlumbing(porcelain *Druid) *proto.Druid {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Druid{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Port = porcelain.Port
	return plumbing
}

func repeatedDruidToPlumbing(porcelains []*Druid) []*proto.Druid {
	var items []*proto.Druid
	for _, porcelain := range porcelains {
		items = append(items, druidToPlumbing(porcelain))
	}
	return items
}

func repeatedDruidToPorcelain(plumbings []*proto.Druid) []*Druid {
	var items []*Druid
	for _, plumbing := range plumbings {
		items = append(items, druidToPorcelain(plumbing))
	}
	return items
}

func sqlServerToPorcelain(plumbing *proto.SQLServer) *SQLServer {
	if plumbing == nil {
		return nil
	}
	porcelain := &SQLServer{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Database = plumbing.Database
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Port = plumbing.Port
	porcelain.OverrideDatabase = plumbing.OverrideDatabase
	return porcelain
}

func sqlServerToPlumbing(porcelain *SQLServer) *proto.SQLServer {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.SQLServer{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Database = porcelain.Database
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Port = porcelain.Port
	plumbing.OverrideDatabase = porcelain.OverrideDatabase
	return plumbing
}

func repeatedSQLServerToPlumbing(porcelains []*SQLServer) []*proto.SQLServer {
	var items []*proto.SQLServer
	for _, porcelain := range porcelains {
		items = append(items, sqlServerToPlumbing(porcelain))
	}
	return items
}

func repeatedSQLServerToPorcelain(plumbings []*proto.SQLServer) []*SQLServer {
	var items []*SQLServer
	for _, plumbing := range plumbings {
		items = append(items, sqlServerToPorcelain(plumbing))
	}
	return items
}

func mongoHybridToPorcelain(plumbing *proto.MongoHybrid) *MongoHybrid {
	if plumbing == nil {
		return nil
	}
	porcelain := &MongoHybrid{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.AuthDatabase = plumbing.AuthDatabase
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Port = plumbing.Port
	porcelain.ReplicaSet = plumbing.ReplicaSet
	porcelain.ConnectToReplica = plumbing.ConnectToReplica
	return porcelain
}

func mongoHybridToPlumbing(porcelain *MongoHybrid) *proto.MongoHybrid {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.MongoHybrid{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.AuthDatabase = porcelain.AuthDatabase
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Port = porcelain.Port
	plumbing.ReplicaSet = porcelain.ReplicaSet
	plumbing.ConnectToReplica = porcelain.ConnectToReplica
	return plumbing
}

func repeatedMongoHybridToPlumbing(porcelains []*MongoHybrid) []*proto.MongoHybrid {
	var items []*proto.MongoHybrid
	for _, porcelain := range porcelains {
		items = append(items, mongoHybridToPlumbing(porcelain))
	}
	return items
}

func repeatedMongoHybridToPorcelain(plumbings []*proto.MongoHybrid) []*MongoHybrid {
	var items []*MongoHybrid
	for _, plumbing := range plumbings {
		items = append(items, mongoHybridToPorcelain(plumbing))
	}
	return items
}

func mongoHostToPorcelain(plumbing *proto.MongoHost) *MongoHost {
	if plumbing == nil {
		return nil
	}
	porcelain := &MongoHost{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.AuthDatabase = plumbing.AuthDatabase
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Port = plumbing.Port
	return porcelain
}

func mongoHostToPlumbing(porcelain *MongoHost) *proto.MongoHost {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.MongoHost{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.AuthDatabase = porcelain.AuthDatabase
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Port = porcelain.Port
	return plumbing
}

func repeatedMongoHostToPlumbing(porcelains []*MongoHost) []*proto.MongoHost {
	var items []*proto.MongoHost
	for _, porcelain := range porcelains {
		items = append(items, mongoHostToPlumbing(porcelain))
	}
	return items
}

func repeatedMongoHostToPorcelain(plumbings []*proto.MongoHost) []*MongoHost {
	var items []*MongoHost
	for _, plumbing := range plumbings {
		items = append(items, mongoHostToPorcelain(plumbing))
	}
	return items
}

func mongoReplicaSetToPorcelain(plumbing *proto.MongoReplicaSet) *MongoReplicaSet {
	if plumbing == nil {
		return nil
	}
	porcelain := &MongoReplicaSet{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.Hostname = plumbing.Hostname
	porcelain.AuthDatabase = plumbing.AuthDatabase
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Username = plumbing.Username
	porcelain.Password = plumbing.Password
	porcelain.Port = plumbing.Port
	porcelain.ReplicaSet = plumbing.ReplicaSet
	porcelain.ConnectToReplica = plumbing.ConnectToReplica
	return porcelain
}

func mongoReplicaSetToPlumbing(porcelain *MongoReplicaSet) *proto.MongoReplicaSet {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.MongoReplicaSet{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.Hostname = porcelain.Hostname
	plumbing.AuthDatabase = porcelain.AuthDatabase
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Username = porcelain.Username
	plumbing.Password = porcelain.Password
	plumbing.Port = porcelain.Port
	plumbing.ReplicaSet = porcelain.ReplicaSet
	plumbing.ConnectToReplica = porcelain.ConnectToReplica
	return plumbing
}

func repeatedMongoReplicaSetToPlumbing(porcelains []*MongoReplicaSet) []*proto.MongoReplicaSet {
	var items []*proto.MongoReplicaSet
	for _, porcelain := range porcelains {
		items = append(items, mongoReplicaSetToPlumbing(porcelain))
	}
	return items
}

func repeatedMongoReplicaSetToPorcelain(plumbings []*proto.MongoReplicaSet) []*MongoReplicaSet {
	var items []*MongoReplicaSet
	for _, plumbing := range plumbings {
		items = append(items, mongoReplicaSetToPorcelain(plumbing))
	}
	return items
}

func athenaToPorcelain(plumbing *proto.Athena) *Athena {
	if plumbing == nil {
		return nil
	}
	porcelain := &Athena{}
	porcelain.ID = plumbing.Id
	porcelain.Name = plumbing.Name
	porcelain.Healthy = plumbing.Healthy
	porcelain.AccessKey = plumbing.AccessKey
	porcelain.SecretAccessKey = plumbing.SecretAccessKey
	porcelain.Output = plumbing.Output
	porcelain.PortOverride = plumbing.PortOverride
	porcelain.Region = plumbing.Region
	return porcelain
}

func athenaToPlumbing(porcelain *Athena) *proto.Athena {
	if porcelain == nil {
		return nil
	}
	plumbing := &proto.Athena{}
	plumbing.Id = porcelain.ID
	plumbing.Name = porcelain.Name
	plumbing.Healthy = porcelain.Healthy
	plumbing.AccessKey = porcelain.AccessKey
	plumbing.SecretAccessKey = porcelain.SecretAccessKey
	plumbing.Output = porcelain.Output
	plumbing.PortOverride = porcelain.PortOverride
	plumbing.Region = porcelain.Region
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

type resourceIteratorImplFetchFunc func() ([]Resource, bool, error)
type resourceIteratorImpl struct {
	buffer      []Resource
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

func (r *resourceIteratorImpl) Value() Resource {
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
