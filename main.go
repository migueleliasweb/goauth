package main

import (
	"components"
	"log"
	"net/http"

	// "github.com/garyburd/redigo/redis"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.POST("/auth", components.AuthHandler{})

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", router))
}
