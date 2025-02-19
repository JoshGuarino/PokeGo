package request

import (
	"encoding/json"
	// "fmt"
	"io"
	"net/http"
	// "os"
	// pokego "github.com/JoshGuarino/PokeGo/pkg"
)

// Make GET request
func Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Make GET request for list of resource
// func GetResourceList(url string) pokego.ResourceList {
// 	body := Get(url)
// 	resourceList := pokego.ResourceList{}
// 	err := json.Unmarshal(body, &resourceList)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	return resourceList
// }

// Make GET request for a specifc resource
func GetSpecificResource[T any](url string) (*T, error) {
	body, errReq := Get(url)
	if errReq != nil {
		return nil, errReq
	}
	resource := new(T)
	errJson := json.Unmarshal(body, resource)
	if errJson != nil {
		return nil, errJson
	}
	return resource, nil
}
