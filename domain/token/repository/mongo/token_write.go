package mongo

import (
	"context"
	"lucy/cashier/domain"
	"net/http"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *tokenRepository) GenerateTokens(ctx context.Context, userId, accessToken, refreshToken string) (aToken, rToken string, code int, err error) {
	tokens := domain.Token{
		UUID:          uuid.NewString(),
		TokenableUUID: userId,
		TokenableType: repo.Type,
		AccessToken:   accessToken,
		RefreshToken:  refreshToken,
	}

	doc := bson.D{
		{Key: "uuid", Value: tokens.UUID},
		{Key: "tokenable_uuid", Value: tokens.TokenableUUID},
		{Key: "tokenable_type", Value: tokens.TokenableType},
		{Key: "access_token", Value: tokens.AccessToken},
		{Key: "refresh_token", Value: tokens.RefreshToken},
	}

	_, err = repo.Collection.InsertOne(ctx, doc)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	return accessToken, refreshToken, http.StatusOK, nil
}
