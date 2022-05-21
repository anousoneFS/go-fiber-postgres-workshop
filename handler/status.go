package handler

type StatusMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

var RequestSuccess = StatusMessage{
	Code:    200,
	Message: "Successful request",
	Status:  "OK",
}

var Created = StatusMessage{
	Code:    201,
	Message: "Request succeeded, and a new resource was created",
	Status:  "OK",
}

var ErrBodyParserFailure = StatusMessage{
	Code:    400,
	Message: "Binds the request body to a struct, the request body is not a valid JSON",
	Status:  "Bad Request",
}
