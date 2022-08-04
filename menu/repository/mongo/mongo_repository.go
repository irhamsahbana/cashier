package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format

type menuRepository struct {
	DB			mongo.Database
	Collection	mongo.Collection
}


// func NewMenuRepository(DB mongo.Database) domain.MenuRepositoryContract {
// 	return &menuRepository{
// 		DB,
// 		*DB.Collection(collectionName),
// 	}
// }

// func (m *menuRepository) InsertOne(ctx context.Context, menu *domain.Menu) (*domain.Menu, error) {

// }

// func (m *menuRepository) UpdateOne(ctx context.Context, id string) (*domain.Menu, error) {

// }

// func (m *menuRepository) DeleteOne(ctx context.Context, id string) (*domain.Menu, error) {

// }