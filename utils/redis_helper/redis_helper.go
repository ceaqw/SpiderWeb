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
	Redis *redis.Client
)

func init() {
	Redis = NewRedisHelper()
}

func NewRedisHelper() *redis.Client {
	addr := conf.GetRedisCfg()
	client := redis.NewClient(&redis.Options{
		Addr:     addr.Addr,
		Password: addr.Password,
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("redis connect error: %#v\n", err.Error()))
	}
	return client
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
