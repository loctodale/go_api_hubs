package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/loctodale/go_api_hubs_microservice/account"
)

type Server struct {
	accountClient *account.Client
	//catalogClient *catalog.Client
	//orderClient   *order.Client
}

func NewGraphqlServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}
	//catalogClient, err := catalogClient.NewClient(catalogUrl)
	//if err != nil {
	//	accountClient.Close()
	//	return nil, err
	//}
	//
	//orderClient, err := orderClient, NewClient(orderUrl)
	//if err != nil {
	//	accountClient.Close()
	//	catelogClient.Close()
	//	return nil, err
	//}

	return &Server{
		accountClient,
		//catalogClient,
		//orderClient,
	}, nil
}

func (s *Server) Account() AccountResolver {
	return &accountResolver{
		server: s,
	}
}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}
}

func (s *Server) ToExcutetableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
