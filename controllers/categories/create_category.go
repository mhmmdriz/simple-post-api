package categories

import (
	"net/http"
	"soal-eksplorasi/configs"
	"soal-eksplorasi/models"

	"github.com/labstack/echo/v4"
)

func CreateCategory(c echo.Context) error {
	category := models.Category{}
	c.Bind(&category)
	category.Posts = []models.Post{}

	err := configs.DB.Create(&category).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:    "error",
			Message:   err.Error(),
			ErrorCode: http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusCreated, models.SuccessResponse{
		Status:  "success",
		Message: "Success create new category",
		Data:    category,
	})
}
