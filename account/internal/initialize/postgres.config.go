package initialize

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/loctodale/go_api_hubs_microservice/account/global"
)

func InitPostgresServer() {
	connectionString := global.Config.AccountService.Database.Postgres
	conn, err := pgx.Connect(global.Ctx, connectionString)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to postgres")
	global.Pdb = conn
}
