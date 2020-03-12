package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type DbService struct {
	*gorm.DB
}

func (db *DbService) InitializeApi(router *gin.Engine) {
	router.POST("/users", db.UserRegistrationHandler)
}
