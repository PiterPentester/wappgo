package models

import (
	"wappgo/utils"

	"github.com/go-redis/redis"
)

var client *redis.Client

func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:     utils.Host + ":" + utils.Port,
		Password: utils.Password,
		DB:       0,
	})
}
