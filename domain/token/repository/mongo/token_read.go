package mongo

import (
	"context"
	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type tokenRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
	Type       domain.TokenableType
}

func NewTokenMongoRepository(DB mongo.Database, t domain.TokenableType) domain.TokenRepositoryContract {
	return &tokenRepository{
		DB:         DB,
		Collection: *DB.Collection("tokens"),
		Type:       t,
	}
}

func (repo *tokenRepository) refreshToken(ctx context.Context, userId, oldAT, oldRT, newAT, newRT string) (aToken, rToken string, code int, err error) {
	return "", "", 200, nil
}

func (repo *tokenRepository) RevokeToken(ctx context.Context, accessToken string) (code int, err error) {
	return 200, err
}
