package components

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

//AccessToken Representation for the return of AuthHandler
type AccessToken struct {
	AccessToken string
}

//AuthHandler Handles authorization
func AuthHandler(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params, jsonParams map[string]interface{}) {
	if username, usernameExists := jsonParams["username"]; usernameExists {
		if _, passwordExists := jsonParams["password"]; passwordExists {

			user, userErr := GetUser(username.(string))

			if userErr != nil {
				JSONError(
					response,
					"User and/or password does not match.",
					http.StatusBadRequest)

				return
			}

			sha512Password := sha512.Sum512([]byte(jsonParams["password"].(string)))
			hexSha512Password := hex.EncodeToString(sha512Password[:])

			if hexSha512Password == user.Password {
				JWTToken := jwt.New(jwt.SigningMethodHS256)

				JWTToken.Claims["iat"] = time.Now()
				JWTToken.Claims["exp"] = time.Now().Add(time.Hour * 2)
				JWTToken.Claims["identity"] = user.Hash

				tokenString, tokenErr := JWTToken.SignedString([]byte(JWTSecret))

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
			} else {
				JSONError(
					response,
					"User and/or password does not match.",
					http.StatusBadRequest)
			}

			return
		}
	}

	JSONError(
		response,
		"Something went wrong.",
		http.StatusInternalServerError)
}
