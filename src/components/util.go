package components

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jsonError struct {
	Error string `json:"error"`
}

//JSONError Helper function to return restful errors
func JSONError(response http.ResponseWriter, errorString string, statusCode int) {

	errorJSONString, _ := json.Marshal(jsonError{Error: errorString})

	response.WriteHeader(statusCode)
	response.Write([]byte(errorJSONString))
}

//GenerateJWTToken Generates the authenticaton token (JWT)
func GenerateJWTToken(subject string, extraTime time.Duration) (tokenString string, err error) {
	JWTToken := jwt.New(jwt.SigningMethodHS256)

	JWTToken.Claims["iat"] = time.Now()
	JWTToken.Claims["exp"] = time.Now().Add(extraTime)
	JWTToken.Claims["sub"] = subject

	return JWTToken.SignedString([]byte(JWTSecret))
}
