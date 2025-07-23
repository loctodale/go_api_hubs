package logs

import (
	"encoding/json"
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/apis/global"
	discordPb "github.com/loctodale/go_api_hubs_microservice/apis/pb/discord_logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"
)

type logService struct {
}

func LogService() *logService {
	return &logService{}
}

func (l *logService) LogApisService(req any, method string, path string, code int32) {
	conn, err := grpc.Dial("discord-logs-bot-service:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonData, err := json.Marshal(req)
	if err != nil {
		return
	}

	var rawMap map[string]interface{}
	err = json.Unmarshal(jsonData, &rawMap)
	if err != nil {
		return
	}

	messageStruct, err := structpb.NewStruct(rawMap)
	if err != nil {
	}
	defer conn.Close()
	discordClient := discordPb.NewDiscordLogsServiceClient(conn)
	_, err = discordClient.SendLogsTracking(global.Ctx, &discordPb.SendLogsTrackingRequest{
		Code:    code,
		Message: messageStruct,
		Service: "account",
		Method:  method,
		Path:    path,
	})
	if err != nil {
		fmt.Println(err)
	}
}
