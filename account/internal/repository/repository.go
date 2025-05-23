package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/loctodale/go_api_hubs_microservice/account/database"
	"github.com/loctodale/go_api_hubs_microservice/account/global"
)

type Repository interface {
	Close()
	CreateNewAccount(a database.AddUserBaseParams) error
	GetOneUserInfo(userAccount string) (database.GetOneUserInfoRow, error)
	GetOneUserInfoAdmin(userAccount string) (database.GetOneUserInfoAdminRow, error)
	LoginUserBase(loginParams database.LoginUserBaseParams) error
	CheckUserBaseExists(userAccount string) (int64, error)
	GetAccounts() []database.GetAccountsRow
}

type postgresRepository struct {
	conn    *pgx.Conn
	queries *database.Queries
}

//func NewPostgresRepository(url string) (Repository, error) {
//	conn, err := pgx.Connect(global.Ctx, url)
//	if err != nil {
//		return nil, err
//	}
//
//	return &postgresRepository{queries: database.New(conn), conn: conn}, nil
//}

func NewAccountRepository() (Repository, error) {
	return &postgresRepository{
		conn:    global.Pdb,
		queries: database.New(global.Pdb),
	}, nil
}

func (r postgresRepository) Close() {
	r.conn.Close(global.Ctx)
}

func (r postgresRepository) CreateNewAccount(a database.AddUserBaseParams) error {
	_, err := r.queries.AddUserBase(global.Ctx, a)
	if err != nil {
		return err
	}
	return nil
}

func (r postgresRepository) GetOneUserInfo(userAccount string) (database.GetOneUserInfoRow, error) {
	result, err := r.queries.GetOneUserInfo(global.Ctx, userAccount)
	if err != nil {
		return database.GetOneUserInfoRow{}, err
	}
	return result, nil

}

func (r postgresRepository) GetOneUserInfoAdmin(userAccount string) (database.GetOneUserInfoAdminRow, error) {
	result, err := r.queries.GetOneUserInfoAdmin(global.Ctx, userAccount)
	if err != nil {
		return database.GetOneUserInfoAdminRow{}, err
	}
	return result, nil
}

func (r postgresRepository) LoginUserBase(loginParams database.LoginUserBaseParams) error {
	err := r.queries.LoginUserBase(global.Ctx, loginParams)
	if err != nil {
		return err
	}
	return nil
}

func (r postgresRepository) CheckUserBaseExists(userAccount string) (int64, error) {
	result, err := r.queries.CheckUserBaseExists(global.Ctx, userAccount)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (r postgresRepository) GetAccounts() []database.GetAccountsRow {
	result, err := r.queries.GetAccounts(global.Ctx)
	if err != nil {
		return []database.GetAccountsRow{}
	}

	return result
}
