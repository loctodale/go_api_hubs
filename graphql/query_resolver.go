package main

import (
	"context"
)

type queryResolver struct {
	server *Server
}

func (q queryResolver) Accounts(ctx context.Context, paginations *PaginationInput) ([]*Account, error) {
	result, err := q.server.accountClient.GetAccounts()
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}

	for _, p := range result.Account {
		accounts = append(accounts, &Account{
			Name: p.Name,
			ID:   p.Id,
		})
	}
	return accounts, nil
}

func (q queryResolver) Products(ctx context.Context, paginations *PaginationInput) ([]*Product, error) {
	//TODO implement me
	panic("implement me")
}

//func (r *queryResolver) Accounts(ctx context.Context, pagination *PaginationInput, id string) ([]*Account, error) {
//
//}
//
//func (r *queryResolver) Products(ctx context.Context, pagination *PaginationInput, query *string, id string) ([]*Product, error) {
//
//}
