package cache

import (
	"task-system/log"

	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
)

var redisdb *redis.Client

// 初始化连接
func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"pong": pong,
			"err":  err,
		}).Fatalln()
	}
	log.Logger.WithFields(logrus.Fields{
		"pong": pong,
		"err":  err,
	})
}
