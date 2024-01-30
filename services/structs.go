package services

import (
	"steam-checker/steam_api"
	"strconv"
)

type Price struct {
	Currency       string `json:"currency"`
	Price          int    `json:"initial"`
	PriceFormatted string `json:"initial_formatted"`
}

func (p Price) FormatPrice() {
	integerPart := p.Price / 100
	fractionalPart := p.Price % 100
	p.PriceFormatted = strconv.Itoa(integerPart) + "," + strconv.Itoa(fractionalPart)
}

type Game struct {
	Type                string                `json:"type"`
	Name                string                `json:"name"`
	SteamAppid          int                   `json:"steam_appid"`
	RequiredAge         string                `json:"required_age"`
	IsFree              bool                  `json:"is_free"`
	ControllerSupport   string                `json:"controller_support"`
	DLC                 []int                 `json:"dlc"`
	DetailedDescription string                `json:"detailed_description"`
	AboutTheGame        string                `json:"about_the_game"`
	ShortDescription    string                `json:"short_description"`
	SupportedLanguages  string                `json:"supported_languages"`
	Reviews             string                `json:"reviews"`
	Website             string                `json:"website"`
	Developers          []string              `json:"developers"`
	Publishers          []string              `json:"publishers"`
	Price               Price                 `json:"price_overview"`
	CurrentLowest       Price                 `json:"current_lowest"`
	Platforms           steam_api.Platforms   `json:"platforms"`
	Metacritic          steam_api.Metacritic  `json:"metacritic"`
	Categories          []steam_api.Category  `json:"categories"`
	Genres              []steam_api.Genre     `json:"genres"`
	ReleaseDate         steam_api.ReleaseDate `json:"release_date"`
}

func FromGameData(gameData steam_api.GameData) Game {
	price := Price{
		Currency:       gameData.PriceOverview.Currency,
		Price:          gameData.PriceOverview.Initial,
		PriceFormatted: strconv.Itoa(gameData.PriceOverview.Final/100) + "," + strconv.Itoa(gameData.PriceOverview.Final%100),
	}
	currentLowest := Price{}
	if gameData.PriceOverview.Initial > gameData.PriceOverview.Final {
		currentLowest = Price{
			Currency:       gameData.PriceOverview.Currency,
			Price:          gameData.PriceOverview.Final,
			PriceFormatted: strconv.Itoa(gameData.PriceOverview.Final/100) + "," + strconv.Itoa(gameData.PriceOverview.Final%100),
		}
	} else {
		currentLowest = price
	}
	price.FormatPrice()
	currentLowest.FormatPrice()
	return Game{
		Type:                gameData.Type,
		Name:                gameData.Name,
		SteamAppid:          gameData.SteamAppid,
		RequiredAge:         gameData.RequiredAge,
		IsFree:              gameData.IsFree,
		ControllerSupport:   gameData.ControllerSupport,
		DLC:                 gameData.DLC,
		DetailedDescription: gameData.DetailedDescription,
		AboutTheGame:        gameData.AboutTheGame,
		ShortDescription:    gameData.ShortDescription,
		SupportedLanguages:  gameData.SupportedLanguages,
		Reviews:             gameData.Reviews,
		Website:             gameData.Website,
		Developers:          gameData.Developers,
		Publishers:          gameData.Publishers,
		Price:               price,
		CurrentLowest:       currentLowest,
		Platforms:           gameData.Platforms,
		Metacritic:          gameData.Metacritic,
		Categories:          gameData.Categories,
		Genres:              gameData.Genres,
		ReleaseDate:         gameData.ReleaseDate,
	}
}

type GameId struct {
	AppId int `json:"appId"`
}
