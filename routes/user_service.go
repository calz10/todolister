package routes

import (
	"github.com/calz10/todolister/request"
	"github.com/calz10/todolister/response"
	"github.com/calz10/todolister/validation"
	"github.com/gin-gonic/gin"
)

type Test struct {
	Response string
}

// UserRegistrationHandler
func (db *DbService) UserRegistrationHandler(c *gin.Context) {
	var request request.UserRequestBody
	var result interface{}
	var errors interface{}
	var statusCode int

	c.BindJSON(&request)

	validationErrors, hasError := validation.Validate(request)

	if hasError {
		errors = validationErrors
		statusCode = response.ValidationErr
	}

	c.JSON(statusCode, response.Response{
		ResponseStatus: response.StatusMap[statusCode],
		Errors:         errors,
		Result:         result,
	})

}
