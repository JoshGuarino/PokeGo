package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JoshGuarino/PokeGo/internal/cache"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Cache instance
var c *cache.Cache

// Initialize cache
func init() {
	c = cache.NewCache()
}

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

// Make GET request for list of resource
func GetResourceList(url string, options models.PaginationOptions) (*models.ResourceList, error) {
	// Set default limit if not provided and build URL
	if options.Limit == 0 {
		options.Limit = 20
	}
	url = fmt.Sprintf("%s?offset=%d&limit=%d", url, options.Offest, options.Limit)

	// Create new ResourceList instance
	resourceList := models.ResourceList{}

	// Check for cached data and return if found
	data, found := c.Get(url)
	if found {
		return data.(*models.ResourceList), nil
	}

	// Make GET request
	body, errReq := Get(url)
	if errReq != nil {
		return nil, errReq
	}

	// Unmarshal JSON response
	errJson := json.Unmarshal(body, &resourceList)
	if errJson != nil {
		return nil, errJson
	}

	// Cache ResourceList if active and return
	if c.GetActive() {
		c.Set(url, &resourceList)
		return &resourceList, nil
	}
	return &resourceList, nil
}

// Make GET request for a specifc resource
func GetSpecificResource[T any](url string) (*T, error) {
	// Check for cached data and return if found
	data, found := c.Get(url)
	if found {
		return data.(*T), nil
	}

	// Make GET request
	body, errReq := Get(url)
	if errReq != nil {
		return nil, errReq
	}

	// Unmarshal JSON response
	resource := new(T)
	errJson := json.Unmarshal(body, resource)
	if errJson != nil {
		return nil, errJson
	}

	// Cache Resource if active and return
	if c.GetActive() {
		c.Set(url, resource)
		return resource, nil
	}
	return resource, nil
}
