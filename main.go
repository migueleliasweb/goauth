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

	router.POST("/auth", components.GoAuthMiddleWare(components.AuthHandler))
	router.GET("/permission", components.GoAuthMiddleWare(components.GetPermissionHandler))

	log.Fatal(http.ListenAndServe(":8000", router))
}
