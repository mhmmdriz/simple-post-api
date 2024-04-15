package users

import (
	"net/http"
	"soal-eksplorasi/configs"
	"soal-eksplorasi/models"
	"soal-eksplorasi/utils"

	"github.com/labstack/echo/v4"
)

func UserRegister(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	// check if name, email, or password is empty
	if user.Name == "" || user.Email == "" || user.Password == "" {
		BaseResponseError := models.ErrorResponse{
			Status:    "error",
			Message:   "Name, email, or password cannot be empty",
			ErrorCode: 400,
		}
		return c.JSON(http.StatusBadRequest, BaseResponseError)
	}

	hash, _ := utils.HashPassword(user.Password)
	user.Password = hash

	result := configs.DB.Create(&user)

	if result.Error != nil {
		BaseResponseError := models.ErrorResponse{
			Status:    "error",
			Message:   "Failed to register user",
			ErrorCode: 500,
		}

		return c.JSON(http.StatusInternalServerError, BaseResponseError)
	}

	successResponse := models.SuccessResponse{
		Status:  "success",
		Message: "Success register user",
		Data:    user,
	}

	return c.JSON(http.StatusCreated, successResponse)

}
