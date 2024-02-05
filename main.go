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

	//err := services.InitiateApp("2071280")
	//
	//if err != nil {
	//	panic(err)
	//}
	services.InitiateApps()
	//res, err := steam_api.GetAppNews(config.GlobalSettings.AppKey, 2071280, 0, 300) // TODO how good is that approach?
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(res)
}
