package controllers

type JSONError struct {
	Error string `json:"error"`
}

func APICustomError(c Context, code int, errMessage string) error {
	jsonError := JSONError{Error: errMessage}
	return c.JSON(code, jsonError)
}
