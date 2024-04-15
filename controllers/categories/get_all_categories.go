package categories

import (
	"net/http"
	"soal-eksplorasi/configs"
	"soal-eksplorasi/models"

	"github.com/labstack/echo/v4"
)

func GetCategories(c echo.Context) error {
	categories := []models.CategoryOnly{}

	err := configs.DB.Table("Categories").Select("id", "name").Scan(&categories).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:    "error",
			Message:   err.Error(),
			ErrorCode: http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Success get all categories",
		Data:    categories,
	})
}
