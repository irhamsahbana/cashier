package bootstrap

import "go.mongodb.org/mongo-driver/mongo"

func Dummy(DB *mongo.Database) error {
	return nil
}
