package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"soal-eksplorasi/models"
	"strings"

	"github.com/labstack/echo/v4"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")
		splittedAuth := strings.Split(authorization, "Bearer ")
		token := splittedAuth[1]
		splittedToken := strings.Split(token, ".")
		payload := splittedToken[1]
		decodedPayload, _ := base64.RawURLEncoding.DecodeString(payload)

		var jsonData map[string]interface{}

		// unmarshal decodedPayload ke dalam jsonData
		err := json.Unmarshal(decodedPayload, &jsonData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Status:    "error",
				Message:   err.Error(),
				ErrorCode: http.StatusInternalServerError,
			})
		}

		if jsonData["isAdmin"] == false {
			return c.JSON(http.StatusForbidden, models.ErrorResponse{
				Status:    "error",
				Message:   "You are not authorized to access this endpoint",
				ErrorCode: http.StatusForbidden,
			})
		}

		return next(c)
	}

}
