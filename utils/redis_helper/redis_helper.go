package redis_helper

import (
	"SpiderWeb/conf"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

type RedisHelper struct {
}

var (
	Redis        *redis.Client
	RedisCluster *redis.ClusterClient
)

func init() {
	redisCfg := conf.GetRedisCfg()
	if redisCfg.ClusterMode == true {
		RedisCluster = NewRedisClusterHelper(redisCfg)
	} else {
		Redis = NewRedisHelper(redisCfg)
	}
}

func NewRedisHelper(redisCfg conf.RedisCfg) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         redisCfg.Addrs[0],
		Password:     redisCfg.Password,
		DB:           0,
		DialTimeout:  redisCfg.DialTimeout,
		ReadTimeout:  redisCfg.ReadTimeout,
		WriteTimeout: redisCfg.WriteTimeout,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("redis connect error: %#v\n", err.Error()))
	}
	return client
}

func NewRedisClusterHelper(redisCfg conf.RedisCfg) *redis.ClusterClient {
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        redisCfg.Addrs,
		Password:     redisCfg.Password,
		DialTimeout:  redisCfg.DialTimeout,
		ReadTimeout:  redisCfg.ReadTimeout,
		WriteTimeout: redisCfg.WriteTimeout,
	})
	if _, err := clusterClient.Ping().Result(); err != nil {
		panic(fmt.Sprintf("redis connect error: %#v\n", err.Error()))
	}
	return clusterClient
}

func (h RedisHelper) SaveRedisToken(user string, token string) error {
	expireTime := viper.GetDuration("jwt.time_out")
	_, err := Redis.Set(user, token, expireTime*time.Minute).Result()
	return err
}

func (h RedisHelper) RemoveRedisToken(user string) error {
	_, err := Redis.Del(user).Result()
	return err
}
