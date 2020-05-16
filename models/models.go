package models

type IModels interface {
	ToJSON() string
	FromJSON(string)
}

// var (
// 	User = "user"
// )

// func New(a string) Modelser { //工厂方法
// 	switch a {
// 	case "user":
// 		return new(user.User)
// 	default:
// 		log.Logger.Print("")
// 		return nil
// 	}
// }
