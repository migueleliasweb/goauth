package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

type userPassAuth struct {
	User string
	Pass string
}

type _appContext struct {
	redisConnection *redis.Conn
}

var appContext _appContext

func authHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))

		return
	}

	bodyBytes, bodyError := ioutil.ReadAll(request.Body)

	if bodyError != nil {
		log.Fatalln(bodyError)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(http.StatusText(http.StatusBadRequest)))

		return
	}

	var userPass userPassAuth
	unmarshalError := json.Unmarshal(bodyBytes, &userPass)

	if unmarshalError != nil {
		log.Fatalln(bodyError)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(http.StatusText(http.StatusBadRequest)))

		return
	}

	userHashMap, redisError := redis.StringMap((*appContext.redisConnection).Do("HGETALL", userPass.User))

	if redisError != nil {
		log.Fatalln(bodyError)
		response.WriteHeader(http.StatusForbidden)
		response.Write([]byte(http.StatusText(http.StatusForbidden)))

		return
	}

	if userHashMap["password"] == userPass.Pass {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("You shall pass!"))
	} else {
		log.Fatalln(bodyError)
		response.WriteHeader(http.StatusForbidden)
		response.Write([]byte(http.StatusText(http.StatusForbidden)))

		return
	}
}

func setupRedis(RedisSchemeURL *string) {
	connection, error := redis.DialURL(*RedisSchemeURL)
	if error != nil {
		log.Fatalln("Could not connect to REDIS:", error)
	} else {
		fmt.Println("Connected to REDIS:", *RedisSchemeURL)
	}

	appContext.redisConnection = &connection
}

func setupFlags() (RedisSchemeURL *string, HTTPAddress *string) {
	RedisSchemeURL = flag.String("redis-url", "redis://@localhost:6379/0", "Url to connect to redis")
	HTTPAddress = flag.String("http-address", ":8080", "Http address to bind to")

	flag.Parse()

	return RedisSchemeURL, HTTPAddress
}

func main() {
	RedisSchemeURL, HTTPAddress := setupFlags()

	setupRedis(RedisSchemeURL)

	http.HandleFunc("/auth", authHandler)
	log.Fatalln(http.ListenAndServe(*HTTPAddress, nil))
}
