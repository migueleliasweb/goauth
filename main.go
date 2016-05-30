package main

import (
	"components"
	"log"
	"net/http"

	// "github.com/garyburd/redigo/redis"

	"github.com/gorilla/mux"
)

func authHandlerFunc(response http.ResponseWriter, request *http.Request, jsonParams map[string]interface{}) {
	//encodedPassword := sha512.Sum512([]byte(jsonParams["username"].(string)))

	user, err := components.GetUser(jsonParams["username"].(string))

	if err == nil {
		response.Write([]byte(user.Name))
	} else {
		response.Write([]byte(err.Error()))
	}

	//convert encodedPassword to slice
	//response.Write(encodedPassword[:])

}

func main() {
	gorillaRouter := mux.NewRouter()
	authHandler := &components.Middleware{CallbackHandler: authHandlerFunc}
	gorillaRouter.Handle("/auth", authHandler).Methods("POST")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", gorillaRouter))
}
