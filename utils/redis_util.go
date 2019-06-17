package utils

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

type RedisUtils struct {
	conn redis.Conn
}

func ErrCheck(err error) {
	if err != nil {
		log.Fatalln("sorry, has some error:", err)
	}
}

func (r *RedisUtils) ConnectByUrl(url string) {
	c, err := redis.DialURL(url)
	ErrCheck(err)
	r.conn = c
}

func (r *RedisUtils) Close() {
	r.conn.Close()
}

func (r *RedisUtils) Keys() ([]string, error) {
	json, getErr := redis.Strings(r.conn.Do("keys", "*"))
	return json, getErr
}

func (r *RedisUtils) Get(key string) (string, error) {
	json, getErr := redis.String(r.conn.Do("get", key))
	return json, getErr
}
