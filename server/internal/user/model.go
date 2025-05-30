package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"index"` // gorm:"uniqueIndex" Говорит что по данному параметру мы будем искать
	Password string
	Name     string
}
