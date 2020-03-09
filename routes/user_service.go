package routes

import (
	"github.com/gin-gonic/gin"
)

type Test struct {
	Response string
}

func (db *DbService) UserRegistrationHandler(c *gin.Context) {
	c.JSON(400, &Test{
		Response: "OK",
	})
}
