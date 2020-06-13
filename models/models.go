package models

type IModels interface {
	ToJSON() string
	FromJSON(string)
}
