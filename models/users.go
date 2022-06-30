package models

type User struct {
	BaseModel `gorm:"embedded"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
