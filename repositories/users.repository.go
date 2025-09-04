package repositories

import (
	"bit_score/entities"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository interface {
	GetByUserName(username string) (entities.Users, error)
}

type usersRepository struct {
	collection *mongo.Collection
}

func NewUsersRepository(db *mongo.Database) UsersRepository {
	return &usersRepository{
		collection: db.Collection("users"),
	}
}

func (r *usersRepository) GetByUserName(username string) (entities.Users, error) {
	var user entities.Users
	filter := map[string]interface{}{"username": username}
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return entities.Users{}, nil
		}
		return entities.Users{}, err
	}
	return user, nil
}
