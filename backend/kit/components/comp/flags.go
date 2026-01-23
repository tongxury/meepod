package comp

import (
	"flag"

	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/third/alipay"
	"gitee.com/meepo/backend/kit/components/third/xinsh"
	"github.com/go-pg/pg/v10"
	redisV9 "github.com/redis/go-redis/v9"
	"github.com/seefan/gossdb/v2/conf"
)

var vflags = &FlagValues{}

func Flags() *FlagValues {
	return vflags
}

type FlagValues struct {
	RepoRedis        redisV9.Options
	RepoPostgres     PGOptions
	RepoSSDB         conf.Config
	RepoSSDBAddr     string
	RepoMysqlDSN     string
	FileDBName       string
	LevelDBPath      string
	RepoClickHouse   ClickHouseOptions
	Log              slf.Config
	GeoipDB          string
	OSSOptions       OssOptions
	AliOSSOptions    AliOSSOptions
	AliSMSOptions    AliSMSOptions
	AlipayOptions    alipay.Options
	UCloudSMSOptions UCloudSMSOptions
	XinshConfig      xinsh.Config
	EthOptions       EthOptions
	TronOptions      TronOptions
	KafkaAddress     string
	ServerOptions    ServerOptions
	custom           map[string]*string
}

type ServerOptions struct {
	Name string
	Test bool
	Port string
}

type ClickHouseOptions struct {
	Addr     string
	Database string
	Username string
	Password string
}

type TronOptions struct {
	Rpc string
}
type EthOptions struct {
	Rpc string
}

type PGOptions struct {
	pg.Options
	PrintSQL bool
}

type OssOptions struct {
	Entrypoint   string
	AccessKey    string
	AccessSecret string
	UseSSL       bool
}

type AliOSSOptions struct {
	AccessKey    string
	AccessSecret string
	Entrypoint   string
	Bucket       string
}

type AliSMSOptions struct {
	AccessKey    string
	AccessSecret string
	Sign         string
	TemplateCode string
	TestPrefix   string
}

type UCloudSMSOptions struct {
	AccessKey    string
	AccessSecret string
	Sign         string
	TemplateCode string
	TestPrefix   string
}

func (c *FlagValues) Parse() *FlagValues {
	flag.Parse()
	return c
}

func (c *FlagValues) GetStr(name string) string {
	return *c.custom[name]
}

func (c *FlagValues) GetBool(name string) bool {
	return *c.custom[name] == "true"
}

func (c *FlagValues) CustomStr(name, defValue, desc string) *FlagValues {
	if c.custom == nil {
		c.custom = map[string]*string{}
	}
	c.custom[name] = flag.String(name, defValue, desc)
	return c
}

func (c *FlagValues) Server() *FlagValues {
	flag.StringVar(&c.ServerOptions.Name, "srv.name", "", "server name")
	flag.BoolVar(&c.ServerOptions.Test, "srv.test", false, "is test")
	flag.StringVar(&c.ServerOptions.Port, "srv.port", "8080", "server port")
	return c
}

func (c *FlagValues) OSS() *FlagValues {
	flag.StringVar(&c.OSSOptions.Entrypoint, "oss.entrypoint", "http://120.27.131.135:8081", "oss entrypoint")
	flag.StringVar(&c.OSSOptions.AccessKey, "oss.accessKey", "M6l71AElFHKllhCtiUR6", "oss access key")
	flag.StringVar(&c.OSSOptions.AccessSecret, "oss.accessSecret", "aAX5RGAeeNyejo6lhqyz8IdguL4VXOe5ULaUPEUf", "oss access secret")
	flag.BoolVar(&c.OSSOptions.UseSSL, "oss.useSsl", false, "oss use ssl")
	return c
}

