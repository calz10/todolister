package db

import (
	"fmt"

	"github.com/calz10/todolister/schemas"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Start() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=todolister password=qwerty sslmode=disable")
	fmt.Print(err)
	defer db.Close()

	db.AutoMigrate(schemas.User{})

	return db
}
