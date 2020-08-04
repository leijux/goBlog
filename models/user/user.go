package user

import (
	"encoding/json"
	"time"

	"goBlog/database"
	"goBlog/database/cache"
	"goBlog/database/orm"
	myerr "goBlog/err"
	"goBlog/log"
	"goBlog/models"
	"goBlog/src/common"

	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

//User 用户结构体
type User struct {
	gorm.Model
	Name      string `gorm:"size:;not null"                       db:"name"      `
	Email     string `gorm:"not null;unique"                      db:"email"     `
	Password  []byte `gorm:"size:30;type:binary;not null"         db:"password"  `
	Avatar    string `                                            db:"avatar"    `
	Authority int    `                                            db:"authority"  `
}
type UserApi struct {
	Name     string `json:"name"         form:"name"       binding:"required"`
	Email    string `json:"email"        form:"email"      binding:"required,email"`
	Password string `json:"password"     form:"password"   binding:"required"`
	Avatar   string `json:"avatar"       form:"avatar"`
}

var _ models.IModels = &User{}

func (u User) ToUserApi() UserApi {
	return UserApi{
		Name:     u.Name,
		Email:    u.Email,
		Password: "",
		Avatar:   u.Avatar,
	}
}

func (u UserApi) ToUser() *User {
	dk, _ := common.Scrypt(u.Password)
	return &User{
		Name:     u.Name,
		Email:    u.Email,
		Password: dk,
		Avatar:   u.Avatar,
	}
}

//AddUser 添加用户
func (user *UserApi) AddUser() (b bool, err error) {
	//res, err := database.Db.NamedExec("insert into user(name,email,password,avatar,created,authority) values(:name,:email,:password,:avatar,:created,:authority)", user)
	u := user.ToUser()
	if orm.Db.NewRecord(u) {
		g := orm.Db.Create(u)
		err = g.Error
	}
	b = orm.Db.NewRecord(user) // => 创建`user`后返回`false`
	return
}

//DelUser 删除用户 *权限管理
func (user *User) DelUser() (l int64, err error) {
	res, err := database.Db.Exec("delete form user where email=?", user.Email)
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
	if user.Email == "" {
		return myerr.ErrStringIsEmpty
	}

	res, Rederr := cache.Get(user.Email)
	if Rederr == redis.Nil || Rederr == cache.ErrRedisOff { //判断是否是空的
		log.Logger.Errorln(Rederr)
		err = database.Db.Get(user, "select * from user where.Email=?", user.Email)
		if err != nil {
			log.Logger.Errorln(err)
			return
		}
		user.Password = nil                            //密码为空的
		err = cache.Set(user.Email, user, 1*time.Hour) //写入缓存
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
