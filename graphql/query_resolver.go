package main

import (
	"context"
)

type queryResolver struct {
	server *Server
}

func (q queryResolver) Accounts(ctx context.Context, paginations *PaginationInput, id *string) ([]*Account, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) Products(ctx context.Context, paginations *PaginationInput, id *string) ([]*Product, error) {
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
