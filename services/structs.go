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

type App struct {
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
	CheckPrice          bool                  `json:"checkPrice"`
	CheckNews           bool                  `json:"checkNews"`
}

func FromAppData(appData steam_api.AppData, checkPrice, checkNews bool) App {
	price := Price{
		Currency:       appData.PriceOverview.Currency,
		Price:          appData.PriceOverview.Initial,
		PriceFormatted: strconv.Itoa(appData.PriceOverview.Final/100) + "," + strconv.Itoa(appData.PriceOverview.Final%100),
	}
	currentLowest := Price{}
	if appData.PriceOverview.Initial > appData.PriceOverview.Final {
		currentLowest = Price{
			Currency:       appData.PriceOverview.Currency,
			Price:          appData.PriceOverview.Final,
			PriceFormatted: strconv.Itoa(appData.PriceOverview.Final/100) + "," + strconv.Itoa(appData.PriceOverview.Final%100),
		}
	} else {
		currentLowest = price
	}
	price.FormatPrice()
	currentLowest.FormatPrice()
	return App{
		Type:                appData.Type,
		Name:                appData.Name,
		SteamAppid:          appData.SteamAppid,
		RequiredAge:         appData.RequiredAge,
		IsFree:              appData.IsFree,
		ControllerSupport:   appData.ControllerSupport,
		DLC:                 appData.DLC,
		DetailedDescription: appData.DetailedDescription,
		AboutTheGame:        appData.AboutTheGame,
		ShortDescription:    appData.ShortDescription,
		SupportedLanguages:  appData.SupportedLanguages,
		Reviews:             appData.Reviews,
		Website:             appData.Website,
		Developers:          appData.Developers,
		Publishers:          appData.Publishers,
		Price:               price,
		CurrentLowest:       currentLowest,
		Platforms:           appData.Platforms,
		Metacritic:          appData.Metacritic,
		Categories:          appData.Categories,
		Genres:              appData.Genres,
		ReleaseDate:         appData.ReleaseDate,
		CheckPrice:          checkPrice,
		CheckNews:           checkNews,
	}
}

type AppInitial struct {
	AppId      int  `json:"appId"`
	CheckPrice bool `json:"checkPrice"`
	CheckNews  bool `json:"checkNews"`
}
