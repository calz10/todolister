package schemas

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Todo struct {
	gorm.Model
	Title       string
	Content     string
	CreatorId   string
	AssignedIds pq.StringArray `gorm:"type:varchar(64)[]"`
	GroudIds    pq.StringArray `gorm:"type:varchar(64)[]"`
	Type        string
}
