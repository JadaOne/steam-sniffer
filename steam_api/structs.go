package steam_api

import (
	"encoding/json"
	"strconv"
)

// Game details

type PriceOverview struct {
	Currency         string `json:"currency"`
	Initial          int    `json:"initial"`
	Final            int    `json:"final"`
	DiscountPercent  int    `json:"discount_percent"`
	InitialFormatted string `json:"initial_formatted"`
	FinalFormatted   string `json:"final_formatted"`
}

type Platforms struct {
	Windows bool `json:"windows"`
	Mac     bool `json:"mac"`
	Linux   bool `json:"linux"`
}

type Metacritic struct {
	Score int    `json:"score"`
	URL   string `json:"url"`
}

type Category struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

type Genre struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type Highlighted struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type ReleaseDate struct {
	ComingSoon bool   `json:"coming_soon"`
	Date       string `json:"date"`
}

type GameData struct {
	Type                string        `json:"type"`
	Name                string        `json:"name"`
	SteamAppid          int           `json:"steam_appid"`
	RequiredAge         string        `json:"-"`
	IsFree              bool          `json:"is_free"`
	ControllerSupport   string        `json:"controller_support"`
	DLC                 []int         `json:"dlc"`
	DetailedDescription string        `json:"detailed_description"`
	AboutTheGame        string        `json:"about_the_game"`
	ShortDescription    string        `json:"short_description"`
	SupportedLanguages  string        `json:"supported_languages"`
	Reviews             string        `json:"reviews"`
	Website             string        `json:"website"`
	Developers          []string      `json:"developers"`
	Publishers          []string      `json:"publishers"`
	PriceOverview       PriceOverview `json:"price_overview"`
	Platforms           Platforms     `json:"platforms"`
	Metacritic          Metacritic    `json:"metacritic"`
	Categories          []Category    `json:"categories"`
	Genres              []Genre       `json:"genres"`
	ReleaseDate         ReleaseDate   `json:"release_date"`
}

func (gd *GameData) UnmarshalJSON(data []byte) error {
	type Alias GameData
	aux := &struct {
		RequiredAge interface{} `json:"required_age"`
		*Alias
	}{
		Alias: (*Alias)(gd),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if v, ok := aux.RequiredAge.(float64); ok {
		gd.RequiredAge = strconv.Itoa(int(v))
	} else if v, ok := aux.RequiredAge.(string); ok {
		if i, err := strconv.Atoi(v); err == nil {
			gd.RequiredAge = strconv.Itoa(i)
		}
	}
	return nil
}

type GameResponse struct {
	Success bool     `json:"success"`
	Data    GameData `json:"data"`
}

// News

type NewsResponse struct {
	Appnews AppNews `json:"appnews"`
}

type AppNews struct {
	Appid     int    `json:"appid"`
	Newsitems []News `json:"newsitems"`
	Count     int    `json:"count"`
}

type News struct {
	Gid           string   `json:"gid"`
	Title         string   `json:"title"`
	Url           string   `json:"url"`
	IsExternalUrl bool     `json:"is_external_url"`
	Author        string   `json:"author"`
	Contents      string   `json:"contents"`
	Feedlabel     string   `json:"feedlabel"`
	Date          int      `json:"date"`
	Feedname      string   `json:"feedname"`
	FeedType      int      `json:"feed_type"`
	Appid         int      `json:"appid"`
	Tags          []string `json:"tags,omitempty"`
}
