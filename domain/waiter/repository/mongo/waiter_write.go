package mongo

import (
	"context"
	"lucy/cashier/domain"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *waiterMongoRepository) UpsertWaiter(ctx context.Context, data *domain.Waiter) (*domain.Waiter, int, error) {
	var waiter domain.Waiter
	var contents bson.M

	filter := bson.M{"uuid": data.UUID}

	countWaiter, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if countWaiter > 0 {
		updatedAt := time.Now().UTC().UnixMicro()

		update := bson.M{
			"$set": bson.M{
				"branch_uuid": data.BranchUUID,
				"name": data.Name,
				"updated_at": updatedAt,
			},
		}

		contents = update
	} else {
		insert := bson.M{
			"$set": bson.M{
				"branch_uuid": data.BranchUUID,
				"name": data.Name,
				"created_at": data.CreatedAt,
			},
		}

		contents = insert
	}

	opts := options.Update().SetUpsert(true)

	if _, err := repo.Collection.UpdateOne(ctx, filter, contents, opts); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&waiter); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &waiter, http.StatusOK, nil
}

func (repo *waiterMongoRepository) DeleteWaiter(ctx context.Context, id string) (*domain.Waiter, int, error) {
	var waiter domain.Waiter

	filter := bson.M{"uuid": id}
	update := bson.A{
				bson.M{
					"$set": bson.M{
						"deleted_at": bson.M{
							"$ifNull": bson.A{
								"$deleted_at",
								time.Now().UTC().UnixMicro(),
							},
						},
					},
				},
			}

	result, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if result.MatchedCount == 0 {
		return nil, http.StatusNotFound, nil
	}

	singleResult := repo.Collection.FindOne(ctx, bson.M{"uuid": id})

	if err = singleResult.Decode(&waiter); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &waiter, http.StatusOK, nil
}