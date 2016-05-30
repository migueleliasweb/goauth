package components

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Middleware Struct to perform common tasks for all routes
type Middleware struct {
	CallbackHandler func(http.ResponseWriter, *http.Request, map[string]interface{})
}

func (MiddleW *Middleware) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	bodyBytes, bodyError := ioutil.ReadAll(request.Body)

	if bodyError != nil {
		log.Fatalln(bodyError)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(http.StatusText(http.StatusBadRequest)))

		return
	}

	var jsonMap map[string]interface{}
	unmarshalError := json.Unmarshal(bodyBytes, &jsonMap)

	if unmarshalError != nil {
		log.Fatalln(bodyError)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid body."))

		return
	}

	MiddleW.CallbackHandler(response, request, jsonMap)
}
