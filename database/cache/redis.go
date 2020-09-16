package cache

import (
	"time"

	"goBlog/config"
	"goBlog/log"
	"goBlog/models"

	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

//Redisdb Redis缓存库链接对象
var redisdb *redis.Client

//ErrRedisOff 数据库关闭
var ErrRedisOff = errors.New("redis off")

// 初始化连接
func init() {
	if open := config.GetBool("database.redis.isOpen"); !open {
		return
	}
	addr := config.GetString("database.redis.addr")
	pas := config.GetString("database.redis.password")
	db := config.GetInt("database.redis.db")

	RClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pas, // no password set
		DB:       db,  // use default DB
	})

	pong, err := RClient.Ping().Result()
	if err != nil {
		log.Fatal("ping failure",
			zap.String("pong", pong),
			zap.Error(err),
		)

	}
	redisdb = RClient
	if config.GetBool("gin.isDebugMode") {
		//TODO  如果是测试模式
		redisdb.FlushDBAsync() //清空本缓存库数据
	}
}

//Get 得到序列化对象
func Get(key string) (value string, err error) {
	if redisdb == nil {
		err = ErrRedisOff
		return
	}
	value, err = redisdb.Get(key).Result()
	return
}

//Set 反序列化对象
func Set(key string, value models.IModels, t time.Duration) (err error) {
	if redisdb == nil {
		err = ErrRedisOff
		return
	}
	err = redisdb.Set(key, value.ToJSON(), t).Err()

	return
}

//Close 关闭链接对象
func Close() {
	if redisdb != nil {
		redisdb.Close()
	}
}
