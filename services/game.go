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

func InitiateApp(appInit AppInitial) error {

	clientOptions := options.Client().ApplyURI(config.GlobalSettings.MongoDBURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(config.GlobalSettings.DatabaseName).Collection("apps")
	filter := bson.D{{"steamappid", appInit.AppId}}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		return errors.New(fmt.Sprintf("App %d already exists.", appInit.AppId))
	}

	apps, err := steam_api.GetAppDetails(config.GlobalSettings.AppKey, []string{}, appInit.AppId)

	if err != nil {
		return err
	}

	appData := apps[0]

	app := FromAppData(appData, appInit.CheckPrice, appInit.CheckNews)

	insertResult, err := collection.InsertOne(context.TODO(), app)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)

	return nil
}
func InitiateApps() {

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

	filter := bson.D{{}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var apps []AppInitial
	if err = cursor.All(context.TODO(), &apps); err != nil {
		log.Fatal(err)
	}

	for _, app := range apps {
		log.Println(app)
		err = InitiateApp(app)
		fmt.Println(err)
	}
}
