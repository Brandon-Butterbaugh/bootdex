package main

type LocationAreaList struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationStub `json:"results"`
}

type LocationStub struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
