package config

import (
	"context"
	"fmt"
	"log"
	"os"

	models "crudecho/model"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	MongoClient *mongo.Client
	PostgresDB  *gorm.DB
}

func Init() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoClient := connectMongo()
	postgresDB := connectPostgres()

	return &Config{
		MongoClient: mongoClient,
		PostgresDB:  postgresDB,
	}
}

func connectMongo() *mongo.Client {
	mongoURI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func connectPostgres() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Debug().AutoMigrate(&models.Manager{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	return db
}
