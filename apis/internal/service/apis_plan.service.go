package service

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/loctodale/go_api_hubs_microservice/apis/database"
	"github.com/loctodale/go_api_hubs_microservice/apis/internal/repository"
	pb "github.com/loctodale/go_api_hubs_microservice/apis/pb/apis"
)

type ApisPlanService interface {
	CreateOneApisPlans(request *pb.CreateApiPlanRequest) error
}

type apisPlanService struct {
	apisPlanRepository repository.ApisPlanRepository
	apisRepository     repository.ApiRepository
}

func NewApisPlanService(apisPlanRepository repository.ApisPlanRepository, apisRepository repository.ApiRepository) ApisPlanService {
	return &apisPlanService{
		apisPlanRepository: apisPlanRepository,
		apisRepository:     apisRepository,
	}
}

func (a *apisPlanService) CreateOneApisPlans(request *pb.CreateApiPlanRequest) error {
	apisId := pgtype.UUID{}
	err := apisId.Scan(request.GetApiId())
	if err != nil {
		return err
	}

	//1. Check api is available
	result, err := a.apisRepository.GetById(apisId)
	if err != nil {
		return err
	}

	if result.Name == "" {
		return err
	}

	var params = database.CreateApisPlanParams{}
	params.Name = request.GetName()
	params.ApiID = apisId
	params.MonthlyLimit = pgtype.Int4{Int32: request.GetMonthlyLimit(), Valid: true}
	params.PricePerCall = pgtype.Int4{Int32: request.GetPricePerCall(), Valid: true}
	params.RateLimit = pgtype.Int4{Int32: request.GetRateLimit(), Valid: true}
	err = a.apisPlanRepository.CreateApisPlan(params)
	if err != nil {
		return err
	}

	return nil
}
