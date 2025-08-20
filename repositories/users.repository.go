package repositories

import (
	"bit_score/entity"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepositoryInterface interface {
	GetByUserName(username string) (entity.Users, error)
}

type usersRepository struct {
	collection *mongo.Collection
}

func NewUsersRepository(db *mongo.Database) UsersRepositoryInterface {
	return &usersRepository{
		collection: db.Collection("users"),
	}
}

func (r *usersRepository) GetByUserName(username string) (entity.Users, error) {
	var user entity.Users
	filter := map[string]interface{}{"username": username}
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return entity.Users{}, nil
		}
		return entity.Users{}, err
	}
	return user, nil
}
