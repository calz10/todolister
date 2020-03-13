package validation

// ValidationError invalid input error format
type ValidationError struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

// ErrorsMap error codes for invalid inputs for code 500+
// Taking field name as first key and validation tag as second key returning the error code
// All map field keys must be in lowercase
var ValidationErrorsMap = map[string]map[string]ValidationError{
	"username": {
		"required": ValidationError{
			Message: "First Name is required",
			Field:   "FirstName",
		},
		"min": ValidationError{
			Message: "First name must have at least 2 characters in length",
			Field:   "firstName",
		},
		"max": ValidationError{
			Message: "First name must not exceed 32 characters in length",
			Field:   "FirstName",
		},
	},
	"email": {
		"required": ValidationError{
			Message: "Email is required",
			Field:   "Email",
		},
		"email": ValidationError{
			Message: "Email is invalid",
			Field:   "Email",
		},
	},
	"password": {
		"required": ValidationError{
			Message: "Password is required",
			Field:   "Password",
		},
		"min": ValidationError{
			Message: "Password must have at least 8 characters in length",
			Field:   "Password",
		},
		"max": ValidationError{
			Message: "Password must not exceed 32 characters in length",
			Field:   "Password",
		},
	},
	"newpassword": {
		"required": ValidationError{
			Message: "Password is required",
			Field:   "Password",
		},
		"min": ValidationError{
			Message: "Password must have at least 6 characters in length",
			Field:   "Password",
		},
		"max": ValidationError{
			Message: "Password must not exceed 32 characters in length",
			Field:   "Password",
		},
	},
}
