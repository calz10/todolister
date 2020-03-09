package response

type Response struct {
	Response Status
	Errors   interface{}
	Result   interface{}
}
