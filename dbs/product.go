package dbs

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2/bson"
	"qcafe/models"
)

func GetProducts() []models.Product {
	sessionClone := MongoSession.Copy()
	defer sessionClone.Close()
	collection := sessionClone.DB("qcafe").C("product")

	var result []models.Product
	if err := collection.Find(nil).All(&result); err != nil {
		log.Error(err)
	}

	return result
}

func GetProductByCode(code string) models.Product {
	sessionClone := MongoSession.Copy()
	defer sessionClone.Close()
	collection := sessionClone.DB("qcafe").C("product")

	var result models.Product
	selector := bson.M{"code": code}
	if err := collection.Find(selector).One(&result); err != nil {
		log.Error(err)
	}

	return result
}

func CreateProduct(data models.Product) models.Product {
	sessionClone := MongoSession.Copy()
	defer sessionClone.Close()
	collection := sessionClone.DB("qcafe").C("product")

	result := models.Product{}
	copier.Copy(&result, &data)

	result.ComputeRelatedFields()
	var err = collection.Insert(result)

	if err != nil {
		log.Error(err)
	}

	return result
}
func UpdateProductByCode(code string, data models.Product) models.Product {
	sessionClone := MongoSession.Copy()
	defer sessionClone.Close()
	collection := sessionClone.DB("qcafe").C("product")

	result := models.Product{}
	copier.Copy(&result, &data)
	selector := bson.M{"code": code}
	result.ComputeRelatedFields()
	var err = collection.Update(selector, &result)

	if err != nil {
		log.Error(err)
	}

	return result
}
