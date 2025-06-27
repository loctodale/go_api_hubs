package main

import "context"

type accountResolver struct {
	server *Server
}

func (a accountResolver) Orders(ctx context.Context, obj *Account) ([]*Order, error) {
	//TODO implement me
	panic("implement me")
}

//func (r *queryResolver) Orders(ctx context.Context, obj *Account) ([]*Order, error)  {
//
//}
