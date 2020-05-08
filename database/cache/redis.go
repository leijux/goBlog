package cache

import (
	"errors"
	"time"

	"task-system/config"
	"task-system/log"
	"task-system/models"

	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
)

//Redisdb Redis缓存库链接对象
var redisdb *redis.Client

var ErrRedisOff = errors.New("redis off !")

// 初始化连接
func init() {
	if !config.Cfg.Database.Redis.IsOpen {
		return
	}
	redisdb = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Database.Redis.Addr,
		Password: config.Cfg.Database.Redis.Password, // no password set
		DB:       config.Cfg.Database.Redis.Db,       // use default DB
	})

	pong, err := redisdb.Ping().Result()
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

func Get(key string) (value string, err error) {
	if redisdb == nil {
		err = ErrRedisOff
		return
	}
	value, err = redisdb.Get(key).Result()
	return
}

func Set(key string, value models.IModels, t time.Duration) (err error) {
	if redisdb == nil {
		err = ErrRedisOff
		return
	}
	err = redisdb.Set(key, value.ToJSON(), t).Err()
	return
}
