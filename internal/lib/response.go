package lib

type errors struct {
	Errors []errorItem `json:"errors"`
}

type errorItem struct {
	Message string `json:"message"`
}

func ErrorResponse(text string) interface{} {
	item := errorItem{Message: text}
	errors := errors{
		Errors: []errorItem{item},
	}
	return errors
}
