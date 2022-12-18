package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
	"strconv"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func init() {
	LoadRedisConfig()
	//fmt.Println("redis 连接密码：", RedisPw)
	RedisConnection()
}

func LoadRedisConfig() {
	// 不在conf包下统一引入使用，是因为存在循环依赖
	//RedisAddr = conf.RedisAddr
	//RedisDb = conf.RedisDb
	//RedisPw = conf.RedisPw
	//RedisDbName = conf.RedisDbName

	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		// 处理读取配置文件异常
		// panic 直译为 运行时恐慌 当panic被抛出异常后，如果我们没有在程序中添加任何保护措施的话，程序就会打印出panic的详细情况之后，终止运行
		fmt.Println("redis 配置文件读取错误，请检查文件路径:", err)
	}

	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func RedisConnection() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		// redis 连接设置
		Addr:     RedisAddr,
		Password: RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	RedisClient = client
}
