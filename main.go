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

type components struct {
	RedisConnection redis.Conn
	RedisSchemeURL  *string
	HTTPAddress     *string
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	//enforce POST
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}

	bodyBytes, bodyError := ioutil.ReadAll(r.Body)

	if bodyError != nil {
		log.Fatalln(bodyError)
	} else {
		var userPass userPassAuth
		UnmarshalError := json.Unmarshal(bodyBytes, &userPass)

		if UnmarshalError != nil {
			log.Println(UnmarshalError)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid json."))

			return
		}

		redisConnection := components.RedisConnection
		redisUser, redisConnectionError := redisConnection.Do("GET", userPass.User)

		if redisConnectionError != nil {
			log.Println(redisConnectionError)
			w.Write([]byte(redisConnectionError))

			return
		}

		log.Println(redisUser)

		defer redisConnection.Close()

		//w.Write([]byte(userPass.User))
	}
}

func setupRedis() {
	connection, error := redis.DialURL(*RedisSchemeURL)
	if error != nil {
		log.Fatalln("Could not connect to REDIS:", error)
	} else {
		fmt.Println("Connected to REDIS:", *RedisSchemeURL)
	}

	components.RedisConnection = connection
}

func setupFlags(c components) {
	RedisSchemeURL = flag.String("redis-url", "redis://@localhost:6379/0", "Url to connect to redis")
	HTTPAddress = flag.String("http-address", ":8080", "Http address to bind to")

	flag.Parse()

	c.RedisSchemeURL = RedisSchemeURL
	c.HTTPAddress = HTTPAddress
}

func main() {
	globalComponents := new(components)

	setupFlags(*globalComponents)
	setupRedis(*globalComponents)

	http.HandleFunc("/auth", authHandler)
	log.Fatalln(http.ListenAndServe(*HTTPAddress, nil))

	defer redisConnection.Close()
}
