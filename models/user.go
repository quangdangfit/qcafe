package models

type User struct {
	UID   int    `json:"uid" bson:"uid"`
	Name  string `json:"name" bson:"name"`
	Phone string `json:"phone" bson:"phone"`
}
