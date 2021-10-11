package models

type News struct {
	ID     uint `gorm:"primarykey"`
	Name   string
	Body   string
	Rating int
}
