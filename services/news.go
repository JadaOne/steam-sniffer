package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"steam-checker/config"
	"steam-checker/steam_api"
)

func InsertNews(news []steam_api.News) error {
	clientOptions := options.Client().ApplyURI(config.GlobalSettings.MongoDBURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(config.GlobalSettings.DatabaseName).Collection("news")

	insertResult, err := collection.InsertMany(context.TODO(), news)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)

	return nil
}

func ProcessNews(appId int, news steam_api.News) error {
	return nil
}

func CollectNews() error {
	return nil
}

func InitialCollectNews() error {
	// !TODO add logic to collect news
	// if any news exists we just do a regular job
	// if there are none, run with 0 count to get total number of news
	// or do 0 count request each time to calculate number of news?
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(config.GlobalSettings.DatabaseName).Collection("apps_list")

	filter := bson.D{{"checkNews", true}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var apps []AppInitial
	if err = cursor.All(context.TODO(), &apps); err != nil {
		log.Fatal(err)
	}
	toBeCollected := make([]int, 0)

	for _, app := range apps {
		toBeCollected = append(toBeCollected, app.AppId)
	}

	for _, app := range toBeCollected {
		_, err = steam_api.GetAppNews(config.GlobalSettings.AppKey, app, 0, 300)
		if err != nil {
			log.Fatal(err)
		}
		news, err := steam_api.GetAppNews(config.GlobalSettings.AppKey, app, 3, 300)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(news)
		err = InsertNews(news.NewsItems)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
