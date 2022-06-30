package models

type User struct {
	BaseModel `gorm:"embedded"`
	Username  string `json:"username" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
}
