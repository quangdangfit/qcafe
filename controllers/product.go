package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
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
	products := dbs.GetProductByCode(code)
	return c.JSON(http.StatusOK, products)
}

func CreateProduct(c echo.Context) error {
	products := models.Product{}

	if err := c.Bind(&products); err != nil {
		fmt.Println("Cannot bind body to product")
	}

	result := dbs.CreateProduct(products)
	return c.JSON(http.StatusOK, result)
}

func UpdateProductByCode(c echo.Context) error {
	code := c.Param("code")
	products := models.Product{}

	if err := c.Bind(&products); err != nil {
		fmt.Println("Cannot bind body to product")
	}

	result := dbs.UpdateProductByCode(code, products)
	return c.JSON(http.StatusOK, result)
}
