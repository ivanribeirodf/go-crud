package models

type User struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Email  string `json:"email" gorm:"unique"`
	Passwd string `json:"-"`
	Role   string `json:"role" gorm:"default:user"`
}
