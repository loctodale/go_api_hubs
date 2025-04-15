package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/loctodale/go_api_hubs_microservice/account"
	"github.com/loctodale/go_api_hubs_microservice/account/database"
)

type Repository interface {
	Close()
	CreateNewAccount(a database.AddUserBaseParams) error
	GetOneUserInfo(userAccount string) (database.GetOneUserInfoRow, error)
	GetOneUserInfoAdmin(userAccount string) (database.GetOneUserInfoAdminRow, error)
	LoginUserBase(loginParams database.LoginUserBaseParams) error
	CheckUserBaseExists(userAccount string) (int64, error)
}

type postgresRepository struct {
	conn    *pgx.Conn
	queries *database.Queries
}

func NewPostgresRepository(url string) (Repository, error) {
	conn, err := pgx.Connect(account.Ctx, url)
	if err != nil {
		return nil, err
	}

	return &postgresRepository{queries: database.New(conn), conn: conn}, nil
}

func (r postgresRepository) Close() {
	r.conn.Close(account.Ctx)
}

func (r postgresRepository) CreateNewAccount(a database.AddUserBaseParams) error {
	_, err := r.queries.AddUserBase(account.Ctx, a)
	if err != nil {
		return err
	}
	return nil
}

func (r postgresRepository) GetOneUserInfo(userAccount string) (database.GetOneUserInfoRow, error) {
	result, err := r.queries.GetOneUserInfo(account.Ctx, userAccount)
	if err != nil {
		return database.GetOneUserInfoRow{}, err
	}
	return result, nil

}

func (r postgresRepository) GetOneUserInfoAdmin(userAccount string) (database.GetOneUserInfoAdminRow, error) {
	result, err := r.queries.GetOneUserInfoAdmin(account.Ctx, userAccount)
	if err != nil {
		return database.GetOneUserInfoAdminRow{}, err
	}
	return result, nil
}

func (r postgresRepository) LoginUserBase(loginParams database.LoginUserBaseParams) error {
	err := r.queries.LoginUserBase(account.Ctx, loginParams)
	if err != nil {
		return err
	}
	return nil
}

func (r postgresRepository) CheckUserBaseExists(userAccount string) (int64, error) {
	result, err := r.queries.CheckUserBaseExists(account.Ctx, userAccount)
	if err != nil {
		return 0, err
	}
	return result, nil
}
