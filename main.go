package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/garyburd/redigo/redis"
)

func fooHandler(w http.ResponseWriter, r *http.Request) {
	os.Exit(1)
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	w.Write([]byte(r.URL.Path))
	log.Panicln("BLURP")
}

func main() {
	RedisSchemeURL := flag.String("redis-url", "redis://@localhost:6379/0", "Url to connect to redis")
	HTTPAddress := flag.String("http-address", ":8080", "Http address to bind to")

	flag.Parse()

	c, err := redis.DialURL(*RedisSchemeURL)
	if err != nil {
		log.Fatalln("Could not connect to REDIS: ", err)
	} else {
		fmt.Println("Connected to REDIS: ", *RedisSchemeURL)
	}

	httpServer := http.NewServeMux()
	httpServer.HandleFunc("/foo/", fooHandler)
	log.Fatalln(http.ListenAndServe(*HTTPAddress, nil))

	defer c.Close()
}
