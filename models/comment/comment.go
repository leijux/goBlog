package comment

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Email   string `gorm:"size:30;not null"       `
	BlogID  uint   `gorm:"not null"`
	content string
}
