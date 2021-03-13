package config

import (
	"gopkg.in/ini.v1"
	"log"
)

const (
	// DefaultBindAddress 监听地址
	DefaultBindAddress = "0.0.0.0:9277"
	// DefaultBucketSize bucket数量
	DefaultBucketSize = 3
	// DefaultBucketName bucket名称
	DefaultBucketName = "dq_bucket_%d"
	// DefaultQueueName 队列名称
	DefaultQueueName = "dq_queue_%s"
	// DefaultQueueBlockTimeout 轮询队列超时时间
	DefaultQueueBlockTimeout = 178
	// DefaultRedisHost Redis连接地址
	DefaultRedisHost = "127.0.0.1:6379"
	// DefaultRedisDb Redis数据库编号
	DefaultRedisDb = 1
	// DefaultRedisPassword Redis密码
	DefaultRedisPassword = ""
	// DefaultRedisMaxIdle Redis连接池闲置连接数
	DefaultRedisMaxIdle = 10
	// DefaultRedisMaxActive Redis连接池最大激活连接数, 0为不限制
	DefaultRedisMaxActive = 0
	// DefaultRedisConnectTimeout Redis连接超时时间,单位毫秒
	DefaultRedisConnectTimeout = 5000
	// DefaultRedisReadTimeout Redis读取超时时间, 单位毫秒
	DefaultRedisReadTimeout = 180000
	// DefaultRedisWriteTimeout Redis写入超时时间, 单位毫秒
	DefaultRedisWriteTimeout = 3000
)

type Config struct {
	BindAddress       string      // http server 监听地址
	BucketSize        int         // bucket数量
	BucketName        string      // bucket在redis中的键名,
	QueueName         string      // ready queue在redis中的键名
	QueueBlockTimeout int         // 调用blpop阻塞超时时间, 单位秒, 修改此项, redis.read_timeout必须做相应调整
	Redis             RedisConfig // redis配置
}

type RedisConfig struct {
	Host           string
	Db             int
	Password       string
	MaxIdle        int // 连接池最大空闲连接数
	MaxActive      int // 连接池最大激活连接数
	ConnectTimeout int // 连接超时, 单位毫秒
	ReadTimeout    int // 读取超时, 单位毫秒
	WriteTimeout   int // 写入超时, 单位毫秒
}

var (
	Setting *Config
)

func Init(path string) {
	Setting = &Config{}

	if path == "" {
		Setting.initDefaultConfig()
		return
	}
	Setting.parse(path)
}

func (config *Config) parse(path string) {

	file, err := ini.Load(path)
	if err != nil {
		log.Fatalln("无法解析配置文件#%s", err.Error())
	}
	section := file.Section("") // 默认分区可以使用空字符串表示
	config.BindAddress = section.Key("bind_address").MustString(DefaultBindAddress)
	config.BucketSize = section.Key("bucket_size").MustInt(DefaultBucketSize)
	config.BucketName = section.Key("bucket_name").MustString(DefaultBucketName)
	config.QueueName = section.Key("queue_name").MustString(DefaultQueueName)
	config.QueueBlockTimeout = section.Key("queue_block_timeout").MustInt(DefaultQueueBlockTimeout)
	config.Redis.Host = section.Key("redis.host").MustString(DefaultRedisHost)
	config.Redis.Db = section.Key("redis.db").MustInt(DefaultRedisDb)
	config.Redis.Password = section.Key("redis.password").MustString(DefaultRedisPassword)
	config.Redis.MaxIdle = section.Key("redis.max_idle").MustInt(DefaultRedisMaxIdle)
	config.Redis.MaxActive = section.Key("redis.max_active").MustInt(DefaultRedisMaxActive)
	config.Redis.ConnectTimeout = section.Key("redis.connect_timeout").MustInt(DefaultRedisConnectTimeout)
	config.Redis.ReadTimeout = section.Key("redis.read_timeout").MustInt(DefaultRedisReadTimeout)
	config.Redis.WriteTimeout = section.Key("redis.write_timeout").MustInt(DefaultRedisWriteTimeout)
}

func (config *Config) initDefaultConfig() {
	config.BindAddress = DefaultBindAddress
	config.BucketSize = DefaultBucketSize
	config.BucketName = DefaultBucketName
	config.QueueName = DefaultQueueName
	config.QueueBlockTimeout = DefaultQueueBlockTimeout
	config.Redis.Host = DefaultRedisHost
	config.Redis.Db = DefaultRedisDb
	config.Redis.Password = DefaultRedisPassword
	config.Redis.MaxIdle = DefaultRedisMaxIdle
	config.Redis.MaxActive = DefaultRedisMaxActive
	config.Redis.ConnectTimeout = DefaultRedisConnectTimeout
	config.Redis.ReadTimeout = DefaultRedisReadTimeout
	config.Redis.WriteTimeout = DefaultRedisWriteTimeout
}
