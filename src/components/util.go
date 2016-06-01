package components

import (
	"encoding/json"
	"net/http"
)

type jsonError struct {
	Error string
}

//JSONError Helper function to return restful errors
func JSONError(response http.ResponseWriter, errorString string, statusCode int) {

	errorJSONString, _ := json.Marshal(jsonError{Error: errorString})

	response.WriteHeader(statusCode)
	response.Write([]byte(errorJSONString))
}
