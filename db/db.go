package db

import (
	"os"

	"github.com/calz10/todolister/env/constants"
	"github.com/calz10/todolister/schemas"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

func Start() *gorm.DB {
	connString := os.Getenv(constants.CONN_URI)
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		defer db.Close()
	}

	db.AutoMigrate(schemas.Profile{}, schemas.Todo{}, schemas.User{})

	var adminUser schemas.User

	db.Take(&adminUser)

	if adminUser.ID == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte("admin1234"), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		db.Create(&schemas.User{Email: "todolistser@gmail.com", Password: string(hash), Username: "todolisteradmin"})
	}

	return db
}
