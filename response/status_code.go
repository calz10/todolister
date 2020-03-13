package response

// Status status codes and messages
type Status struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"-"`
}

// StatusCodes code keys for mapping
var StatusCodes = []int{
	100, 200, 201, 202, 203, 204, 300, 301, 400,
	401, 402, 403, 404, 405, 406, 407, 408, 409,
	500, 501, 502, 503, 504, 600,
}

// StatusMap status code map for code 1xx-500
var StatusMap = map[int]Status{
	100: Status{
		Code:    100,
		Message: "Continue",
		Success: true,
	},
	200: Status{
		Code:    200,
		Message: "OK",
		Success: true,
	},
	201: Status{
		Code:    201,
		Message: "Created",
		Success: true,
	},
	202: Status{
		Code:    202,
		Message: "Accepted",
		Success: true,
	},
	203: Status{
		Code:    203,
		Message: "Partial Information",
		Success: true,
	},
	204: Status{
		Code:    204,
		Message: "No Content",
		Success: true,
	},
	300: Status{
		Code:    300,
		Message: "Multiple Choices",
		Success: true,
	},
	301: Status{
		Code:    301,
		Message: "Moved Permanently",
		Success: true,
	},
	400: Status{
		Code:    400,
		Message: "Bad Request",
		Success: false,
	},
	401: Status{
		Code:    401,
		Message: "Unauthorized",
		Success: false,
	},
	402: Status{
		Code:    402,
		Message: "Payment Required",
		Success: false,
	},
	403: Status{
		Code:    403,
		Message: "Forbidden",
		Success: false,
	},
	404: Status{
		Code:    404,
		Message: "Not Found",
		Success: false,
	},
	405: Status{
		Code:    405,
		Message: "Method Not Allowed",
		Success: false,
	},
	406: Status{
		Code:    406,
		Message: "Not Acceptable",
		Success: false,
	},
	407: Status{
		Code:    407,
		Message: "Request Time-Out",
		Success: false,
	},
	408: Status{
		Code:    408,
		Message: "Authentication Failed",
		Success: false,
	},
	409: Status{
		Code:    409,
		Message: "Data already exist",
		Success: false,
	},
	500: Status{
		Code:    500,
		Message: "Internal Server Error",
		Success: false,
	},
	501: Status{
		Code:    501,
		Message: "Not Implemented",
		Success: false,
	},
	502: Status{
		Code:    502,
		Message: "Bad Gateway",
		Success: false,
	},
	503: Status{
		Code:    503,
		Message: "Service Unavailable",
		Success: false,
	},
	504: Status{
		Code:    504,
		Message: "Gateway Timeout",
		Success: false,
	},
	600: Status{
		Code:    600,
		Message: "Validation Error",
		Success: false,
	},
}
