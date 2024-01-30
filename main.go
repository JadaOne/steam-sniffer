package main

import (
	"steam-checker/config"
	"steam-checker/services"
)

func main() {
	//games, err := steam_api.GetAppDetails(apiKey, []string{}, "363440")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(games)

	config.LoadSettings()

	//err := services.InitiateGame("2071280")
	//
	//if err != nil {
	//	panic(err)
	//}
	services.InitiateGames()
}
