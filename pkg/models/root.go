package models

// Root is a list of availble resource endpoints
type Root struct {
}

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

// Pagination options
type PaginationOptions struct {
	Limit  int
	Offest int
}
