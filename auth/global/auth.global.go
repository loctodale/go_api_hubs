package global

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/loctodale/go_api_hubs_microservice/pkg"
)

var (
	Config *pkg.Config
	Pdb    *pgx.Conn
	Ctx    = context.Background()
)
