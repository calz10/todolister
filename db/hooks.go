package db

import (
	"fmt"

	"github.com/calz10/todolister/schemas"
	"github.com/jinzhu/gorm"
)

func (u User) AfterCreate(scope *gorm.Scope) (err error) {
	if u.ID > 0 {
		scope.DB().Create(&schemas.Profile{UserId: fmt.Sprint(u.ID)})
	}
	return
}
