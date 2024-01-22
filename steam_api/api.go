package steam_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const SteamUrl = "https://store.steampowered.com/api/"

func GetAppDetails(apiKey string, filters []string, appIds ...string) ([]GameData, error) {

	if len(filters) == 0 && len(appIds) > 1 {
		return []GameData{}, fmt.Errorf("to query multiple apps, at least one filter should be provided")
	}

	requestURL := SteamUrl + "appdetails"

	// Create a new URL values object
	params := url.Values{}
	// Add the API key and App ID as query parameters
	params.Add("key", apiKey)

	for _, appId := range appIds {
		params.Add("appids", appId)
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

	var response map[string]GameResponse
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		fmt.Printf("client: could not unmarshal json: %s\n", err)
		os.Exit(1)
	}

	result := make([]GameData, 0)
	for _, value := range response {
		if value.Success == true {
			result = append(result, value.Data)
		}
	}

	return result, nil
}
