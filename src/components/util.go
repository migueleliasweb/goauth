package components

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//*GoAuthDB Pointer to the database
var DB *GoAuthDB

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

//SetupTestDB For test purposes
func SetupTestDB() {
	_permissions := make(map[string]*Permission)
	_permissions["admin"] = &Permission{Name: "admin"}

	_users := make(map[string]User, 1)
	_users["miguel"] = User{
		Username:    "miguel",
		Password:    "1ea3aa8cbcf860693559a55b36f041a879a3b2e62842ff49c4762ce860b98999b3d42b9ce3c7153bbe3b84f51950fe3b174b9db3f84d1f29039f901f69c6899f",
		Permissions: _permissions,
	}

	DB = InMemoryDatabase{
		Users:       _users,
		Permissions: _permissions,
	}
}
