package authuser

import (
	"context"
	"lucy/cashier/bootstrap"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUser(id string) (*dto.UserResponse, int, error) {
	c := context.Background()
	ctx, cancel := context.WithTimeout(c, bootstrap.App.Config.GetDuration("context.timeout")*time.Second)
	defer cancel()

	DB := bootstrap.App.Mongo.Database(bootstrap.App.Config.GetString("mongo.name"))
	collectionUsers := DB.Collection("users")
	collectionUserRoles := DB.Collection("user_roles")

	filter := bson.M{
		"$and": bson.A{
			bson.M{"uuid": id},
			bson.M{"deleted_at": nil},
		},
	}

	resultUser := collectionUsers.FindOne(ctx, filter)
	if resultUser.Err() != nil {
		if resultUser.Err().Error() == mongo.ErrNoDocuments.Error() {
			return nil, http.StatusNotFound, resultUser.Err()
		}
		return nil, http.StatusInternalServerError, resultUser.Err()
	}

	var user domain.User
	err := resultUser.Decode(&user)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	filter = bson.M{
		"uuid": user.RoleUUID,
	}

	resultRole := collectionUserRoles.FindOne(ctx, filter)
	if resultRole.Err() != nil {
		if resultUser.Err().Error() == mongo.ErrNoDocuments.Error() {
			return nil, http.StatusNotFound, resultUser.Err()
		}
		return nil, http.StatusInternalServerError, resultRole.Err()
	}

	var userRole domain.UserRole
	err = resultRole.Decode(&userRole)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var resp dto.UserResponse
	resp.UUID = user.UUID
	resp.BranchUUID = user.BranchUUID
	resp.Name = user.Name
	resp.Role = userRole.Name
	respCreatedAt := time.UnixMicro(user.CreatedAt).UTC()
	resp.CreatedAt = &respCreatedAt
	if user.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*user.UpdatedAt).UTC()
		resp.UpdatedAt = &respUpdatedAt
	}
	if user.DeletedAt != nil {
		respDeletedAt := time.UnixMicro(*user.DeletedAt).UTC()
		resp.DeletedAt = &respDeletedAt
	}

	return &resp, http.StatusOK, nil
}
