package posts

import (
	"net/http"
	"soal-eksplorasi/configs"
	"soal-eksplorasi/models"

	"github.com/labstack/echo/v4"
)

func GetPostByID(c echo.Context) error {
	id := c.Param("id")
	var post models.PostResponse

	err := configs.DB.Raw("SELECT posts.id, category_id, title, content, categories.name as category_name FROM posts INNER JOIN categories ON category_id = categories.id WHERE posts.id = ?", id).Scan(&post).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Success get post with ID " + id,
		Data:    post,
	})
}
