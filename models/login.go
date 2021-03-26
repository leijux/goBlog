package models

import (
	"goBlog/database/orm"
	"goBlog/src/common"
)

const (
	emptyString = ""
)

//login 登入结构体
type Login struct {
	Email string `gorm:"email"   `
	Pwd   []byte `gorm:"password"`
}

type LoginApi struct {
	Email string `form:"email" json:"email" binding:"required,email"`
	Pwd   string `form:"pwd"   json:"pwd"   binding:"required"`
}

func NewLogin() LoginApi {
	return LoginApi{}
}

func (l LoginApi) toLogin() Login {
	dk, _ := common.Scrypt(l.Pwd)
	return Login{
		Email: l.Email,
		Pwd:   dk,
	}
}

//PwdCheck 验证登入
func (login *LoginApi) PwdCheck() (b bool, userApi UserApi, err error) {
	l := login.toLogin()
	u := User{}

	//err = database.Db.Get(&user, "select * from user where email=? and password=?", login.Email, login.Pwd)
	err = orm.Db.Where("email = ? and password=?", l.Email, l.Pwd).First(&u).Error
	if err != nil {
		b = false
		return
	}
	b = true
	userApi = u.ToUserApi()
	return
}

//EmailCheck 邮箱验证
func (login *LoginApi) EmailCheck() (b bool, err error) {
	u := User{}
	//database.Db.Get(&u, "select email from user where email=? ", login.Email)
	err = orm.Db.Where("email=?", login.Email).First(&u).Error
	if err != nil {
		b = false
		return
	}
	if u.Email == emptyString {
		b = true
		return
	}
	b = false
	return
}
