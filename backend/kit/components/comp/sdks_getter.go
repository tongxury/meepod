package comp

import (
	"gitee.com/meepo/backend/kit/components/assist/loc"
	"gitee.com/meepo/backend/kit/components/chain/coingecko"
	"gitee.com/meepo/backend/kit/components/chain/cryptocom"
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
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-pg/pg/v10"
	"github.com/minio/minio-go/v7"
	"github.com/oschwald/geoip2-golang"
	"github.com/redis/go-redis/v9"
	"github.com/seefan/gossdb/v2"
	"github.com/seefan/gossdb/v2/pool"
	"github.com/segmentio/kafka-go"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/wzshiming/ssdb"
	"gorm.io/gorm"
)

func (s *Sdk) Loc() *loc.Client {
	if s.locClient == nil {
		panic("loc client is not initialized")
	}

	return s.locClient
}

func (s *Sdk) Alipay() *alipay.Client {
	if s.alipayClient == nil {
		panic("alipay client is not initialized")
	}

	return s.alipayClient
}

func (s *Sdk) Xinsh() *xinsh.Client {
	if s.xinshClient == nil {
		panic("xinsh client is not initialized")
	}

	return s.xinshClient
}

func (s *Sdk) Coingecko() *coingecko.Client {
	return s.coingeckoClient
}

func (s *Sdk) SSDB2() *ssdb.Client {
	if s.ssdb2Client == nil {
		panic("ssdb db is not initialized")
	}

	return s.ssdb2Client
}

func (s *Sdk) SSDB() *pool.Client {
	//if s.ssdbClient == nil {
	//	panic("ssdb db is not initialized")
	//}

	return gossdb.Client()
}

func (s *Sdk) LevelDB() *leveldb.DB {
	if s.levelDBClient == nil {
		panic("level db is not initialized")
	}

	return s.levelDBClient
}

func (s *Sdk) FileDB() *bolt.DB {
	if s.fileDBClient == nil {
		panic("file db is not initialized")
	}

	return s.fileDBClient
}

func (s *Sdk) OSS() *minio.Client {
	if s.ossClient == nil {
		panic("oss client is not initialized")
	}

	return s.ossClient
}
func (s *Sdk) Geoip() *geoip2.Reader {
	if s.geoipClient == nil {
		panic("geoip client is not initialized")
	}

	return s.geoipClient
}

func (s *Sdk) Redis() *redis.Client {
	if s.redisClient == nil {
		panic("redis client is not initialized")
	}

	return s.redisClient
}
func (s *Sdk) LocalCache() *bigcache.BigCache {
	if s.cacheClient == nil {
		panic("local cache client is not initialized")
	}

	return s.cacheClient
}

func (s *Sdk) Postgres() *pg.DB {
	if s.pgClient == nil {
		panic("pg client is not initialized")
	}

	return s.pgClient
}
func (s *Sdk) Mysql() *gorm.DB {
	if s.mysqlClient == nil {
		panic("mysql client is not initialized")
	}

	return s.mysqlClient
}

func (s *Sdk) ClickHouse() driver.Conn {
	if s.clickhouseClient == nil {
		panic("click house client is not initialized")
	}

	return *s.clickhouseClient
}
func (s *Sdk) KafkaWriter() *kafka.Writer {
	if s.kafkaWriter == nil {
		panic("kafka writer  is not initialized")
	}

	return s.kafkaWriter
}

func (s *Sdk) EthClient() *ethclient.Client {
	if len(s.ethClients) == 0 {
		panic("eth client is not initialized")
	}

	return s.ethClients.Get()
}

func (s *Sdk) AliOSS() *aliyun.OSSClient {
	if s.aliOSSClient == nil {
		panic("ali oss is not initialized")
	}

	return s.aliOSSClient
}

func (s *Sdk) AliSMS() *aliyun.SMSClient {
	if s.aliSMSClient == nil {
		panic("ali sms is not initialized")
	}

	return s.aliSMSClient
}

func (s *Sdk) UCloudSMS() *ucloud.SMSClient {
	if s.ucloudSMSClient == nil {
		panic("ucloud sms is not initialized")
	}

	return s.ucloudSMSClient
}

func (s *Sdk) EthClientV2() *eth.Client {
	if s.ethClientV2 == nil {
		panic("eth client is not initialized")
	}

	return s.ethClientV2
}

func (s *Sdk) EthClientV3() *eth.Client {
	if s.ethClientV3 == nil {
		panic("eth client is not initialized")
	}

	return s.ethClientV3
}

func (s *Sdk) EthGraphClient() *eth.GraphClient {
	if s.ethGraphClient == nil {
		panic("eth graph client is not initialized")
	}

	return s.ethGraphClient
}

func (s *Sdk) EtherScanClient() *eth.EtherScanClient {
	if s.etherScanClient == nil {
		panic("eth scan client is not initialized")
	}

	return s.etherScanClient
}

func (s *Sdk) EtherScanCrawlerClient() *eth.EtherScanCrawlerClient {
	if s.etherScanCrawlerClient == nil {
		panic("eth scan client is not initialized")
	}

	return s.etherScanCrawlerClient
}

func (s *Sdk) UniApolloClient() *swap.UniApolloClient {
	if s.uniApolloClient == nil {
		panic("uniswap client is not initialized")
	}

	return s.uniApolloClient
}

func (s *Sdk) SushiApolloClient() *swap.SushiApolloClient {
	if s.sushiApolloClient == nil {
		panic("sushiswap client is not initialized")
	}

	return s.sushiApolloClient
}

func (s *Sdk) TronClient() *tronclient.TronClient {
	if s.tronClient == nil {
		panic("tron client is not initialized")
	}

	return s.tronClient
}

func (s *Sdk) TronGridClient() *tronclient.TronGridClient {
	if s.tronGridClient == nil {
		panic("tron grid client is not initialized")
	}

	return s.tronGridClient
}

func (s *Sdk) CryptoComClient() *cryptocom.Client {
	return cryptocom.NewCryptoComClient()
}
