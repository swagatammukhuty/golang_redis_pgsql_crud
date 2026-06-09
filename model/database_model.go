package model

type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}
