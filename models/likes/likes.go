package likes

import "gorm.io/gorm"

type Likes struct {
	gorm.Model
	BlogID uint `gorm:"not null"`
	Email string `gorm:"size:30;not null"       `
}
