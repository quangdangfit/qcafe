package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"qcafe/dbs"
	"qcafe/models"
)

func GetProducts(c echo.Context) error {
	products := dbs.GetProducts()
	return c.JSON(http.StatusOK, products)
}

func GetProductByCode(c echo.Context) error {
	code := c.Param("code")
	product := dbs.GetProductByCode(code)
	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {
	product := models.Product{}

	if err := c.Bind(&product); err != nil {
		log.Error("Cannot bind body to product")
	}

	validate := validator.New()
	err := validate.Struct(product)
	if err != nil {
		log.Error("SCHEDULE: Bad request: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "BAD_REQUEST",
		})
	}

	result := dbs.CreateProduct(product)
	return c.JSON(http.StatusOK, result)
}

func UpdateProductByCode(c echo.Context) error {
	code := c.Param("code")
	products := models.Product{}

	if err := c.Bind(&products); err != nil {
		log.Error("Cannot bind body to product")
	}

	result := dbs.UpdateProductByCode(code, products)
	return c.JSON(http.StatusOK, result)
}

func DeleteProductByCode(c echo.Context) error {
	code := c.Param("code")
	dbs.DeleteProductByCode(code)

	return c.JSON(http.StatusOK, map[string]string{
		"status": "deleted",
	})
}
