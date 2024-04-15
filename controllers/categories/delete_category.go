package categories

import (
	"net/http"
	"soal-eksplorasi/configs"
	"soal-eksplorasi/models"

	"github.com/labstack/echo/v4"
)

func DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	category := models.Category{}

	err := configs.DB.Preload("Posts").First(&category, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if category.ID == 0 {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Status:  "error",
			Message: "Category not found",
		})
	}

	err = configs.DB.Delete(&category).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Success delete category with ID " + id,
		Data:    category,
	})
}
