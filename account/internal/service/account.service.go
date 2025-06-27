package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/account/database"
	"github.com/loctodale/go_api_hubs_microservice/account/global"
	repository2 "github.com/loctodale/go_api_hubs_microservice/account/internal/repository"
	"github.com/loctodale/go_api_hubs_microservice/account/pb"
	"github.com/loctodale/go_api_hubs_microservice/account/utils"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"strings"
	"time"
)

type Service interface {
	PostAccount(userAccount string, userPassword string) error
	GetAccounts() ([]database.GetAccountsRow, error)
	RegisterAccount(userAccount string) (pb.BaseResponseMessage, error)
	LoginAccount(username string, password string) (pb.LoginResponse, error)
	VerifyAccount(account string, otp string) (pb.BaseResponseMessage, error)
}
type accountService struct {
	repository      repository2.Repository
	utils           utils.Utils
	tokenRepository repository2.TokenRepository
}

func NewAccountService(repository repository2.Repository, utils utils.Utils, tokenRepository repository2.TokenRepository) Service {
	return &accountService{
		repository,
		utils,
		tokenRepository,
	}
}

func (a *accountService) PostAccount(userAccount string, userPassword string) error {
	//1. Check account is existed
	isExisted, err := a.repository.CheckUserBaseExists(userAccount)
	if err != nil {
		return err
	}

	if isExisted != 0 {
		return errors.New("Account already existed")
	}
	userSalt, err := a.utils.GenerateSalt(16)
	if err != nil {
		return err
	}

	saltPassword := a.utils.HashPasswordSalt(userPassword, userSalt)

	//2. Generate account salt
	accountBaseModel := database.AddUserBaseParams{
		userAccount,
		saltPassword,
		userSalt,
		1,
	}
	if _, err = a.repository.CreateNewAccount(accountBaseModel); err != nil {
		return err
	}
	return nil
}

func (a *accountService) GetAccounts() ([]database.GetAccountsRow, error) {
	result := a.repository.GetAccounts()

	return result, nil
}

func (a *accountService) RegisterAccount(userAccount string) (pb.BaseResponseMessage, error) {
	isExisted, err := a.repository.CheckUserBaseExists(userAccount)
	if err != nil {
		return pb.BaseResponseMessage{
			Code:    400,
			Message: err.Error(),
		}, err
	}
	if isExisted != 0 {
		return pb.BaseResponseMessage{
			Code:    400,
			Message: "Account already existed",
		}, errors.New("Account already existed")
	}

	hashKey := a.utils.GetHash(strings.ToLower(userAccount))
	userKey := fmt.Sprintf("u:%s:otp", hashKey)
	fmt.Println("Hash key: ", hashKey)
	fmt.Println("UserKey: ", userKey)

	otpFound, err := global.Rdb.Get(global.Ctx, userKey).Result()
	fmt.Println("otpFound: ", otpFound)
	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist")
		break
	case err != nil:
		fmt.Println("Get failed: ", err)
		return pb.BaseResponseMessage{Code: 400, Message: err.Error()}, err
	case otpFound != "":
		return pb.BaseResponseMessage{Code: 400, Message: "otp empty"}, errors.New("otp empty")
	}
	otp := a.utils.GenerateSixDigitOtp()
	fmt.Println("otp: ", otp)
	messageBody := make(map[string]interface{})
	messageBody["otp"] = otp
	messageBody["email"] = userAccount

	kafkaBody, err := json.Marshal(messageBody)

	if err != nil {
		fmt.Println(err)
		return pb.BaseResponseMessage{Code: 400, Message: err.Error()}, err
	}
	kafkaMessage := kafka.Message{
		Key:   []byte(userKey),
		Value: kafkaBody,
	}
	done := make(chan error)
	go func() {
		err = global.Rdb.Set(global.Ctx, userKey, otp, time.Duration(2)*time.Minute).Err()
		if err != nil {
			done <- err
		}

		fmt.Println("Kafka message sent", kafkaBody)
		err = global.KafkaProducer.WriteMessages(global.Ctx, kafkaMessage)

		if err != nil {
			fmt.Println("kafka error: ", err)
			done <- err
		}
		done <- nil
	}()

	if err = <-done; err != nil {
		fmt.Println(err)
		return pb.BaseResponseMessage{Code: 400, Message: err.Error()}, err
	}
	return pb.BaseResponseMessage{
		Code:    200,
		Message: "Đăng ký tài khoản thành công",
	}, nil
}

