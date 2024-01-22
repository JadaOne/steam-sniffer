package main

import (
	"fmt"
	"steam-checker/steam_api"
)

const apiKey = "3CE1C3E26481CE20CBA13961D1123A52"

func main() {
	games, err := steam_api.GetAppDetails(apiKey, []string{}, "363440")
	if err != nil {
		panic(err)
	}
	fmt.Println(games)
}
