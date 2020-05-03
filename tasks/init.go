package tasks

import (
	"os"
	"strconv"

	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
)

func intialize() *gocelery.CeleryClient {
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(os.Getenv("BROKER_URL"))
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}

	// initialize celery client
	workers, _ := strconv.Atoi(os.Getenv("WORKER_POOL"))
	cli, _ := gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&gocelery.RedisCeleryBackend{Pool: redisPool},
		workers,
	)
	return cli

}
