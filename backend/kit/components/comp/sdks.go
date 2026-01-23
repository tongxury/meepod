package comp

import (
	"gitee.com/meepo/backend/kit/components/assist/ip"
	"gitee.com/meepo/backend/kit/components/assist/loc"
	"gitee.com/meepo/backend/kit/components/chain/coingecko"
	"gitee.com/meepo/backend/kit/components/chain/eth"
	"gitee.com/meepo/backend/kit/components/chain/eth/swap"
	"gitee.com/meepo/backend/kit/components/chain/tron/tronclient"
	"gitee.com/meepo/backend/kit/components/third/alipay"
	"gitee.com/meepo/backend/kit/components/third/aliyun"
	"gitee.com/meepo/backend/kit/components/third/ucloud"
	"gitee.com/meepo/backend/kit/components/third/xinsh"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/allegro/bigcache/v3"
	"github.com/boltdb/bolt"
	"github.com/go-pg/pg/v10"
	"github.com/minio/minio-go/v7"
	"github.com/oschwald/geoip2-golang"
	"github.com/redis/go-redis/v9"
	"github.com/seefan/gossdb/v2/pool"
	"github.com/segmentio/kafka-go"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/wzshiming/ssdb"
	"gorm.io/gorm"
)

func SDK() *Sdk {
	return preIns.sdk
}

var preIns = &Pre{
	sdk: &Sdk{},
}

type Pre struct {
	sdk *Sdk
}

type Sdk struct {
	locClient *loc.Client

	alipayClient    *alipay.Client
	aliOSSClient    *aliyun.OSSClient
	aliSMSClient    *aliyun.SMSClient
	xinshClient     *xinsh.Client
	ucloudSMSClient *ucloud.SMSClient

	ossClient        *minio.Client
	geoipClient      *geoip2.Reader
	ipClient         *ip.Client
	redisClient      *redis.Client
	fileDBClient     *bolt.DB
	ssdbClient       *pool.Client
	ssdb2Client      *ssdb.Client
	levelDBClient    *leveldb.DB
	cacheClient      *bigcache.BigCache
	pgClient         *pg.DB
	mysqlClient      *gorm.DB
	clickhouseClient *driver.Conn

	ethClients  eth.Clients
	ethClientV2 *eth.Client
	ethClientV3 *eth.Client

	ethGraphClient         *eth.GraphClient
	etherScanClient        *eth.EtherScanClient
	etherScanCrawlerClient *eth.EtherScanCrawlerClient
	coingeckoClient        *coingecko.Client

	uniApolloClient   *swap.UniApolloClient
	sushiApolloClient *swap.SushiApolloClient

	tronClient     *tronclient.TronClient
	tronGridClient *tronclient.TronGridClient

	kafkaWriter *kafka.Writer
}
