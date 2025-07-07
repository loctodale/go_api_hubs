package service

import (
	"errors"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/loctodale/go_api_hubs_microservice/apis/database"
	"github.com/loctodale/go_api_hubs_microservice/apis/global"
	"github.com/loctodale/go_api_hubs_microservice/apis/internal/repository"
	aPb "github.com/loctodale/go_api_hubs_microservice/apis/pb/account"
	pb "github.com/loctodale/go_api_hubs_microservice/apis/pb/apis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ApisService interface {
	CreateOneApi(params *pb.CreateApiRequest) (pgconn.CommandTag, error)
	GetById(id string) (pb.GetOneApiResponse, error)
}

type apisService struct {
	apiRepository     repository.ApiRepository
	apiPlanRepository repository.ApisPlanRepository
	apiKeyRepository  repository.ApisKeyRepository
}

func (a apisService) GetById(id string) (pb.GetOneApiResponse, error) {
	var uuid pgtype.UUID
	err := uuid.Scan(id)
	if err != nil {
		return pb.GetOneApiResponse{
			BaseResponse: &pb.BaseResponse{
				Message: err.Error(),
				Code:    400,
			},
		}, err
	}
	result, err := a.apiRepository.GetById(uuid)
	if err != nil {
		return pb.GetOneApiResponse{
			BaseResponse: &pb.BaseResponse{
				Message: err.Error(),
				Code:    400,
			},
		}, err
	}
	return pb.GetOneApiResponse{
		BaseResponse: &pb.BaseResponse{
			Message: "Get success",
			Code:    200,
		},
		ApiModel: &pb.ApiModel{
			Name:        result.Name,
			Slug:        result.Slug,
			Category:    result.Category.String,
			BaseUrl:     result.BaseUrl.String,
			Description: result.Description.String,
			Status:      string(result.Status.ApiStatus),
			CreatedAt:   result.CreatedAt.Time.String(),
			DocsUrl:     result.DocUrl.String,
			ProviderId:  result.ProviderID.String(),
		},
	}, nil
}

func (a apisService) CreateOneApi(req *pb.CreateApiRequest) (pgconn.CommandTag, error) {
	conn, err := grpc.Dial("account-service:6000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return pgconn.CommandTag{}, err
	}
	defer conn.Close()

	accountClient := aPb.NewPrivateAccountServiceClient(conn)
	//Check provider is existed
	accountFound, err := accountClient.GetAccount(global.Ctx, &aPb.GetAccountByIdRequest{Id: req.ProviderId})
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	if accountFound == nil {
		return pgconn.CommandTag{}, errors.New("provider not found")
	}
	//fmt.Print(acoountClient.RegisterAccount())
	uuid := pgtype.UUID{}
	err = uuid.Scan(req.GetProviderId())
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	result, err := a.apiRepository.CreateOneApi(database.CreateApisParams{
		Name: req.GetName(),
		Status: database.NullApiStatus{
			ApiStatus: database.ApiStatus(req.GetStatus()),
			Valid:     true,
		},
		BaseUrl: pgtype.Text{
			String: req.GetBaseUrl(),
			Valid:  true,
		},
		Category: pgtype.Text{
			String: req.GetCategory(),
			Valid:  true,
		},
		Description: pgtype.Text{
			String: req.GetDescription(),
			Valid:  true,
		},
		DocUrl: pgtype.Text{
			String: req.GetDocsUrl(),
			Valid:  true,
		},
		Slug:       slug.Make(req.GetName()),
		ProviderID: uuid,
	})
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	return result, nil
}

func NewApiService(
	apiRepository repository.ApiRepository,
	apiPlanRepository repository.ApisPlanRepository,
	apiKeyRepository repository.ApisKeyRepository,
) ApisService {
	return &apisService{
		apiRepository:     apiRepository,
		apiPlanRepository: apiPlanRepository,
		apiKeyRepository:  apiKeyRepository,
	}
}
