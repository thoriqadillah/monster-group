package model

type Product struct {
	Id       int    `json:"ID"`
	Name     string `json:"Name"`
	Price    int    `json:"Price"`
	Category string `json:"Category"`
	Brand    string `json:"Brand"`
}
