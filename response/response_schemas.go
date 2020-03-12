package response

type Response struct {
	ResponseStatus Status
	Errors         interface{}
	Result         interface{}
}
