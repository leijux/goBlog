package login

import (
	"goBlog/database"
	"goBlog/log"
	"goBlog/models/user"
	"goBlog/src/common"
)

//login 登入结构体
type Login struct {
	Email string `db:"email"   `
	Pwd   []byte `db:"password"`
}
type LoginApi struct {
	Email string `form:"email" json:"email" binding:"required"`
	Pwd   string `form:"pwd"   json:"pwd"   binding:"required"`
}

//TODO写到这里
func (l LoginApi) toLogin() Login {
	dk, _ := common.Scrypt(l.Pwd)
	return Login{
		Email: l.Email,
		Pwd:   dk,
	}
}

//PwdCheck 验证登入
func (login *Login) PwdCheck() (b bool, user user.User, err error) {

	err = database.Db.Get(&user, "select * from user where email=? and password=?", login.Email, login.Pwd)
	if err != nil {
		log.Logger.Errorln(err)
		b = false
		return
	}
	if user.Email == "" {
		b = false
		return
	}

	b = true
	user.Password = nil
	return
}

//EmailCheck 邮箱验证
func (login *Login) EmailCheck() (b bool) {
	var u user.User
	database.Db.Get(&u, "select email from user where email=? ", login.Email)
	if u.Email == "" {
		b = true
		return
	}
	b = false
	return
}
