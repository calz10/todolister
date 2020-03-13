package routes

import (
	"github.com/calz10/todolister/models"
	"github.com/calz10/todolister/request"
	"github.com/calz10/todolister/response"
	"github.com/calz10/todolister/validation"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserRegistrationHandler
func (db *DbService) UserRegistrationHandler(c *gin.Context) {
	var request request.UserRequestBody
	var result interface{}
	var errors interface{}
	var statusCode int

	defer func() {
		if err := recover(); err != nil {
			statusCode = 500

			c.JSON(statusCode, response.Response{
				ResponseStatus: response.StatusMap[statusCode],
				Errors:         err,
			})
		}
	}()

	c.BindJSON(&request)

	validationErrors, hasError := validation.Validate(request)

	if hasError {
		errors = validationErrors
		statusCode = response.ValidationErr
	} else {

		var userSchema models.User
		db.Where("email = ?", request.Email).Or("username = ?", request.Username).Find(&userSchema)

		if userSchema.Username != "" {
			statusCode = response.ExistingData

			var customedResponse response.Status
			customedResponse = response.StatusMap[statusCode]

			customedResponse.Message = "Email or username was already taken"

			c.JSON(statusCode, response.Response{
				ResponseStatus: customedResponse,
				Errors:         errors,
				Result:         nil,
			})
			return
		} else {
			hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
			err := db.Create(&models.User{Email: request.Email, Password: string(hash), Username: request.Username})
			if err != nil {
				db.Take(&userSchema, &models.User{Email: request.Email})
				statusCode = response.Created
				result = response.UserRegistrationResult{
					Email:    request.Email,
					Username: request.Username,
					ID:       int(userSchema.ID),
				}
			}
		}
	}

	c.JSON(statusCode, response.Response{
		ResponseStatus: response.StatusMap[statusCode],
		Errors:         errors,
		Result:         result,
		Success:        response.StatusMap[statusCode].Success,
	})

}

func (db *DbService) UserLoginHandler(c *gin.Context) {
	// s := claims.GenerateNewToken("test")
	// t, _ := claims.VerifyToken(s)
	// fmt.Println(t)
}
