package repositories

import (
	"bit_score/entities"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ICreateTeam struct {
	Name         string
	PrimaryColor string
	SecondColor  string
}

type IUpdateTeam struct {
	Name         string
	PrimaryColor string
	SecondColor  string
}

type TeamsRepository interface {
	Create(team ICreateTeam) error
	GetAll() ([]entities.Teams, error)
	GetByID(id string) (*entities.Teams, error)
	Update(id string, team IUpdateTeam) error
	Delete(id string) error
}

type teamsRepository struct {
	db *mongo.Collection
}

func NewTeamsRepository(db *mongo.Database) TeamsRepository {
	return &teamsRepository{
		db: db.Collection("teams"),
	}
}

func (r *teamsRepository) Create(team ICreateTeam) error {
	_, err := r.db.InsertOne(context.Background(), map[string]interface{}{
		"name":          team.Name,
		"primary_color": team.PrimaryColor,
		"second_color":  team.SecondColor,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *teamsRepository) GetAll() ([]entities.Teams, error) {
	var teams []entities.Teams
	cursor, err := r.db.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var team entities.Teams
		if err := cursor.Decode(&team); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	return teams, nil
}

func (r *teamsRepository) GetByID(id string) (*entities.Teams, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var team entities.Teams
	err = r.db.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&team)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func (r *teamsRepository) Update(id string, team IUpdateTeam) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{}
	if team.Name != "" {
		update["name"] = team.Name
	}
	if team.PrimaryColor != "" {
		update["primary_color"] = team.PrimaryColor
	}
	if team.SecondColor != "" {
		update["second_color"] = team.SecondColor
	}

	_, err = r.db.UpdateOne(context.Background(), bson.M{"_id": objectID}, bson.M{"$set": update})
	if err != nil {
		return err
	}

	return nil
}

func (r *teamsRepository) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.db.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}
