package http

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/gopkg/middleware"
	"github.com/phamtrung99/user-service/config"
	"github.com/phamtrung99/user-service/delivery/http/v1/user"
	"github.com/phamtrung99/user-service/delivery/http/v1/userfavorite"
	"github.com/phamtrung99/user-service/repository"
	"github.com/phamtrung99/user-service/usecase"
)

// NewHTTPHandler .
func NewHTTPHandler(repo *repository.Repository, ucase *usecase.UseCase) *echo.Echo {
	e := echo.New()
	cfg := config.GetConfig()

	skipper := func(c echo.Context) bool {
		p := c.Request().URL.Path

		return strings.Contains(p, "/health_check") ||
			strings.Contains(p, "/login")
	}

	e.Use(middleware.Auth(cfg.Jwt.Key, skipper, false))

	apiV1 := e.Group("/v1")

	user.Init(apiV1.Group("/users"), ucase)
	userfavorite.Init(apiV1.Group("/favorites"), ucase)

	return e
}
