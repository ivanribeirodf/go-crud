package models

type User struct {
	ID 		uint	`json:"id" gorm:"primaryKey"`
	Nome 	string 	`json:"name"`
	Email 	string 	`json:"email" gorm:"unique"`
}