package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//  ------------------------------------
//  afterCreate hooks
// --------------------------------------

// User Hooks

// AfterCreate user hooks
func (u *User) AfterCreate(scope *gorm.Scope) (err error) {
	if u.ID > 0 {
		scope.DB().Create(&Profile{UserId: fmt.Sprint(u.ID)})
	}

	return
}
