package categories

import (
	"net/http"
	"soal-eksplorasi/configs"
	"soal-eksplorasi/models"

	"github.com/labstack/echo/v4"
)

func GetDetailCategory(c echo.Context) error {
	id := c.Param("id")
	category := models.Category{}

	err := configs.DB.Preload("Posts").First(&category, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:    "error",
			Message:   err.Error(),
			ErrorCode: http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Success get detail category",
		Data:    category,
	})

}
