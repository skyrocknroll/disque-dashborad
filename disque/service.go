package disque

import (
	"github.com/garyburd/redigo/redis"
	"github.com/EverythingMe/go-disque/disque"
)

func dial(addr string) (redis.Conn, error) {
	return redis.Dial("tcp", addr)
}

func GetPool() *disque.Pool {

	pool := disque.NewPool(disque.DialFunc(dial), "127.0.0.1:7711")

	return pool
}
