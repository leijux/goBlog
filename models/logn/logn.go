package login

import (
	"task-system/database"
	"task-system/log"
	"task-system/models/user"
)

//Login 登入结构体
type Login struct {
	Emeil string `db:"emeil"    form:"emeil" json:"emeil" binding:"required"`
	Pwd   string `db:"password" form:"pwd"   json:"pwd"   binding:"required"`
}

//PwdCheck 验证登入
func (login *Login) PwdCheck() (b bool, user user.User, err error) {
	err = database.Db.Get(&user, "select * from user where emeil=? and password=?", login.Emeil, login.Pwd)
	if err != nil {
		log.Logger.Errorln(err)
	}
	if user.Emeil == "" {
		b = false
		return
	}

	b = true
	user.Password = ""
	return
}
