package user

import (
	"time"

	"task-system/database"
	"task-system/log"
)

type User struct {
	Id        int       `xorm:"pk autoincr" db:"id" json:"id"`
	Name      string    `xorm:"varchar(12) not null        'name' comment('用户名')"         db:"name"  json:"name"`
	Emeil     string    `xorm:"varchar(25) not null unique 'emeil' comment('用户的邮箱')"     db:"emeil"     json:"emeil"`
	Password  string    `xorm:"varchar(12) not null         'password' comment('用户的密码')" db:"password"  json:"password"`
	Avatar    string    `xorm:"varchar(25)  'avatar' comment('用户头像地址')"                  db:"avatar"    json:"avatar"`
	Created   time.Time `xorm:"created"                                                      db:"created"  json:"created"`
	Updated   time.Time `xorm:"updated"                                                      db:"updated"  json:"updated"  `
	Authority int       `xorm:"int(1) not null 'authority' comment('权限')"                   db:"authority" json:"authority"`
}

func (user *User) AddUser() (id int64, err error) {
	res, err :=database.Db.NamedExec("insert into user values(:name,:emeil,:password,:avatar,:creaTime,:authority)", user)
	if err != nil {
		log.Logger.Errorln(err)
		return 0, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		log.Logger.Errorln(err)
		return 0, err
	}
	return
}

func (user *User) DelUser() (err error) {
	return nil
}

func (u *User) UpUser() (err error) {
	return nil
}

func (user *User) GetUser(emeil string) (err error) {
	err = database.Db.Get(&user, "select * from user where emeil=?", user.Emeil)
	if err != nil {
		log.Logger.Errorln(err)
		return err
	}
	return nil
}

func (user *User) GerUsers() (users []User, err error) {
	return nil, nil
}
