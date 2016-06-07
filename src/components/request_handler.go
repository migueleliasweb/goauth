package components

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

//AccessToken Representation for the return of AuthHandler
type AccessToken struct {
	AccessToken string `json:"access_token"`
}

//AuthHandler Handles authorization
func AuthHandler(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params, jsonParams map[string]interface{}) {
	if username, usernameExists := jsonParams["username"]; usernameExists {
		if _, passwordExists := jsonParams["password"]; passwordExists {

			user := DB.GetUser(username.(string))

			if user != nil {
				JSONError(
					response,
					"User and/or password does not match.",
					http.StatusBadRequest)

				return
			}

			sha512Password := sha512.Sum512([]byte(jsonParams["password"].(string)))
			hexSha512Password := hex.EncodeToString(sha512Password[:])

			if hexSha512Password == user.Password {
				tokenString, tokenErr := GenerateJWTToken(user.Username, time.Hour*2)

				if tokenErr == nil {
					tokenString, _ := json.Marshal(AccessToken{AccessToken: tokenString})
					response.WriteHeader(http.StatusOK)
					response.Write([]byte(tokenString))
				} else {
					JSONError(
						response,
						"Could not create Json Web Token",
						http.StatusInternalServerError)
				}

				return
			}
		}
	}

	JSONError(
		response,
		"User and/or password does not match.",
		http.StatusBadRequest)

	return
}

//PermissionHandler Adds a permisson to a given user
func PermissionPUTHandler(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params, jsonParams map[string]interface{}) {

}
