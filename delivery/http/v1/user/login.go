package user

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/auth-service/proto"
	"github.com/phamtrung99/gopkg/apperror"
	"github.com/phamtrung99/gopkg/utils"
	authservice "github.com/phamtrung99/user-service/client/authService"
)

type LoginRequest struct {
	Email    string
	Password string
}

func (r *Route) Login(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
		req      = LoginRequest{}
	)

	// Bind order by
	if err := c.Bind(&req); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	reqLogin := &proto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	client := authservice.GetClient()

	res, err := client.Login(ctx, reqLogin)

	if err != nil {
		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	return utils.Response.Success(ctx, res)
}
