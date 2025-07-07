package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/loctodale/go_api_hubs_microservice/apis/database"
	"github.com/loctodale/go_api_hubs_microservice/apis/internal/repository"
	pb "github.com/loctodale/go_api_hubs_microservice/apis/pb/apis"
	"github.com/loctodale/go_api_hubs_microservice/apis/utils"
	"google.golang.org/protobuf/types/known/structpb"
	"io"
	"net/http"
)

type ApisKeyService interface {
	CreateApisKey(*pb.CreateApisKeyRequest) error
	CallApiByKey(request *pb.CallApiRequest) (*pb.CallApiResponse, error)
}

type apisKeyService struct {
	apisKeyRepository repository.ApisKeyRepository
	apiRepository     repository.ApiRepository
}

func NewApisKeyService(apisKeyRepository repository.ApisKeyRepository, apisRepository repository.ApiRepository) ApisKeyService {
	return &apisKeyService{
		apisKeyRepository: apisKeyRepository,
		apiRepository:     apisRepository,
	}
}

func (a apisKeyService) CreateApisKey(req *pb.CreateApisKeyRequest) error {
	var apiId pgtype.UUID
	err := apiId.Scan(req.GetApiId())
	if err != nil {
		return err
	}

	var planId pgtype.UUID
	err = planId.Scan(req.GetPlanId())
	if err != nil {
		return err
	}

	var userId pgtype.UUID
	err = userId.Scan(req.GetUserId())
	if err != nil {
		return err
	}
	var quotaReset pgtype.Timestamp
	err = quotaReset.Scan(req.QuotaResetAt)
	params := database.CreateApisKeyParams{
		ApiID:  apiId,
		PlanID: planId,
		UserID: userId,
		ApiKey: pgtype.Text{
			Valid:  true,
			String: utils.NewApiKeyRandom().GenerateRandomApiKey(16),
		},
		IsActive: pgtype.Bool{
			Valid: true,
			Bool:  req.GetIsActive(),
		},
		QuotaResetAt: quotaReset,
		QuotaUsed: pgtype.Int4{
			Valid: true,
			Int32: req.QuotaUsed,
		},
	}
	err = a.apisKeyRepository.CreateApisKey(params)

	if err != nil {
		return err
	}

	return nil
}

func (a apisKeyService) CallApiByKey(request *pb.CallApiRequest) (*pb.CallApiResponse, error) {

	//1. Lấy thông tin api key
	result, err := a.apisKeyRepository.GetByKey(pgtype.Text{
		Valid:  true,
		String: request.GetApiKey(),
	})
	if err != nil {
		return &pb.CallApiResponse{
			BaseResponse: &pb.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
			Payload: nil,
		}, err
	}

	//2. Lấy thông tin api

	api, err := a.apiRepository.GetById(result.ApiID)
	if err != nil {
		return nil, err
	}

	//3. Map thông tin body trong request

	jsonData, err := json.Marshal(request.Body.AsMap())

	fmt.Print(jsonData)
	if err != nil {
		return &pb.CallApiResponse{
			BaseResponse: &pb.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
			Payload: nil,
		}, err
	}

	//4. Thực hiện request

	req, err := http.NewRequest(request.Method, api.BaseUrl.String, bytes.NewBuffer(jsonData))
	if err != nil {
		return &pb.CallApiResponse{
			BaseResponse: &pb.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
			Payload: nil,
		}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	// Map response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return &pb.CallApiResponse{
			BaseResponse: &pb.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
			Payload: nil,
		}, err
	}

	var jsonMap map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &jsonMap); err != nil {
		return &pb.CallApiResponse{
			BaseResponse: &pb.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
			Payload: nil,
		}, err
	}

	respPayload, err := structpb.NewStruct(jsonMap)

	if err != nil {
		return &pb.CallApiResponse{
			BaseResponse: &pb.BaseResponse{
				Code:    500,
				Message: err.Error(),
			},
			Payload: nil,
		}, err
	}

	return &pb.CallApiResponse{
		BaseResponse: &pb.BaseResponse{
			Code:    200,
			Message: "Success",
		},
		Payload: respPayload,
	}, nil

}
