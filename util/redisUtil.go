package util

import (
	"fmt"

	redis "github.com/go-redis/redis"
	ini "gopkg.in/ini.v1"
)

//GetRedisConnection get the redis connection with the provided configuration
func GetRedisConnection(config *ini.File) *redis.Client {

	section := config.Section("redis")

	number, err := section.Key("DB").Int()

	if err != nil {
		fmt.Println(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     section.Key("Address").String(),
		Password: section.Key("Password").String(),
		DB:       number,
	})

	return rdb
}
