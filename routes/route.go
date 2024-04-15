package routes

import (
	"soal-eksplorasi/constants"
	"soal-eksplorasi/controllers/categories"
	"soal-eksplorasi/controllers/posts"
	"soal-eksplorasi/controllers/users"
	"soal-eksplorasi/middlewares"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// Middleware logger
	e.Use(middleware.Logger())
	// Middleware rate limiter. Limit to 20 request/second
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.POST("/api/v1/register", users.UserRegister)
	e.POST("/api/v1/login", users.UserLogin)

	jwtAuth := e.Group("")
	jwtAuth.Use(echojwt.JWT([]byte(constants.PRIVATE_KEY_JWT)))
	jwtAuth.GET("/api/v1/posts", posts.GetPosts)
	jwtAuth.GET("/api/v1/posts/:id", posts.GetPostByID)
	jwtAuth.POST("/api/v1/posts", posts.CreatePost)
	jwtAuth.PUT("/api/v1/posts/:id", posts.UpdatePost)
	jwtAuth.DELETE("/api/v1/posts/:id", posts.DeletePost)

	adminAuth := e.Group("")
	adminAuth.Use(echojwt.JWT([]byte(constants.PRIVATE_KEY_JWT)), middlewares.IsAdmin)
	adminAuth.GET("/api/v1/categories", categories.GetCategories)
	adminAuth.GET("/api/v1/categories/:id", categories.GetDetailCategory)
	adminAuth.POST("/api/v1/categories", categories.CreateCategory)
	adminAuth.DELETE("/api/v1/categories/:id", categories.DeleteCategory)

	return e
}
