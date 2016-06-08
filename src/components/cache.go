package components

import (
	"bytes"
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

const userPreffix = "user_"

var (
	//Cache Redis instance
	Cache CacheWrapper
)

//CacheWrapper "De facto" where to search for information
type CacheWrapper struct {
	redisConn redis.Conn
	redisErr  error
}

//usernameKeyHelper Returns the cache key for the given user
func usernameKeyHelper(username *string) string {
	var buffer bytes.Buffer

	buffer.WriteString(userPreffix)
	buffer.WriteString(*username)

	return buffer.String()
}

//UserExists Returns true if a user with the given name exists
func (CW *CacheWrapper) UserExists(username *string) bool {
	userExists, err := redis.Bool(CW.redisConn.Do("EXISTS", usernameKeyHelper(username)))

	if err != nil {
		log.Println("Could not check user" + *username + " for existance.")
		log.Fatalln(err)
	}

	return userExists
}

//GetEncodedPassword Returns the encoded password for the given username
func (CW *CacheWrapper) GetEncodedPassword(username *string) *string {
	password, err := redis.String(CW.redisConn.Do("HGET", usernameKeyHelper(username), "password"))

	if err != nil {
		log.Println("Could not fetch user's password")
		log.Fatalln(err)
	}

	return &password
}

func init() {
	Cache := new(CacheWrapper)
	Cache.redisConn, Cache.redisErr = redis.DialURL(os.Getenv("REDIS_URL"))

	if Cache.redisErr != nil {
		log.Fatalln("Could not connect to redis.")
	}

}
