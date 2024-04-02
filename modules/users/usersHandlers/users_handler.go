package usersHandlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/choerryp/modules/users"
	"github.com/ppp3ppj/choerryp/modules/users/usersUsecases"
)

type userHandlerErrCode string

const (
    signupErr userHandlerErrCode = "users-001"
)

type IUserHandler interface {
    Signup(ctx echo.Context) error
}

type userHandler struct {
    ctx *echo.Echo
    usersUsecase usersUsecases.IUserUsecase
}

func UsersHandler(ctx *echo.Echo, usersUsecase usersUsecases.IUserUsecase) IUserHandler {
    return &userHandler{
        ctx: ctx,
        usersUsecase: usersUsecase,
    }
}

func (h *userHandler) Signup(ctx echo.Context) error {
    req := new(users.UserRegisterReq)

    if err := ctx.Bind(req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("failed to bind request: %v", err))
    }

    err := h.usersUsecase.InsertUser(req)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return ctx.JSON(http.StatusOK, "success")
}
