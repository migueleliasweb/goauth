package components

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//AuthHandler Handles authorization
func AuthHandler(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params, jsonParams map[string]interface{}) {
	//encodedPassword := sha512.Sum512([]byte(jsonParams["username"].(string)))

	if username, usernameExists := jsonParams["username"]; usernameExists {
		if _, passwordExists := jsonParams["username"]; passwordExists {
			user, err := GetUser(username.(string))

			if err == nil {
				response.Write([]byte(user.Password))
			} else {
				response.Write([]byte(err.Error()))
			}

			return
		}
	}

	response.WriteHeader(http.StatusBadRequest)
	response.Write([]byte(http.StatusText(http.StatusBadRequest)))
}
