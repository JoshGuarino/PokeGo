package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/internal/logger"
)

// Make GET request
func Get(url string) ([]byte, error) {
	// Create new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers and make request
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Read response body and return
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Get data from URL or cache
func GetData[T any](url string) (T, error) {
	// Create new data struct
	dataStruct := new(T)

	// Check for cached data and return if found
	data, found := cache.CACHE.Get(url)
	if found {
		return data.(T), nil
	}

	// Make GET request
	body, errReq := Get(url)
	if errReq != nil {
		return *dataStruct, errReq
	}

	// Unmarshal JSON response
	errJson := json.Unmarshal(body, dataStruct)
	if errJson != nil {
		return *dataStruct, errJson
	}

	// Cache ResourceList if active and return
	if cache.CACHE.Active() {
		cache.CACHE.Set(url, *dataStruct)
		return *dataStruct, nil
	}
	return *dataStruct, nil
}

// Make GET request for list of resource
func GetResourceList[T any](url string, limit int, offset int) (*T, error) {
	url = fmt.Sprintf("%s?limit=%d&offset=%d", url, limit, offset)
	logger.LOG.Info(
		"Get resource list called",
		"url", url,
	)
	data, err := GetData[*T](url)
	if err != nil {
		logger.LOG.Error("Error getting resource list", "err", err)
		return nil, err
	}
	return data, nil
}

// Make GET request for a specifc resource
func GetResource[T any](url string) (*T, error) {
	logger.LOG.Info(
		"Get resource called",
		"url", url,
	)
	data, err := GetData[*T](url)
	if err != nil {
		logger.LOG.Error("Error getting specifc resource", "err", err)
		return nil, err
	}
	return data, nil
}

// Make GET request for a slice of resources
func GetResourceSlice[T any](url string) ([]*T, error) {
	logger.LOG.Info(
		"Get resource slice called",
		"url", url,
	)
	data, err := GetData[[]*T](url)
	if err != nil {
		logger.LOG.Error("Error getting resource slice", "err", err)
		return nil, err
	}
	return data, nil
}
