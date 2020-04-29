package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gopkg.in/go-playground/validator.v9"
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
	category := models.ProductCategory{}

	if err := c.Bind(&category); err != nil {
		log.Error("Cannot bind body to product")
	}

	validate := validator.New()
	err := validate.Struct(category)
	if err != nil {
		log.Error("SCHEDULE: Bad request: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "BAD_REQUEST",
			"msg":    "Body invalid",
		})
	}

	result := dbs.CreateCategory(category)
	return c.JSON(http.StatusOK, result)
}

func UpdateCategoryByCode(c echo.Context) error {
	code := c.Param("code")
	category := models.ProductCategory{}

	if err := c.Bind(&category); err != nil {
		log.Error("Cannot bind body to product")
	}

	result := dbs.UpdateCategoryByCode(code, category)
	return c.JSON(http.StatusOK, result)
}

func DeleteCategoryByCode(c echo.Context) error {
	code := c.Param("code")
	dbs.DeleteCategoryByCode(code)

	return c.JSON(http.StatusOK, map[string]string{
		"status": "deleted",
	})
}
