package initialize

import (
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/account/global"
	"github.com/redis/go-redis/v9"
)

func InitRedisServer() {
	r := global.Config.AccountService
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Database.Redis.Host, r.Database.Redis.Port),
		Password: r.Database.Redis.Password, // no password set
		DB:       r.Database.Redis.Database, // use default DB
		PoolSize: r.Database.Redis.Poll,
	})
	if rdb != nil {
		fmt.Println("Redis initialize success")
	}
	global.Rdb = rdb
}
