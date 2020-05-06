package cache

import (
	"task-system/config"
	"task-system/log"

	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
)
//Redisdb Redis缓存库链接对象
var Redisdb *redis.Client

// 初始化连接
func init() {
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Database.Redis.Addr,
		Password: config.Cfg.Database.Redis.Password, // no password set
		DB:       config.Cfg.Database.Redis.Db,  // use default DB
	})

	pong, err := Redisdb.Ping().Result()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"pong": pong,
			"err":  err,
		}).Fatalln()
	}
	log.Logger.WithFields(logrus.Fields{
		"pong": pong,
		"err":  err,
	}).Infoln()
}
