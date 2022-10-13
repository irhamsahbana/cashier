package mongo

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/lib/logger"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
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
		logger.Log(logrus.Fields{
			"error": err,
		}).Error("failed to find user by uuid")
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
		logger.Log(logrus.Fields{
			"error": err,
		}).Error("failed to remove token")
		return nil, http.StatusInternalServerError, err
	}

	user, code, err := repo.FindUserBy(ctx, "uuid", userId, false)
	if err != nil {
		logger.Log(logrus.Fields{
			"error": err,
		}).Error("failed to find user by uuid")
		return nil, code, err
	}

	user.Tokens = []string{tokenId}

	return user, http.StatusOK, nil
}

func (repo *userRepository) UpsertUser(ctx context.Context, user *domain.User) (*domain.User, int, error) {
	filter := bson.D{
		{Key: "$and", Value: bson.A{
			bson.D{{Key: "uuid", Value: user.UUID}},
			bson.D{{Key: "branch_uuid", Value: user.BranchUUID}},
		}},
	}

	countUser, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if countUser == 0 { // insert
		_, err := repo.Collection.InsertOne(ctx, user)
		if err != nil {
			logger.Log(logrus.Fields{}).Error(err)
			return nil, http.StatusInternalServerError, err
		}
	} else { // update
		updatedAt := time.Now().UnixMicro()
		user.UpdatedAt = &updatedAt

		db, code, err := repo.FindUserBy(ctx, "uuid", user.UUID, true)
		if err != nil {
			logger.Log(logrus.Fields{}).Error(err)
			return nil, code, err
		}
		user.CreatedAt = db.CreatedAt
		update := bson.M{"$set": user}

		_, err = repo.Collection.UpdateOne(ctx, filter, update)
		if err != nil {
			logger.Log(logrus.Fields{}).Error(err)
			return nil, http.StatusInternalServerError, err
		}
	}

	return user, http.StatusOK, nil
}
