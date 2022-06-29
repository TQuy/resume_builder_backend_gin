package models

type Resume struct {
	BaseModel BaseModel `gorm:"embedded"`
	AuthorId  uint
	Author    User   `json:"author"`
	Name      string `json:"name"`
	Content   string `json:"content" gorm:"size:1000"`
}
