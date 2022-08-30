package mongo

import (
	"context"
	"lucy/cashier/domain"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *userRepository) InsertToken(ctx context.Context, userId, tokenId string) (*domain.User, int, error) {
	filter := bson.M{"uuid": userId}
	update := bson.M{"$push": bson.M{"tokens": tokenId}}

	_, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	user, code, err := repo.FindUserBy(ctx, "uuid", userId, false)
	if err != nil {
		return nil, code, err
	}

	user.Tokens = []string{tokenId}
	return user, http.StatusOK, nil
}

func (repo *userRepository) RemoveToken(ctx context.Context, userId, tokenId string) (*domain.User, int, error) {
	filter := bson.M{"uuid": userId}
	update := bson.M{"$pull": bson.M{"tokens": tokenId}}

	_, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	user, code, err := repo.FindUserBy(ctx, "uuid", userId, false)
	if err != nil {
		return nil, code, err
	}

	user.Tokens = []string{tokenId}

	return user, http.StatusOK, nil
}
