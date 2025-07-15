package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/loctodale/go_api_hubs_microservice/account/database"
	"github.com/loctodale/go_api_hubs_microservice/account/global"
)

type TokenRepository interface {
	GetTokenByUser(userId pgtype.UUID) (database.GetTokenByUserRow, error)
	AddToken(params database.CreateNewTokenParams) error
}

type tokenRepository struct {
	conn    *pgx.Conn
	queries *database.Queries
}

func NewTokenRepository() TokenRepository {
	return &tokenRepository{
		conn:    global.Pdb,
		queries: database.New(global.Pdb),
	}
}

func (t tokenRepository) GetTokenByUser(userId pgtype.UUID) (database.GetTokenByUserRow, error) {
	token, err := t.queries.GetTokenByUser(global.Ctx, userId)
	if err != nil {
		return database.GetTokenByUserRow{}, err
	}

	return token, nil
}

func (t tokenRepository) AddToken(params database.CreateNewTokenParams) error {
	err := t.queries.CreateNewToken(global.Ctx, params)
	if err != nil {
		return err
	}
	return nil
}
