package server

type ResultData struct {
	Code    string
	Message string
}

func Message(code string, message string) *ResultData {
	return &ResultData{
		Code:    code,
		Message: message,
	}
}
