package repository

import (
	"github.com/loctodale/go_api_hubs_microservice/account/global"
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	GetByUserKey(userKey string) (string, error)
}

type repository struct {
	rdb *redis.Client
}

func NewRepository(rdb *redis.Client) Repository {
	return &repository{
		rdb: global.Rdb,
	}
}

func (r *repository) GetByUserKey(userKey string) (string, error) {
	key, error := r.rdb.Get(global.Ctx, userKey).Result()
	if error != nil {
		return "", error
	}
	return key, nil
}
