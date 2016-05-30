package components

import "net/http"

//Middleware Struct to perform common tasks for all routes
type Middleware struct {
	CallbackHandler func(http.ResponseWriter, *http.Request, map[string]interface{})
}

func (MiddleW *Middleware) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	jsonMap := map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": []interface{}{
			"Gomez",
			"Morticia",
		},
	}

	MiddleW.CallbackHandler(response, request, jsonMap)
}
