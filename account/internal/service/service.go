package service

import (
	"errors"
	"github.com/loctodale/go_api_hubs_microservice/account/database"
	repository2 "github.com/loctodale/go_api_hubs_microservice/account/internal/repository"
	"github.com/loctodale/go_api_hubs_microservice/account/utils"
)

type Service interface {
	PostAccount(userAccount string, userPassword string) error
	GetAccounts() ([]database.GetAccountsRow, error)
	RegisterAccount(userAccount string) error
}
type accountService struct {
	repository repository2.Repository
	utils      utils.CryptoUtils
}

func NewAccountService(repository repository2.Repository, utils utils.CryptoUtils) Service {
	return &accountService{
		repository,
		utils,
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
	if err = a.repository.CreateNewAccount(accountBaseModel); err != nil {
		return err
	}
	return nil
}

func (a *accountService) GetAccounts() ([]database.GetAccountsRow, error) {
	result := a.repository.GetAccounts()

	return result, nil
}

func (a *accountService) RegisterAccount(userAccount string) error {
	isExisted, err := a.repository.CheckUserBaseExists(userAccount)
	if err != nil {
		return err
	}
	if isExisted != 0 {
		return errors.New("Account already existed")
	}
	return nil
}
