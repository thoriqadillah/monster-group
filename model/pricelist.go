package model

type Pricelist struct {
	Id       int    `json:"ID"`
	Category string `json:"Category"`
	Brand    string `json:"Brand"`
	Price    int    `json:"Price"`
}
