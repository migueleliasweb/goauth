package main

import (
	// "encoding/json"
	// "flag"
	// "fmt"
	// "io/ioutil"
	"log"
	"net/http"

	"components"

	// "github.com/garyburd/redigo/redis"

	"github.com/gorilla/mux"
)

func authHandlerFunc(response http.ResponseWriter, request *http.Request, jsonParams map[string]interface{}) {
	response.Write([]byte("Teste"))
}

func main() {
	gorillaRouter := mux.NewRouter()
	// Routes consist of a path and a handler function.

	authHandler := &components.Middleware{CallbackHandler: authHandlerFunc}

	gorillaRouter.Handle("/auth", authHandler).Methods("POST")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", gorillaRouter))
}
