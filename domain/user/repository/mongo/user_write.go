package mongo

import (
	"context"
	"lucy/cashier/domain"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *userRepository) UpdateUser(ctx context.Context, data *domain.User) (*domain.User, int, error) {
	var user domain.User

	filter := bson.M{"uuid": data.UUID}
	opts := options.Update().SetUpsert(true)

	updateFields := bson.M{
		"$set": bson.M{

			"updated_at": time.Now().UTC().UnixMicro(),
		},
	}

	if _, err := repo.Collection.UpdateOne(ctx, filter, updateFields, opts); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusOK, nil
}

func (repo *userRepository) GenerateTokens(ctx context.Context, userId, accessToken, refreshToken string) (aToken, rToken string, code int, err error) {

	tokens := domain.Token{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}

	filter := bson.M{"uuid": userId}
	updateFields := bson.A{
						bson.M{
							"$set": bson.M{
								"tokens": bson.M{
									"$ifNull": bson.A{
										bson.M{"$concatArrays": bson.A{"$tokens", bson.A{tokens}}},
										bson.A{tokens},
									},
								},
							},
						},
					}

	_, err = repo.Collection.UpdateOne(ctx, filter, updateFields)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}



	return accessToken, refreshToken, http.StatusOK, nil
}

func (repo *userRepository) RefreshToken(ctx context.Context, oldAT, oldRT, newAT, newRT, userId string) (aToken, rToken string, code int, err error) {
	filter := bson.M{"uuid": userId}

	updateFields := bson.M{
		"$pull" : bson.M{
			"tokens": bson.M{
				"access_token": oldAT,
				"refresh_token": oldRT,
			},
		},
	}

	result, err := repo.Collection.UpdateOne(ctx, filter, updateFields)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if result.ModifiedCount == 0 {
		return "", "", http.StatusNotFound, nil
	}

	updateFields = bson.M{
		"$push" : bson.M{
			"tokens": bson.M{
				"access_token": newAT,
				"refresh_token": newRT,
			},
		},
	}

	_, err = repo.Collection.UpdateOne(ctx, filter, updateFields)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	return newAT, newRT, http.StatusOK, nil
}

func (repo *userRepository) RevokeToken(ctx context.Context, accessToken string) (code int, err error) {
	filter := bson.M{"tokens.access_token": accessToken}

	updateFields := bson.M{
		"$pull" : bson.M{
			"tokens": bson.M{
				"access_token": accessToken,
			},
		},
	}

	_, err = repo.Collection.UpdateOne(ctx, filter, updateFields)
	if err != nil {
		return  http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

