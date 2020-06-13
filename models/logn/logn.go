package logn

import (
	"goBlog/database"
	"goBlog/log"
	"goBlog/models/user"
)

//Logn 登入结构体
type Logn struct {
	Emeil string `db:"emeil"    form:"emeil" json:"emeil" binding:"required"`
	Pwd   string `db:"password" form:"pwd"   json:"pwd"   binding:"required"`
}

//PwdCheck 验证登入
func (login *Logn) PwdCheck() (b bool, user user.User, err error) {
	err = database.Db.Get(&user, "select * from user where emeil=? and password=?", login.Emeil, login.Pwd)
	if err != nil {
		log.Logger.Errorln(err)
		b = false
		return
	}
	if user.Emeil == "" {
		b = false
		return
	}

	b = true
	user.Password = ""
	return
}

//EmeilChack 邮箱验证
func (login *Logn) EmeilChack() (b bool) {
	var u user.User
	database.Db.Get(&u, "select emeil from user where emeil=? ", login.Emeil)
	if u.Emeil == "" {
		b = true
		return
	}
	b = false
	return
}
