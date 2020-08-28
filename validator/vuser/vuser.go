package vuser

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func VuserName(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}
