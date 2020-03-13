package redisClient

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
)


func NewPool(maxidle int, maxactive int) *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: maxidle,
		// max number of connections
		MaxActive: maxactive,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(os.Getenv("REDIS_URL"))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func RedisClientSet(c redis.Conn, key string, value string) string {
	_, err := c.Do("SET", key, value)
	if err != nil {
		log.Fatal(err)
	}
	return "Value of "+key+" has been stored."
}

func RedisClientGet(c redis.Conn,key string) string{
	values, err := redis.String(c.Do("GET", key))
	if err == redis.ErrNil {
		return "Key "+key+" does not exist !!! .\nSET first with /set command !!!"
	} else if err != nil {
		log.Fatal(err)
	}
	return values
}
func RedisClientFlush(c redis.Conn) string {
	err := c.Send("FLUSHALL")
	if err != nil {
		log.Fatal(err)
	}
	return "All key values has been deleted."
}

func RedisClientDelete(c redis.Conn,key string) string{
	values, err := redis.String(c.Do("DEL", key))
	if err == redis.ErrNil {
		return "Key "+key+" does not exist !!! .\nSET first with /set command !!!"
	} else if err != nil {
		log.Fatal(err)
	}
	return values
}
