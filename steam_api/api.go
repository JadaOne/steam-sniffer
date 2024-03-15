package steam_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const SteamUrl = "https://store.steampowered.com/"
const SteamAPIUrl = SteamUrl + "api/"

const ApiSteamUrl = "http://api.steampowered.com/"

func GetAppDetails(apiKey string, filters []string, appIds ...int) ([]AppData, error) {

	if len(filters) == 0 && len(appIds) > 1 {
		return []AppData{}, fmt.Errorf("to query multiple apps, at least one filter should be provided")
	}

	requestURL := SteamAPIUrl + "appdetails"

	// Create a new URL values object
	params := url.Values{}
	// Add the API key and App ID as query parameters
	params.Add("key", apiKey)

	for _, appId := range appIds {
		params.Add("appids", strconv.Itoa(appId))
	}

	// Add any filters as query parameters
	for _, filter := range filters {
		params.Add("filters", filter)
	}

	// Append the encoded parameters to the URL
	requestURL = requestURL + "?" + params.Encode()

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1) // !TODO what is it and how to handle?
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	var response map[string]AppResponse
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		fmt.Printf("client: could not unmarshal json: %s\n", err)
		os.Exit(1)
	}

	result := make([]AppData, 0)
	for _, value := range response {
		if value.Success == true {
			result = append(result, value.Data)
		}
	}

	return result, nil
}

func GetAppNews(apiKey string, appId int, count, maxLength int) (AppNews, error) {

	requestURL := ApiSteamUrl + "ISteamNews/GetNewsForApp/v0002/"
	params := url.Values{}

	if maxLength == 0 {
		maxLength = 300
	}

	// Add the API key and App ID as query parameters
	params.Add("key", apiKey)
	params.Add("count", strconv.Itoa(count))
	params.Add("format", "json")
	params.Add("appId", strconv.Itoa(appId))
	params.Add("maxlenth", strconv.Itoa(maxLength))

	// Append the encoded parameters to the URL
	requestURL = requestURL + "?" + params.Encode()

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1) // !TODO what is it and how to handle?
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	var response NewsResponse
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		fmt.Printf("client: could not unmarshal json: %s\n", err)
		os.Exit(1)
	}

	return response.Appnews, nil
}
