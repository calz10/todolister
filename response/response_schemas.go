package response

type Response struct {
	Success        bool        `json:"success"`
	Errors         interface{} `json:"errors"`
	ResponseStatus Status      `json:"status"`
	Result         interface{} `json:"result"`
}

type UserRegistrationResult struct {
	ID       int    `uri:"id" json:"_id"`
	Username string `uri:"username" json:"username"`
	Email    string `uri:"email" json:"email"`
}
