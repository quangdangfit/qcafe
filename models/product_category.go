package models

type ProductCategory struct {
	Code string `validate:"required";json:"code" bson:"code"`
	Name string `validate:"required";json:"name" bson:"name"`
}
