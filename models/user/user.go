package user

import (
	"encoding/json"
	"time"

	"task-system/database"
	"task-system/database/cache"
	myerr "task-system/err"
	"task-system/log"
	"task-system/models"

	"github.com/go-redis/redis/v7"
)

//User 用户结构体
type User struct {
	ID        int       ` db:"id" json:"id"`
	Name      string    `db:"name"       json:"name"         form:"name"       binding:"required"`
	Emeil     string    `db:"emeil"      json:"emeil"        form:"emeil"      binding:"required"`
	Password  string    `db:"password"   json:"password"     form:"password"   binding:"required"`
	Avatar    string    `db:"avatar"     json:"avatar"         form:"avatar"`
	Created   time.Time `db:"created"    json:"created"       form:"created"`
	Updated   time.Time `db:"updated"    json:"updated"       form:"updated"`
	Authority int       `db:"authority"  json:"authority"     form:"authority"  binding:"required"`
}

var _ models.IModels = &User{}

//AddUser 添加用户
func (user *User) AddUser() (id int64, err error) {
	res, err := database.Db.NamedExec("insert into user(name,emeil,password,avatar,created,authority) values(:name,:emeil,:password,:avatar,:created,:authority)", user)
	if err != nil {
		log.Logger.Errorln(err)
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		log.Logger.Errorln(err)
		return
	}
	return
}

//DelUser 删除用户 *权限管理
func (user *User) DelUser() (l int64, err error) {
	res, err := database.Db.Exec("delete form user where emeil=?", user.Emeil)
	if err != nil {
		log.Logger.Errorln(err)
		return
	}
	l, err = res.RowsAffected()
	return
}

//UpUser 更新用户数据
func (user *User) UpUser() (err error) {
	return nil
}

//GetUser 得到用户数据
func (user *User) GetUser() (err error) {
	if user.Emeil == "" {
		return myerr.ErrStringIsEmpty
	}

	res, Rederr := cache.Get(user.Emeil)
	if Rederr == redis.Nil || Rederr == cache.ErrRedisOff { //判断是否是空的
		log.Logger.Errorln(Rederr)
		err = database.Db.Get(user, "select * from user where emeil=?", user.Emeil)
		if err != nil {
			log.Logger.Errorln(err)
			return
		}
		err = cache.Set(user.Emeil, user, 1*time.Hour) //写入缓存
		return
	}
	user.FromJSON(res)
	log.Logger.Infoln(res, "缓存读取成功")
	return
}

//GetUsers 得到全部的用户数据
func (user *User) GetUsers() (users []User, err error) {
	return nil, nil
}

//ToJSON ...
func (user *User) ToJSON() string {
	j, err := json.Marshal(user)
	if err != nil {
		log.Logger.Debugln(err)
		return ""
	}
	return string(j)
}

func (user *User) FromJSON(data string) {
	err := json.Unmarshal([]byte(data), user)
	if err != nil {
		log.Logger.Debugln(err)
		return
	}
	return
}
