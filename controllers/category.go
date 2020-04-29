package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"qcafe/dbs"
	"qcafe/models"
)

func GetCategories(c echo.Context) error {
	categories := dbs.GetCategories()
	return c.JSON(http.StatusOK, categories)
}

func GetCategoryByCode(c echo.Context) error {
	code := c.Param("code")
	category := dbs.GetCategoryByCode(code)
	return c.JSON(http.StatusOK, category)
}

func CreateCategory(c echo.Context) error {
	products := models.ProductCategory{}

	if err := c.Bind(&products); err != nil {
		log.Error("Cannot bind body to product")
	}

	result := dbs.CreateCategory(products)
	return c.JSON(http.StatusOK, result)
}

func UpdateCategoryByCode(c echo.Context) error {
	code := c.Param("code")
	products := models.ProductCategory{}

	if err := c.Bind(&products); err != nil {
		log.Error("Cannot bind body to product")
	}

	result := dbs.UpdateCategoryByCode(code, products)
	return c.JSON(http.StatusOK, result)
}

func DeleteCategoryByCode(c echo.Context) error {
	code := c.Param("code")
	dbs.DeleteCategoryByCode(code)

	return c.JSON(http.StatusOK, map[string]string{
		"status": "deleted",
	})
}
