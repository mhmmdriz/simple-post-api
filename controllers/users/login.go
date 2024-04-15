package users

import (
	"net/http"
	"soal-eksplorasi/configs"
	"soal-eksplorasi/middlewares"
	"soal-eksplorasi/models"
	"soal-eksplorasi/utils"

	"github.com/labstack/echo/v4"
)

func UserLogin(c echo.Context) error {
	var user models.User
	var userDB models.User

	c.Bind(&user)

	result := configs.DB.Where("email = ?", user.Email).First(&userDB).Error
	BaseResponseError := models.ErrorResponse{
		Status:    "error",
		Message:   "Email or password is wrong",
		ErrorCode: http.StatusUnauthorized,
	}
	if result != nil {
		return c.JSON(http.StatusUnauthorized, BaseResponseError)
	}

	// compare password
	if !utils.CheckPasswordHash(user.Password, userDB.Password) {
		return c.JSON(http.StatusUnauthorized, BaseResponseError)
	}

	// generate Token JWT
	var userResponse models.UserLogin
	userResponse.Name = userDB.Name
	userResponse.Email = userDB.Email
	userResponse.Id = userDB.Id
	userResponse.IsAdmin = userDB.IsAdmin
	userResponse.Token = middlewares.GenerateTokenJWT(userDB.Id, userDB.Name, userDB.IsAdmin)

	successResponse := models.SuccessResponse{
		Status:  "success",
		Message: "Success login",
		Data:    userResponse,
	}

	return c.JSON(http.StatusOK, successResponse)
}
