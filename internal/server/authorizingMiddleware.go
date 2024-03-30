package server

import "github.com/labstack/echo/v4"

type authorizingMiddleware struct {
    logger echo.Logger
}

func (m *authorizingMiddleware) UserAuthorizing(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        m.logger.Info("Authorizing user")
        return next(c)
    }
}
