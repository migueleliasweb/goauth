package main

import (
	// "encoding/json"
	// "flag"
	// "fmt"
	// "io/ioutil"
	"log"
	"net/http"

	// "github.com/garyburd/redigo/redis"

	"components"

	"github.com/gorilla/mux"
)

func yourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	a = components.Debug

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", yourHandler).Methods("POST")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
