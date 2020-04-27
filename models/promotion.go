package models

type Promotion struct {
	Name    string  `json:"name" bson:"name"`
	Percent float64 `json:"percent" bson:"percent"`
}
