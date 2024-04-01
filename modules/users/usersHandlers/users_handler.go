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
    fmt.Println("signup")
    req := new(users.UserRegisterReq)
    req.Username = "pppp"
    req.Password = "123456"
    req.Firstname = "pppp"
    req.Lastname = "pppp"
    req.Email = "pppp"

    err := h.usersUsecase.InsertUser(req)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return ctx.JSON(http.StatusOK, "success")
}
