package components

import (
	"crypto/sha512"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

//AuthHandler Handles authorization
func AuthHandler(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params, jsonParams map[string]interface{}) {

	if username, usernameExists := jsonParams["username"]; usernameExists {
		if _, passwordExists := jsonParams["password"]; passwordExists {

			user, userErr := GetUser(username.(string))

			if userErr != nil {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte(http.StatusText(http.StatusInternalServerError)))

				return
			}

			sha512Password := sha512.Sum512([]byte(jsonParams["password"].(string)))
			hexSha512Password := hex.EncodeToString(sha512Password[:])

			if hexSha512Password == user.Password {
				JWTToken := jwt.New(jwt.SigningMethodHS256)

				JWTToken.Claims["iat"] = time.Now()
				JWTToken.Claims["exp"] = time.Now().Add(time.Hour * 2)
				JWTToken.Claims["identity"] = user.Hash

				tokenString, tokenErr := JWTToken.SignedString([]byte("asdasdasd"))

				if tokenErr == nil {
					response.Write([]byte(tokenString))
				} else {
					response.WriteHeader(http.StatusInternalServerError)
					response.Write([]byte(tokenErr.Error()))
				}

				return
			}
		}
	}

	log.Println("Else")
	response.WriteHeader(http.StatusBadRequest)
	response.Write([]byte(http.StatusText(http.StatusBadRequest)))
}
