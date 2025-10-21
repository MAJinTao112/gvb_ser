package core

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"log"
	"time"
)

func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)

}
func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       db,
		PoolSize: redisConf.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Errorf("redis connect err:%v", redisConf.Addr)
		return nil
		//return nil
	}
	log.Println("redis connect success")
	return rdb
}
