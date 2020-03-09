package schemas

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title       string
	Content     string
	CreatorId   string
	AssignedIds []string `gorm:"type:varchar(64)[]"`
	GroudIds    []string `gorm:"type:varchar(64)[]"`
	Type        string
}
