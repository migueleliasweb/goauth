package main

import (
	"components"
	"log"
	"net/http"

	// "github.com/garyburd/redigo/redis"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//DB := components.SetupTestDB()

	router := httprouter.New()

	router.POST("/auth", components.GoAuthMiddleWare(components.AuthHandler))

	log.Fatal(http.ListenAndServe(":8000", router))
}
