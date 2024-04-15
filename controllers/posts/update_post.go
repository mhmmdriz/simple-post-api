package posts

import (
	"net/http"
	"soal-eksplorasi/configs"
	"soal-eksplorasi/models"

	"github.com/labstack/echo/v4"
)

func UpdatePost(c echo.Context) error {
	id := c.Param("id")
	post := models.Post{}

	err := configs.DB.Where("id = ?", id).First(&post).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if post.ID == 0 {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Status:  "error",
			Message: "Post with ID " + id + " not found",
		})
	}

	c.Bind(&post)
	err = configs.DB.Save(&post).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Success edit post with ID " + id,
		Data:    post,
	})
}
