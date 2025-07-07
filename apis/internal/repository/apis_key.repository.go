package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/loctodale/go_api_hubs_microservice/apis/database"
	"github.com/loctodale/go_api_hubs_microservice/apis/global"
)

type ApisKeyRepository interface {
	CreateApisKey(params database.CreateApisKeyParams) error
	GetByKey(key pgtype.Text) (database.TblApisKey, error)
}

type apisKeyRepository struct {
	conn  *pgx.Conn
	query *database.Queries
}

func NewApisKeyRepository() ApisKeyRepository {
	return &apisKeyRepository{
		conn:  global.Pdb,
		query: database.New(global.Pdb),
	}
}

func (r *apisKeyRepository) CreateApisKey(params database.CreateApisKeyParams) error {
	_, err := r.query.CreateApisKey(global.Ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (r *apisKeyRepository) GetByKey(key pgtype.Text) (database.TblApisKey, error) {
	result, err := r.query.GetApisKeyByKey(global.Ctx, key)
	if err != nil {
		return database.TblApisKey{}, err
	}

	return result, nil
}
