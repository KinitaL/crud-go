package models

type User struct {
	ID           uint `gorm:"primarykey"`
	Access_token string
}
