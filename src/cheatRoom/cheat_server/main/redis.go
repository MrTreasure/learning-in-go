package main

import "github.com/garyburd/redigo/redis"
import "time"

var pool *redis.pool

func initRedis(addr string, idelConn, maxConn int, idelTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle: idelConn,
		MaxActive: maxConn,
		IdleTimeout: idelTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		}
	}
	return
}

func GetConn() redis.Conn {
	return pool.Get()
}

fun putConn(conn redis.Conn) {
	conn.close()
}