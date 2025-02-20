package request

// Resource list for an endpoint
type ResourceList struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

// Resource list result
type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
