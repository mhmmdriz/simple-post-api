package posts

import (
	"net/http"
	"soal-eksplorasi/configs"
	"soal-eksplorasi/models"

	"github.com/labstack/echo/v4"
)

func CreatePost(c echo.Context) error {
	post := models.Post{}
	c.Bind(&post)

	err := configs.DB.Create(&post).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.SuccessResponse{
		Status:  "success",
		Message: "Success add post",
		Data:    post,
	})
}
