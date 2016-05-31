package components

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//GoAuthHandler Type for using to route callbacks
type GoAuthHandler func(
	response http.ResponseWriter,
	request *http.Request,
	routeParams httprouter.Params,
	jsonParams map[string]interface{})

//GoAuthMiddleWare Default Middleware
func GoAuthMiddleWare(CH GoAuthHandler) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {

		bodyBytes, bodyError := ioutil.ReadAll(request.Body)

		if bodyError != nil {
			log.Println(bodyError)
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(http.StatusText(http.StatusBadRequest)))

			return
		}

		var jsonMap map[string]interface{}

		if len(bodyBytes) > 0 {
			unmarshalError := json.Unmarshal(bodyBytes, &jsonMap)

			if unmarshalError != nil {
				log.Println(bodyError)
				response.WriteHeader(http.StatusBadRequest)
				response.Write([]byte("Invalid body."))

				return
			}
		}

		CH(response, request, routeParams, jsonMap)
	}
}
