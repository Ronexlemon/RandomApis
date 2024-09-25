package models

type Grocery struct{
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	Unit string `json:"unit"`
	Price int  `json:"price"`
}