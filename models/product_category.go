package models

type ProductCategory struct {
	Code string `json:"code" bson:"code"`
	Name string `json:"name" bson:"name"`
}
