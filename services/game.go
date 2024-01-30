package services

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"steam-checker/config"
	"steam-checker/steam_api"
)

func InitiateGame(appId int) error {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// Get a handle for your collection
	collection := client.Database(config.GlobalSettings.DatabaseName).Collection("api")
	filter := bson.D{{"steamappid", appId}}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		return errors.New(fmt.Sprintf("Game %s already exists.", appId))
	}

	games, err := steam_api.GetAppDetails(config.GlobalSettings.AppKey, []string{}, appId)

	if err != nil {
		return err
	}

	gameData := games[0]

	game := FromGameData(gameData)
	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), game)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)

	return nil
}
func InitiateGames() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get a handle for your collection
	collection := client.Database(config.GlobalSettings.DatabaseName).Collection("games_list")

	// Define an empty filter to match all documents
	filter := bson.D{{}}

	// Find all documents
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var games []GameId
	if err = cursor.All(context.TODO(), &games); err != nil {
		log.Fatal(err)
	}

	for _, game := range games {
		// Now you can work with each game
		log.Println(game)
		err = InitiateGame(game.AppId)
		fmt.Println(err)
	}
}
