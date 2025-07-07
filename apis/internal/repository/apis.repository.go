package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/loctodale/go_api_hubs_microservice/apis/database"
	"github.com/loctodale/go_api_hubs_microservice/apis/global"
)

type ApiRepository interface {
	CreateOneApi(params database.CreateApisParams) (pgconn.CommandTag, error)
	GetById(id pgtype.UUID) (database.TblApi, error)
}

type apiRepository struct {
	conn    *pgx.Conn
	queries *database.Queries
}

func NewApiRepository() ApiRepository {
	return &apiRepository{
		conn:    global.Pdb,
		queries: database.New(global.Pdb),
	}
}

func (a apiRepository) CreateOneApi(params database.CreateApisParams) (pgconn.CommandTag, error) {
	result, err := a.queries.CreateApis(context.Background(), params)
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return result, nil
}

func (a apiRepository) GetById(id pgtype.UUID) (database.TblApi, error) {

	result, err := a.queries.GetApiById(global.Ctx, id)
	if err != nil {
		return database.TblApi{}, err
	}
	return result, nil
}
