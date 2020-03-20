package redis

import (
	"MyBlog/configs"
	"errors"
	"fmt"

	"github.com/go-redis/redis"
)

var Client *redis.Client

// InitRedis 初始化Redis
func InitRedis() (err error) {

	ip := configs.GetString("redis.host")
	if ip == "" {
		return errors.New("get_redis_host_error")
	}
	port := configs.GetInt("redis.port")
	if port == 0 {
		return errors.New("get_redis_port_error")
	}
	password := configs.GetString("redis.password")
	if password == "" {
		return errors.New("get_redis_password_error")
	}

	//ip := "127.0.0.1"
	//port := 6379
	//password := "123456"

	Client = redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", ip, port),
		Password:   password,
		DB:         0,
		MaxRetries: 3,
	})

	return nil
}