func (a *accountService) VerifyAccount(account string, otp string) (pb.BaseResponseMessage, error) {
	hashKey := a.utils.GetHash(strings.ToLower(account))
	userKey := fmt.Sprintf("u:%s:otp", hashKey)
	foundOTP, err := global.Rdb.Get(global.Ctx, userKey).Result()
	if err != nil {
		return pb.BaseResponseMessage{Code: 400, Message: err.Error()}, err
	}

	if otp != foundOTP {
		return pb.BaseResponseMessage{Code: 400, Message: "Mã OTP không chính xác, vui lòng thử lại"}, errors.New("Mã OTP không chính xác, vui lòng thử lại")
	}

	saltKey, err := a.utils.GenerateSalt(16)
	if err != nil {
		return pb.BaseResponseMessage{Code: 400, Message: err.Error()}, err
	}

	saltPassword := a.utils.HashPasswordSalt(otp, saltKey)

	resp, err := a.repository.CreateNewAccount(database.AddUserBaseParams{
		UserSalt:     saltKey,
		UserPassword: saltPassword,
		UserAccount:  account,
		UserRole:     1,
	})
	if err != nil {
		return pb.BaseResponseMessage{Code: 400, Message: err.Error()}, err
	}

	// Create consumer (Kong)
	err = a.utils.CreateConsumer(account, resp.String())
	if err != nil {
		return pb.BaseResponseMessage{Code: 400, Message: err.Error()}, err
	}

	// Create ACL Group for Consumer (Kong)
	err = a.utils.AddConsumerToACLGroup(account, "member")
	if err != nil {
		return pb.BaseResponseMessage{Code: 400, Message: err.Error()}, err
	}
	return pb.BaseResponseMessage{
		Code:    200,
		Message: "Verify account success",
	}, nil
}

func (a *accountService) LoginAccount(username string, password string) (pb.LoginResponse, error) {
	foundAccount, err := a.repository.GetLoginAccount(username)
	if err != nil {
		return pb.LoginResponse{
			AccessToken:  "",
			RefreshToken: "nil",
		}, err
	}

	matchPassword := a.utils.MatchPassword(foundAccount.UserPassword, password, foundAccount.UserSalt)
	if !matchPassword {
		return pb.LoginResponse{}, errors.New("Invalid username or password")
	}

	privateKey, publicKey, err := a.utils.GenerateKey()
	if err != nil {
		return pb.LoginResponse{
			RefreshToken: "",
			AccessToken:  "",
		}, err
	}
	accessToken, refreshToken, err := a.utils.GenerateToken(username, privateKey, foundAccount.UserID.String())
	if err != nil {
		return pb.LoginResponse{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}

	pubPEMString, err := a.utils.PublicKeyToString(publicKey)
	if err != nil {
		fmt.Println(err)
		return pb.LoginResponse{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}
	// save public key, refresh token into database
	err = a.tokenRepository.AddToken(database.CreateNewTokenParams{
		RefreshToken: refreshToken,
		PublicKey:    pubPEMString,
		UserID:       foundAccount.UserID,
	})
	if err != nil {
		fmt.Println(err)
		return pb.LoginResponse{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}
	// save public key into kong gateway
	err = a.utils.CreateConsumer(foundAccount.UserAccount, foundAccount.UserID.String())
	if err != nil {
		fmt.Println(err)
		return pb.LoginResponse{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}

	err = a.utils.AddConsumerToACLGroup(foundAccount.UserAccount, "member")
	if err != nil {
		fmt.Println(err)
		return pb.LoginResponse{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}

	err = a.utils.AddConsumerJWTCredential(foundAccount.UserAccount, pubPEMString, foundAccount.UserID.String())
	if err != nil {
		fmt.Println(err)
		return pb.LoginResponse{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}

	return pb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
