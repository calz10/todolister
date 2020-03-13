package response

type Response struct {
	ResponseStatus Status      `json:"status"`
	Errors         interface{} `json:"errors"`
	Result         interface{} `json:"result"`
}
