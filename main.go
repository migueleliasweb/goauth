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
		} else {
			w.Write([]byte(userPass.User))
		}
	}
}

func setupRedis(RedisSchemeURL *string) (connection redis.Conn) {
	connection, error := redis.DialURL(*RedisSchemeURL)
	if error != nil {
		log.Fatalln("Could not connect to REDIS:", error)
	} else {
		fmt.Println("Connected to REDIS:", *RedisSchemeURL)
	}

	return connection
}

func setupFlags() (RedisSchemeURL *string, HTTPAddress *string) {
	RedisSchemeURL = flag.String("redis-url", "redis://@localhost:6379/0", "Url to connect to redis")
	HTTPAddress = flag.String("http-address", ":8080", "Http address to bind to")

	flag.Parse()

	return RedisSchemeURL, HTTPAddress
}

func main() {
	RedisSchemeURL, HTTPAddress := setupFlags()
	redisConnection := setupRedis(RedisSchemeURL)

	http.HandleFunc("/auth", authHandler)
	log.Fatalln(http.ListenAndServe(*HTTPAddress, nil))

	defer redisConnection.Close()
}
