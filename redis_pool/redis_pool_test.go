package redis_pool

import (
	"fmt"
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
)

func TestRedisPool(t *testing.T) {
	InitRedisConfig(&Conf{
		Host: "127.0.0.1",
		Port: 6379,
	}, redis.DialConnectTimeout(10*time.Second))

	fmt.Println(GetRedisInstance().Ping())

}
