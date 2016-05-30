package components

import "net/http"

//AuthHandler Struct that handles authorization
type AuthHandler struct {
	Middleware
}

//CallbackHandler Handles authorization
func (MW *Middleware) CallbackHandler(response http.ResponseWriter, request *http.Request, jsonParams map[string]interface{}) {
	//encodedPassword := sha512.Sum512([]byte(jsonParams["username"].(string)))

	user, err := GetUser(jsonParams["username"].(string))

	if err == nil {
		response.Write([]byte(user.Name))
	} else {
		response.Write([]byte(err.Error()))
	}
}
