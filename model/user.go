package model

type User struct {
	ID     int    `gorm:"id"`
	Email  string `gorm:"email"`
	Mobile string `gorm:"mobile"`
}
