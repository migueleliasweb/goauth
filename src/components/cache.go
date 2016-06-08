package components

import (
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

const userPreffix = "user_"

var (
	//RedisWrap Redis instance
	RedisWrap RedisWrapper
)

//RedisWrapper "De facto" where to search for information
type RedisWrapper struct {
	redisConn redis.Conn
	redisErr  error
}

//UserExists Returns true if a user with the given name exists
func (RWrap *RedisWrapper) UserExists(username string) bool {
	userExists, err := redis.Bool(RWrap.redisConn.Do("EXISTS", userPreffix+username))

	if err != nil {
		log.Println("Could not check user" + username + " for existance.")
		log.Fatalln(err)
	}

	return userExists
}

func init() {
	RedisWrap := new(RedisWrapper)
	RedisWrap.redisConn, RedisWrap.redisErr = redis.DialURL(os.Getenv("REDIS_URL"))

	if RedisWrap.redisErr != nil {
		log.Fatalln("Could not connect to redis.")
	}

}