func (c *FlagValues) AliOSS() *FlagValues {
	flag.StringVar(&c.AliOSSOptions.Entrypoint, "ali.oss.entrypoint", "oss-cn-beijing.aliyuncs.com", "ali oss entrypoint")
	flag.StringVar(&c.AliOSSOptions.AccessKey, "ali.oss.accessKey", "", "ali oss access key")
	flag.StringVar(&c.AliOSSOptions.AccessSecret, "ali.oss.accessSecret", "", "ali oss access secret")
	flag.StringVar(&c.AliOSSOptions.Bucket, "ali.oss.bucket", "eimg", "ali oss bucket")
	return c
}

func (c *FlagValues) AliSMS() *FlagValues {
	flag.StringVar(&c.AliSMSOptions.Sign, "ali.sms.sign", "", "ali sms sign")
	flag.StringVar(&c.AliSMSOptions.TemplateCode, "ali.sms.templateCode", "", "ali sms template code")
	flag.StringVar(&c.AliSMSOptions.AccessKey, "ali.sms.accessKey", "", "ali sms access key")
	flag.StringVar(&c.AliSMSOptions.AccessSecret, "ali.sms.accessSecret", "", "ali sms access secret")
	flag.StringVar(&c.AliSMSOptions.TestPrefix, "ali.sms.testPrefix", "93", "ali sms test prefix")
	return c
}

func (c *FlagValues) Alipay() *FlagValues {

	var privateKey string
	var alipayPublicKey string

	flag.StringVar(&c.AlipayOptions.Entrypoint, "ali.pay.entrypoint", "https://openapi.alipay.com/gateway.do", "ali pay entrypoint")
	flag.StringVar(&c.AlipayOptions.AppId, "ali.pay.appId", "2021004101679599", "ali pay appId")
	flag.StringVar(&c.AlipayOptions.PrivateKey, "ali.pay.privateKey", privateKey, "ali pay privateKey")
	flag.StringVar(&c.AlipayOptions.AlipayPublicKey, "ali.pay.accessSecret", alipayPublicKey, "ali pay alipayPublicKey")
	flag.StringVar(&c.AlipayOptions.Income.Name, "ali.pay.income.name", "福彩科技（北京）有限公司", "ali pay income name")
	flag.StringVar(&c.AlipayOptions.Income.Pid, "ali.pay.income.pid", "2088641357051477", "ali pay income pid")
	flag.StringVar(&c.AlipayOptions.Income.Memo, "ali.pay.income.memo", "福彩科技", "ali pay income memo")
	flag.StringVar(&c.AlipayOptions.Income.LoginName, "ali.pay.income.loginName", "lottery@example.com", "ali pay income loginName")
	flag.Float64Var(&c.AlipayOptions.Income.Rate, "ali.pay.income.rate", 0.005, "ali pay income rate")

	return c
}

func (c *FlagValues) Xinsh() *FlagValues {
	return c
}

func (c *FlagValues) UCloudSMS() *FlagValues {
	flag.StringVar(&c.UCloudSMSOptions.Sign, "ucloud.sms.sign", "福彩科技", "ucloud sms sign")
	flag.StringVar(&c.UCloudSMSOptions.TemplateCode, "ucloud.sms.templateCode", "UTA230808KQJPZH", "ucloud sms template code")
	flag.StringVar(&c.UCloudSMSOptions.AccessKey, "ucloud.sms.accessKey", "7PT9vS5XSDpFa0n7MqlIsmINWsP50cTF21CrRJzDm5", "ucloud sms access key")
	flag.StringVar(&c.UCloudSMSOptions.AccessSecret, "ucloud.sms.accessSecret", "BiFJKxnx8raF7U7WZtZAm87PoKpEjG7PhApHGa8nTnZ749PRYYnMwHHC0tAYJLw6u5", "ucloud sms access secret")
	flag.StringVar(&c.UCloudSMSOptions.TestPrefix, "ucloud.sms.testPrefix", "93", "ucloud sms test prefix")
	return c
}

func (c *FlagValues) Geoip() *FlagValues {
	flag.StringVar(&c.GeoipDB, "geoip.db", "./GeoIP2-City.mmdb", "geoip mmdb file path")
	return c
}

