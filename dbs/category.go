package dbs

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2/bson"
	"qcafe/models"
)

func GetCategories() []models.ProductCategory {
	sessionClone := MongoSession.Copy()
	defer sessionClone.Close()
	collection := sessionClone.DB("qcafe").C("category")

	var result []models.ProductCategory
	if err := collection.Find(nil).All(&result); err != nil {
		log.Error(err)
	}

	return result
}

func GetCategoryByCode(code string) models.ProductCategory {
	sessionClone := MongoSession.Copy()
	defer sessionClone.Close()
	collection := sessionClone.DB("qcafe").C("category")

	var result models.ProductCategory
	selector := bson.M{"code": code}
	if err := collection.Find(selector).One(&result); err != nil {
		log.Error(err)
	}

	return result
}

func CreateCategory(data models.ProductCategory) models.ProductCategory {
	sessionClone := MongoSession.Copy()
	defer sessionClone.Close()
	collection := sessionClone.DB("qcafe").C("category")

	result := models.ProductCategory{}
	copier.Copy(&result, &data)

	var err = collection.Insert(result)

	if err != nil {
		log.Error(err)
	}

	return result
}

func UpdateCategoryByCode(code string, data models.ProductCategory) models.ProductCategory {
	sessionClone := MongoSession.Copy()
	defer sessionClone.Close()
	collection := sessionClone.DB("qcafe").C("category")

	result := models.ProductCategory{}
	copier.Copy(&result, &data)
	selector := bson.M{"code": code}
	var err = collection.Update(selector, &result)

	if err != nil {
		log.Error(err)
	}

	return result
}

func DeleteCategoryByCode(code string) {
	sessionClone := MongoSession.Copy()
	defer sessionClone.Close()
	collection := sessionClone.DB("qcafe").C("category")

	selector := bson.M{"code": code}
	var err = collection.Remove(selector)

	if err != nil {
		log.Error(err)
	}
}
