package main

import "context"

type mutationResolver struct {
	server *Server
}

func (m mutationResolver) Login(ctx context.Context, model *LoginModel) (*LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) CreateAccount(ctx context.Context, account *AccountInput) (*Account, error) {
	err := m.server.accountClient.PostAccount(account.UserAccount, account.UserPassword)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (m mutationResolver) RegisterAccount(ctx context.Context, account *RegisterAccount) (*SampleResponse, error) {
	err := m.server.accountClient.RegisterAccount(account.UserAccount)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (m mutationResolver) CreateProduct(ctx context.Context, product *ProductInput) (*Product, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) CreateOrder(ctx context.Context, order *OrderInput) (*Order, error) {
	//TODO implement me
	panic("implement me")
}

//func (r *mutationResolver) createAccount(ctx context.Context, in AccountInput) (*Account, error) {
//
//}
//
//func (r *mutationResolver) createProduct(ctx context.Context, in ProductInput) (*Product, error) {
//
//}
//
//func (r *mutationResolver) createOrder(ctx context.Context, in OrderInput) (*Order, error) {
//
//}
