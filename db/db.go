package db

import (
	"os"

	"github.com/calz10/todolister/env/constants"
	"github.com/calz10/todolister/schemas"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Start() *gorm.DB {
	connString := os.Getenv(constants.CONN_URI)
	db, _ := gorm.Open("postgres", connString)
	defer db.Close()

	db.CreateTable(schemas.User{})
	db.CreateTable(schemas.Profile{})
	db.CreateTable(schemas.Todo{})

	return db
}
