package user

import (
	"time"

	"task-system/database"
	"task-system/log"
)
//User 用户结构体
type User struct {
	ID        int       `xorm:"pk autoincr" db:"id" json:"id"`
	Name      string    `xorm:"varchar(12) not null        'name' comment('用户名')"           db:"name"       json:"name"         form:"name"       binding:"required"`
	Emeil     string    `xorm:"varchar(25) not null unique 'emeil' comment('用户的邮箱')"       db:"emeil"      json:"emeil"        form:"emeil"      binding:"required"`
	Password  string    `xorm:"varchar(12) not null         'password' comment('用户的密码')"   db:"password"   json:"password"     form:"password"   binding:"required"`
	Avatar    string    `xorm:"varchar(25)  'avatar' comment('用户头像地址')"                   db:"avatar"    json:"avatar"        form:"avatar"`
	Created   time.Time `xorm:"created"                                                      db:"created"    json:"created"       form:"created"`
	Updated   time.Time `xorm:"updated"                                                      db:"updated"    json:"updated"       form:"updated"`
	Authority int       `xorm:"int(1) not null 'authority' comment('权限')"                   db:"authority"  json:"authority"     form:"authority"  binding:"required"`
}
//AddUser 添加用户
func (user *User) AddUser() (id int64, err error) {
	res, err :=database.Db.NamedExec("insert into user(name,emeil,password,avatar,created,authority) values(:name,:emeil,:password,:avatar,:created,:authority)", user)
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
//DelUser 删除用户
func (user *User) DelUser() (err error) {
	return nil
}

//UpUser 更新用户数据
func (user *User) UpUser() (err error) {
	return nil
}

//GetUser 得到用户数据
func (user *User) GetUser(emeil string) (err error) {
	err = database.Db.Get(&user, "select * from user where emeil=?", user.Emeil)
	if err != nil {
		log.Logger.Errorln(err)
		return err
	}
	return nil
}
//GetUsers 得到全部的用户数据
func (user *User) GetUsers() (users []User, err error) {
	return nil, nil
}
