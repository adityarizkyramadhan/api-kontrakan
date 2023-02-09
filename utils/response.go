package utils

type Response struct {
	Success bool
	Message string
	Body    interface{}
}

func ResponseWhenFail(message string) Response {
	return Response{
		Success: false,
		Message: message,
		Body:    nil,
	}
}

func ResponseWhenSuccess(message string, body interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Body:    body,
	}
}