func (c *FlagValues) Logger() *FlagValues {

	flag.StringVar(&c.Log.Level, "log.level", "debug", "log level")
	flag.StringVar(&c.Log.Output, "log.output", "stdout", "log output")
	flag.StringVar(&c.Log.Formatter, "log.formatter", "json", "log formatter, eg. json|text|logfmt")
	flag.StringVar(&c.Log.File.FileName, "log.filename", "", "output will be 'stdout' if filename is not set, eg. /logs/%level-%Y-%m-%d.log")

	return c
}

func (c *FlagValues) Redis() *FlagValues {

	flag.StringVar(&c.RepoRedis.Addr, "repo.redis.addr", "127.0.0.1:6379", "redis addr(host:port)")
	flag.StringVar(&c.RepoRedis.Username, "repo.redis.username", "default", "redis username")
	flag.StringVar(&c.RepoRedis.Password, "repo.redis.password", "1qaz@WSX", "redis password")
	return c
}

func (c *FlagValues) Postgres() *FlagValues {

	flag.StringVar(&c.RepoPostgres.Addr, "repo.pg.address", "127.0.0.1:5432", "postgres addr(host:port)")
	flag.StringVar(&c.RepoPostgres.Password, "repo.pg.password", "1qaz@WSX", "postgres password")
	flag.StringVar(&c.RepoPostgres.Database, "repo.pg.database", "", "postgres database")
	flag.BoolVar(&c.RepoPostgres.PrintSQL, "repo.pg.printsql", false, "postgres print sql ")
	return c
}

func (c *FlagValues) Mysql() *FlagValues {
	flag.StringVar(&c.RepoMysqlDSN, "repo.mysql.dsn", "", "mysql dsn(user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local)")
	return c
}

func (c *FlagValues) FileDB() *FlagValues {
	flag.StringVar(&c.FileDBName, "repo.filedb.name", "__file.db", "")
	return c
}

func (c *FlagValues) LevelDB() *FlagValues {
	flag.StringVar(&c.LevelDBPath, "repo.leveldb.path", "__level.db", "")
	return c
}

func (c *FlagValues) SSDB() *FlagValues {
	flag.StringVar(&c.RepoSSDB.Host, "repo.ssdb.host", "localhost", "")
	flag.IntVar(&c.RepoSSDB.Port, "repo.ssdb.port", 8888, "")
	return c
}

func (c *FlagValues) SSDB2() *FlagValues {
	flag.StringVar(&c.RepoSSDBAddr, "repo.ssdb.addr", "localhost:8888", "")
	return c
}

func (c *FlagValues) LocalCache() *FlagValues {
	return c
}

func (c *FlagValues) Kafka() *FlagValues {
	flag.StringVar(&c.KafkaAddress, "kafka.address", "", "kafka brokers(eg. localhost:9092,localhost:9093,localhost:9094)")
	return c
}

func (c *FlagValues) ClickHouse() *FlagValues {

	flag.StringVar(&c.RepoClickHouse.Addr, "repo.ch.addr", "", "clickhouse addr(host1:port1,host2:port2)")
	flag.StringVar(&c.RepoClickHouse.Username, "repo.ch.username", "", "clickhouse username")
	flag.StringVar(&c.RepoClickHouse.Password, "repo.ch.password", "", "clickhouse password")
	flag.StringVar(&c.RepoClickHouse.Database, "repo.ch.database", "", "clickhouse database")
	return c
}

func (c *FlagValues) Eth() *FlagValues {

	//flag.StringVar(&c.EthOptions.Rpc, "eth.rpc.url", "https://bsc-dataseed1.binance.org", "eth rpc addr(http://node1,http://node2)")
	flag.StringVar(&c.EthOptions.Rpc, "eth.rpc.url", "http://195.201.8.156:8545", "eth rpc addr(http://node1,http://node2)")
	return c
}

func (c *FlagValues) Tron() *FlagValues {
	flag.StringVar(&c.TronOptions.Rpc, "tron.rpc.url", "", "tron rpc addr")
	return c
}
