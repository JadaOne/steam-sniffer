package config

import "os"

type Settings struct {
	MongoDBURI   string
	DatabaseName string
	AppKey       string
}

var GlobalSettings Settings

func LoadSettings() {
	GlobalSettings = Settings{
		MongoDBURI:   os.Getenv("MONGODB_URI"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
		AppKey:       os.Getenv("AppKey"),
	}
}
