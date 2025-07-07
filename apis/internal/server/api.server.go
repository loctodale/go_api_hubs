package server

import (
	"context"
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/apis/internal/service"
	pb "github.com/loctodale/go_api_hubs_microservice/apis/pb/apis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type grpcServer struct {
	pb.UnimplementedApiServiceServer
	pb.UnimplementedApisPlanServiceServer
	pb.UnimplementedApisKeyServiceServer
	apisService     service.ApisService
	apisPlanService service.ApisPlanService
	apisKeyService  service.ApisKeyService
}

func ListenGRPC(apisService service.ApisService, apisPlanService service.ApisPlanService, apisKeyService service.ApisKeyService, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterApisPlanServiceServer(serv, &grpcServer{
		apisPlanService: apisPlanService,
		apisKeyService:  apisKeyService,
		apisService:     apisService,
	})
	pb.RegisterApiServiceServer(serv, &grpcServer{
		apisPlanService: apisPlanService,
		apisKeyService:  apisKeyService,
		apisService:     apisService,
	})
	pb.RegisterApisKeyServiceServer(serv, &grpcServer{
		apisPlanService: apisPlanService,
		apisKeyService:  apisKeyService,
		apisService:     apisService,
	})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) AddOneApi(ctx context.Context, req *pb.CreateApiRequest) (*pb.BaseResponse, error) {
	_, err := s.apisService.CreateOneApi(req)
	if err != nil {
		fmt.Println("Error::", err)
		return &pb.BaseResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}
	return &pb.BaseResponse{
		Message: "Tạo mới thành công",
		Code:    200,
	}, nil
}
func (s *grpcServer) GetApiById(ctx context.Context, req *pb.GetOneApiRequest) (*pb.GetOneApiResponse, error) {
	result, err := s.apisService.GetById(req.GetId())
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (s *grpcServer) AddOneApiPlan(ctx context.Context, req *pb.CreateApiPlanRequest) (*pb.BaseResponse, error) {
	fmt.Println("AddOneApiPlan")
	err := s.apisPlanService.CreateOneApisPlans(req)
	if err != nil {
		fmt.Println("Error::", err)
		return &pb.BaseResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}
	return &pb.BaseResponse{
		Code:    200,
		Message: "Tạo api plan thành công",
	}, nil
}
func (s *grpcServer) AddOneApisKey(ctx context.Context, req *pb.CreateApisKeyRequest) (*pb.BaseResponse, error) {
	err := s.apisKeyService.CreateApisKey(req)
	if err != nil {
		fmt.Println("Error::", err)
		return &pb.BaseResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}
	return &pb.BaseResponse{
		Code:    200,
		Message: "Tao thành công",
	}, nil
}

func (s *grpcServer) CallApiByKey(ctx context.Context, req *pb.CallApiRequest) (*pb.CallApiResponse, error) {
	result, err := s.apisKeyService.CallApiByKey(req)
	if err != nil {
		fmt.Println("Error:: ", err)
		return nil, err
	}

	return result, nil
}
