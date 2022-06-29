package models

type User struct {
	BaseModel BaseModel `gorm:"embedded"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
}
