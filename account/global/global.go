package global

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/loctodale/go_api_hubs_microservice/pkg"
	"github.com/redis/go-redis/v9"
)

var (
	Config *pkg.Config
	Rdb    *redis.Client
	Pdb    *pgx.Conn
	Ctx    = context.Background()
)
