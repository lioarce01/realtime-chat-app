package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}

func GetJWTSecret() string {
	fmt.Println("JWT_SECRET", os.Getenv("JWT_SECRET"))
	return os.Getenv("JWT_SECRET")
}

func GetPort() string {
	return os.Getenv("PORT")
}

func ConnectDB() {
	clientOptions := options.Client().ApplyURI(GetMongoURI())

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	DB = client.Database("chatapp")
}