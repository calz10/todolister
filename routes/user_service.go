package routes

import (
	"fmt"

	"github.com/calz10/todolister/request"
	"github.com/calz10/todolister/response"
	"github.com/calz10/todolister/schemas"
	"github.com/calz10/todolister/validation"
	"github.com/gin-gonic/gin"
)

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
	} else {

		var userSchema schemas.User
		db.Where("email=?", request.Email).Or("username=?", request.Username).Find(&userSchema)

		if userSchema.Username != "" {
			statusCode = response.ExistingData

			var customedResponse response.Status
			customedResponse = response.StatusMap[statusCode]

			customedResponse.Message = "Either email or username was already taken"

			c.JSON(statusCode, response.Response{
				ResponseStatus: customedResponse,
				Errors:         errors,
				Result:         nil,
			})
			return
		} else {
			fmt.Print(request)
		}
	}

	c.JSON(statusCode, response.Response{
		ResponseStatus: response.StatusMap[statusCode],
		Errors:         errors,
		Result:         result,
	})

}
