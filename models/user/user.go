package user

import (
	
	"time"

	"goBlog/database/cache"
	"goBlog/database/orm"
	"goBlog/log"
	"goBlog/models"
	"goBlog/models/blog"
	"goBlog/src/common"
	
	json "github.com/json-iterator/go"
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	emptyString = ""
)

var (
	ErrDataEmpty    = errors.New("data is empty!")
	ErrEmailIsEmpty = errors.New("Email is empty!")
	ErrNoNewRecord  = errors.New("no new record")
)

//User 用户结构体
type User struct {
	gorm.Model
	Name      string      `gorm:"size:15;not null"                       `
	Email     string      `gorm:"size:30;not null;unique"                `
	Password  []byte      `gorm:"type:binary(32);not null"               `
	Image     string      `gorm:"size:50"                                `
	Authority int         `                                              `
	Blogs     []blog.Blog `gorm:"foreignkey:Email;references:Email"      `
}

type UserApi struct {
	Name     string `json:"name"      form:"name"      binding:"required"      `
	Email    string `json:"email"     form:"email"     binding:"required,email"`
	Password string `json:"password"  form:"password"  binding:"required"      `
	Image    string `json:"image"     form:"image"                             `
}

var _ models.IModels = &User{}

func NewUser() UserApi {
	return UserApi{}
}

func (u User) ToUserApi() UserApi {
	return UserApi{
		Name:     u.Name,
		Email:    u.Email,
		Password: emptyString,
		Image:    u.Image,
	}
}

func (u UserApi) ToUser() *User {
	var dk []byte
	if u.Password != emptyString {
		var err error
		dk, err = common.Scrypt(u.Password)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return &User{
		Name:     u.Name,
		Email:    u.Email,
		Password: dk,
		Image:    u.Image,
	}
}

//CreateUser 添加用户
func (user *UserApi) CreateUser() (bool, error) {
	//res, err := database.Db.NamedExec("insert into user(name,email,password,avatar,created,authority) values(:name,:email,:password,:avatar,:created,:authority)", user)
	u := user.ToUser()
	return createUser(u)
}

func createUser(u *User) (bool, error) {
	if result := orm.Create(u); result.Error != nil {
		return false, errors.Wrap(result.Error, "CreateUser")
	}
	return true, nil
	//return false, ErrNoNewRecord
}

//DelUser 删除用户 *权限管理
func (user *User) DelUser() (l int64, err error) {
	// res, err := database.Db.Exec("delete form user where email=?", user.Email)
	// if err != nil {
	// 	err = errors.Wrap(err, "DelUser err")
	// 	log.Errorln(err)
	// 	return
	// }
	// l, err = res.RowsAffected()
	// err = errors.Wrap(err, "DelUser err")
	// return
	return
}

//UpUser 更新用户数据
func (user *User) UpUser() (err error) {
	return nil
}

//GetUser 得到用户数据
func (user *UserApi) GetUser() error {
	u := user.ToUser()
	err := getUser(u)
	*user = u.ToUserApi()
	return err
}

func getUser(u *User) error {
	if u.Email == "" {
		return ErrEmailIsEmpty
	}

	if res, b := userWithCache(u.Email); b {
		u.FromJSON(res)
		log.Infoln("cache read success")
	}
	err := orm.Db.Where("email = ?", u.Email).First(&u).Error
	if err != nil {
		log.Errorln(err)
		return errors.Cause(err)
	}
	err = cache.Set(u.Email, u, 1*time.Hour) //写入缓存
	if err != nil {
		return errors.Wrap(err, "cache write err")
	}
	return nil
}

func userWithCache(email string) (string, bool) {
	switch res, Rederr := cache.Get(email); {
	case Rederr == redis.Nil:
		return res, true
	case Rederr == cache.ErrRedisOff:
		return emptyString, false
	default:
		log.Errorln(Rederr)
		return emptyString, false
	}
}

//GetUsers 得到全部的用户数据
func (user UserApi) GetUsers() (users []User, err error) {
	return nil, nil
}
//用户有多少文章
func (user UserApi) Xx() (i int64,err error) {
	u := user.ToUser()
	//u.ID = 1
	u.Email = "123@122f213.com"
	association := orm.Db.Model(&u).Association("Blogs")
	i=association.Count()
	err = association.Error
	if err != nil {
		return
	}
	return
}

//ToJSON ...
func (user *User) ToJSON() string {
	j, err := json.Marshal(user)
	if err != nil {
		log.Debugln(err)
		return emptyString
	}
	return string(j)
}

func (user *User) FromJSON(data string) {
	err := json.Unmarshal([]byte(data), user)
	if err != nil {
		log.Debugln(err)
	}
	return
}
