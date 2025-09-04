package db

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DdConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func LoadDbConfig() DdConfig {
	return DdConfig{
		Host:     getEnv("DB_HOST", "mongodb"),
		Port:     getEnvAsInt("DB_PORT", 27017),
		User:     getEnv("DB_USER", "admin"),
		Password: getEnv("DB_PASSWORD", "password123"),
		Dbname:   getEnv("DB_DATABASE", "bitscore"),
	}
}

func ConnectDb() (*mongo.Database, error) {
	config := LoadDbConfig()
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin",
		config.User, config.Password, config.Host, config.Port, config.Dbname)

	clientOptions := options.Client().ApplyURI(dsn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(config.Dbname)
	return db, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else {
		return defaultValue
	}
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
