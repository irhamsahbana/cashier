package mongo

import (
	"context"
	"lucy/cashier/domain"
	"net/http"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *tokenRepository) GenerateTokens(ctx context.Context, userId, accessToken, refreshToken string) (tokenUUID string, code int, err error) {
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
		return "", http.StatusInternalServerError, err
	}

	return tokens.UUID, http.StatusOK, nil
}

func (repo *tokenRepository) RefreshTokens(ctx context.Context, userId, oldAT, oldRT, newAT, newRT string) (tokenUUID string, code int, err error) {
	filter := bson.D{
		{Key: "access_token", Value: oldAT},
		{Key: "refresh_token", Value: oldRT},

		{Key: "tokenable_uuid", Value: userId},
		{Key: "tokenable_type", Value: repo.Type},
	}

	_, err = repo.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return repo.GenerateTokens(ctx, userId, newAT, newRT)
}

func (repo *tokenRepository) RevokeTokens(ctx context.Context, userId, accessToken string) (uuid string, code int, err error) {
	filter := bson.D{
		{Key: "access_token", Value: accessToken},
		{Key: "tokenable_uuid", Value: userId},
		{Key: "tokenable_type", Value: repo.Type},
	}

	var doc domain.Token

	repo.Collection.FindOne(ctx, filter).Decode(&doc)

	_, err = repo.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return doc.UUID, http.StatusOK, nil
}
