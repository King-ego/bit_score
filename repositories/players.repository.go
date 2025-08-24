package repositories

import "go.mongodb.org/mongo-driver/mongo"

type ICreatePlayer struct {
	Name string
}

type PlayersRepository interface {
	Create(player ICreatePlayer) error
}

type playersRepository struct {
	db *mongo.Collection
}

func NewPlayersRepository(db *mongo.Database) PlayersRepository {
	return &playersRepository{
		db: db.Collection("players"),
	}
}

func (r *playersRepository) Create(player ICreatePlayer) error {
	_, err := r.db.InsertOne(nil, map[string]interface{}{
		"name": player.Name,
	})
	if err != nil {
		return err
	}
	return nil
}
