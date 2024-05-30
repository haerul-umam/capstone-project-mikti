package web

type WebResponse struct {
	Code 		int 				`json:"code"`
	Message string 			`json:"message"`
	Data 		interface{} `json:"data"`
}

func ResponseToClient(code int, status string, data interface{}) WebResponse {
	return WebResponse{
		Code: code,
		Message: status,
		Data: data,
	}
}