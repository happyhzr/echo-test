package redis

import (
	"fmt"
	"time"

	"github.com/insisthzr/echo-test/cookbook/twitter/conf"

	"github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool
)

func init() {
	pool = newPool(conf.REDIS_HOST)
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func Ping() error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("PING")
	return err
}

func Get(key string) ([]byte, error) {
	conn := pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}

func Set(key string, value []byte) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value, "ex", conf.REDIS_EX)
	return err
}
