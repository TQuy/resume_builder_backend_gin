package models

type Resume struct {
	BaseModel `gorm:"embedded"`
	AuthorId  uint   `gorm:"uniqueIndex:unique_idx_resume"`
	Author    User   `json:"author"`
	Name      string `json:"name" gorm:"uniqueIndex:unique_idx_resume"`
	Content   string `json:"content" gorm:"size:1000"`
}
