package models

type Resume struct {
	BaseModel `gorm:"embedded"`
	AuthorId  uint   `json:",omitempty" gorm:"uniqueIndex:unique_idx_resume;not null"`
	Author    User   `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	Name      string `json:"name" gorm:"uniqueIndex:unique_idx_resume;not null"`
	Content   string `json:"content" gorm:"size:1000;not null"`
}
