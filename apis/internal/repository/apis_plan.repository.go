package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/loctodale/go_api_hubs_microservice/apis/database"
	"github.com/loctodale/go_api_hubs_microservice/apis/global"
)

type ApisPlanRepository interface {
	CreateApisPlan(params database.CreateApisPlanParams) error
	GetById(id pgtype.UUID) (database.TblApisPlan, error)
}

type apiPlanRepository struct {
	conn    *pgx.Conn
	queries *database.Queries
}

func NewApiPlanRepository() ApisPlanRepository {
	return &apiPlanRepository{
		conn:    global.Pdb,
		queries: database.New(global.Pdb),
	}
}

func (a apiPlanRepository) CreateApisPlan(params database.CreateApisPlanParams) error {
	_, err := a.queries.CreateApisPlan(global.Ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (a apiPlanRepository) GetById(uuid pgtype.UUID) (database.TblApisPlan, error) {
	result, err := a.queries.GetApisPlanById(global.Ctx, uuid)
	if err != nil {
		return database.TblApisPlan{}, err
	}

	return result, nil
}
