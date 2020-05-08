package models

import (
	//"task-system/models/logn"
	//"task-system/models/user"
)

type IModels interface {
	ToJSON() string
	FromJSON(string)
}

//var _ IModels = &user.User{}

// func New() {//工厂方法

// }
