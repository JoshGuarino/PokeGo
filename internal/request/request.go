package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JoshGuarino/PokeGo/internal/cache"
)

// Make GET request
func Get(url string) ([]byte, error) {
	// Create new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}

	// Send request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Read response body
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Return response body
	return body, nil
}

// Get data from URL or cache
func GetData[T any](url string) (*T, error) {
	// Check for cached data and return if found
	data, found := cache.C.Get(url)
	if found {
		return data.(*T), nil
	}

	// Make GET request
	body, errReq := Get(url)
	if errReq != nil {
		return nil, errReq
	}

	// Unmarshal JSON response
	dataStruct := new(T)
	errJson := json.Unmarshal(body, dataStruct)
	if errJson != nil {
		return nil, errJson
	}

	// Cache ResourceList if active and return
	if cache.C.GetActive() {
		cache.C.Set(url, dataStruct)
		return dataStruct, nil
	}
	return dataStruct, nil
}

// Make GET request for list of resource
func GetResourceList[T any](url string, limit int, offset int) (*T, error) {
	// Append limit and offset to URL
	url = fmt.Sprintf("%s?limit=%d&offset=%d", url, limit, offset)
	return GetData[T](url)
}

// Make GET request for a specifc resource
func GetResource[T any](url string) (*T, error) {
	return GetData[T](url)
}
