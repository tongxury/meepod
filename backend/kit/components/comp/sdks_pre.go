package comp

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/assist/loc"
	"gitee.com/meepo/backend/kit/components/chain/coingecko"
	"gitee.com/meepo/backend/kit/components/chain/eth"
	"gitee.com/meepo/backend/kit/components/chain/eth/swap"
	"gitee.com/meepo/backend/kit/components/chain/tron/tronclient"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/third/alipay"
	"gitee.com/meepo/backend/kit/components/third/aliyun"
	"gitee.com/meepo/backend/kit/components/third/ucloud"
	"gitee.com/meepo/backend/kit/components/third/xinsh"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/allegro/bigcache/v3"
	"github.com/boltdb/bolt"
	"github.com/go-pg/pg/v10"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/oschwald/geoip2-golang"
	"github.com/redis/go-redis/v9"
	"github.com/seefan/gossdb/v2"
	"github.com/seefan/gossdb/v2/conf"
	"github.com/segmentio/kafka-go"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/wzshiming/ssdb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

func (c *Pre) SSDB2(addr string) *Pre {

	if addr == "" {
		addr = "localhost:8888"
	}

	connect, err := ssdb.Connect(
		ssdb.Addr(addr),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = connect.SetX("__health_check", ssdb.Value("1"), 1*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	c.sdk.ssdb2Client = connect
	return c
}

func (c *Pre) SSDB(config conf.Config) *Pre {

	if config.Host == "" {
		config.Host = "localhost"
	}

	if config.Port == 0 {
		config.Port = 8888
	}

	//config.WriteBufferSize = 10240
	//config.PoolSize = 100
	//config.MinPoolSize = 100
	//config.MinPoolSize = 100

	err := gossdb.Start(&config)
	if err != nil {
		log.Fatal(err)
	}
	//defer gossdb.Shutdown()
	cli, err := gossdb.NewClient()
	if err != nil {
		log.Fatal("gossdb ", err)
	}
	if !cli.Ping() {
		log.Fatal("gossdb start err")
	}

	c.sdk.ssdbClient = cli

	return c
}

func (s *Sdk) Preparing() *Pre {
	return preIns
}

func (c *Pre) LevelDB(filepath string) *Pre {
	if filepath == "" {
		filepath = "__level.db"
	}

	db, err := leveldb.OpenFile(filepath, nil)
	if err != nil {
		log.Fatal(fmt.Errorf("level db: %v", err))
	}

	c.sdk.levelDBClient = db

	return c
}

func (c *Pre) FileDB(name string) *Pre {
	if name == "" {
		name = "_file.db"
	}
	db, err := bolt.Open(name, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(fmt.Errorf("file db: %v", err))
	}

	db.Path()
	c.sdk.fileDBClient = db

	return c
}

func (c *Pre) Alipay(options alipay.Options) *Pre {
	c.sdk.alipayClient = alipay.NewAlipayClient(options)
	return c
}

func (c *Pre) AliOSS(options AliOSSOptions) *Pre {
	cli, err := aliyun.NewOSSClient(options.Entrypoint, options.AccessKey, options.AccessSecret)
	if err != nil {
		log.Fatal(err)
	}
	c.sdk.aliOSSClient = cli
	return c
}
func (c *Pre) AliSMS(options AliSMSOptions) *Pre {
	cli, err := aliyun.NewSMSClient(options.AccessKey, options.AccessSecret)
	if err != nil {
		log.Fatal(err)
	}
	c.sdk.aliSMSClient = cli
	return c
}

func (c *Pre) Xinsh(conf xinsh.Config) *Pre {

	c.sdk.xinshClient = xinsh.NewXinShPayClient(conf)
	return c
}
func (c *Pre) UCloudSMS(options UCloudSMSOptions) *Pre {
	cli, err := ucloud.NewSMSClient(options.AccessKey, options.AccessSecret)
	if err != nil {
		log.Fatal(err)
	}
	c.sdk.ucloudSMSClient = cli
	return c
}

func (c *Pre) OSS(options OssOptions) *Pre {

	entrypoint := options.Entrypoint
	useSSL := false
	if strings.HasPrefix(options.Entrypoint, "https://") {
		entrypoint = strings.ReplaceAll(options.Entrypoint, "https://", "")
		useSSL = true
	} else if strings.HasPrefix(options.Entrypoint, "http://") {
		entrypoint = strings.ReplaceAll(options.Entrypoint, "http://", "")
		useSSL = false
	}

	minioClient, err := minio.New(entrypoint, &minio.Options{
		Creds:  credentials.NewStaticV4(options.AccessKey, options.AccessSecret, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatal(err)
	}

	if minioClient.IsOffline() {
		log.Fatal(fmt.Errorf("oss IsOffline"))
	}

	c.sdk.ossClient = minioClient
	return c
}

func (c *Pre) Geoip(dbFile string) *Pre {
	db, err := geoip2.Open(dbFile)
	if err != nil {
		log.Fatal(err)
	}

	c.sdk.geoipClient = db
	return c
}

func (c *Pre) Logger(options slf.Config) *Pre {
	if err := slf.InitLogger(options); err != nil {
		log.Fatal(err)
	}
	return c
}

func (c *Pre) Redis(options redis.Options) *Pre {
	cli := redis.NewClient(&options)

	_, err := cli.Set(context.Background(), "__health", "1", 1*time.Millisecond).Result()
	if err != nil {
		log.Fatal(fmt.Errorf("redis: %v", err))
	}

	c.sdk.redisClient = cli
	return c
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	sql, err := q.FormattedQuery()
	slf.WithError(err).Debugw("", slf.String("sql", strings.ReplaceAll(string(sql), "\"", "")))
	return nil
}

func (c *Pre) Loc() *Pre {
	c.sdk.locClient = loc.NewLocClient()

	return c
}

func (c *Pre) Postgres(options PGOptions) *Pre {

	cli := pg.Connect(&options.Options)

	if options.PrintSQL {
		cli.AddQueryHook(dbLogger{})
	}

	_, err := cli.Exec("select 1")
	if err != nil {
		log.Fatal("postgres", err)
	}

	c.sdk.pgClient = cli

	return c
}

func (c *Pre) KafkaWriter(addresses string) *Pre {

	addrs := strings.Split(addresses, ",")

	c.sdk.kafkaWriter = &kafka.Writer{
		Addr: kafka.TCP(addrs...),
		// NOTE: When Topic is not defined here, each Message must define it instead.
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: false,
	}

	return c
}

func (c *Pre) LocalCache() *Pre {

	//  todo
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(100*time.Minute))
	if err != nil {
		log.Fatal("local cache", err)
	}

	c.sdk.cacheClient = cache
	return c
}

func (c *Pre) Mysql(dsn string) *Pre {

	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	if dsn == "" {
		log.Fatal("mysql dsn is required")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		log.Fatal("mysql", err)
	}

	err = db.Select("select 1").Error
	if err != nil {
		log.Fatal("mysql", err)
	}

	c.sdk.mysqlClient = db

	return c
}

func (c *Pre) Eth(options EthOptions) *Pre {

	urls := strings.Split(options.Rpc, ",")

	clients, err := eth.NewEthClients(urls)
	if err != nil {
		log.Fatal("eth clients", err)
	}

	// todo

	c.sdk.ethClients = clients

	c1 := clients[:1]
	c2 := clients[1:]
	c.sdk.ethClientV2, _ = eth.NewEthClient(c1)
	c.sdk.ethClientV3, _ = eth.NewEthClient(c2)

	c.sdk.ethGraphClient = eth.NewGraphClient(urls[0])
	c.sdk.etherScanClient = eth.NewEtherScanClient("")
	c.sdk.etherScanCrawlerClient = eth.NewEtherScanCrawlerClient()
	c.sdk.uniApolloClient = swap.NewUniApolloClient()
	c.sdk.sushiApolloClient = swap.NewSushiApolloClient()
	c.sdk.coingeckoClient = coingecko.NewClient()

	return c
}

func (c *Pre) ClickHouse(options ClickHouseOptions) *Pre {

	addrs := strings.Split(options.Addr, ",")

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr:             addrs,
		ConnOpenStrategy: clickhouse.ConnOpenRoundRobin,
		Auth: clickhouse.Auth{
			Database: options.Database,
			Username: options.Username,
			Password: options.Password,
		},
	})
	if err != nil {
		log.Fatal("clickhouse", err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("clickhouse", err)
	}

	c.sdk.clickhouseClient = &conn

	return c
}

func (c *Pre) Tron(options TronOptions) *Pre {

	apiKey := "dd5d773e-4392-4bcb-95f1-811542e16dd7"

	client, err := tronclient.NewTronClient(options.Rpc, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.GrpcClient().GetNowBlock()
	if err != nil {
		log.Fatal(err)
	}

	c.sdk.tronClient = client

	gridClient, err := tronclient.NewTronGridClient(options.Rpc, apiKey)
	if err != nil {
		return nil
	}

	_, err = client.GrpcClient().GetNowBlock()
	if err != nil {
		log.Fatal(err)
	}

	c.sdk.tronGridClient = gridClient

	return c
}
