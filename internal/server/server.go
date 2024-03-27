package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/ppp3ppj/choerryp/internal/config"
)


type echoServer struct {
    app *echo.Echo
    conf *config.Config
}

var (
    server *echoServer
    once  sync.Once
)

func NewEchoServer(conf *config.Config) *echoServer {
    echoApp := echo.New()
    echoApp.Logger.SetLevel(log.DEBUG)

    once.Do(func() {
        server = &echoServer{
            app: echo.New(),
            conf: config.ConfigGetting(),
        }
    })

    return server
}

func (s *echoServer) Start() {
    s.app.Use(middleware.Recover())
    s.app.Use(middleware.Logger())


    s.app.GET("/v1/health", s.healthCheck)

    s.httpListening()
}

func (s *echoServer) httpListening() {
    url := fmt.Sprintf(":%d", s.conf.Server.Port)
    if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
        s.app.Logger.Fatalf("shutting down the server: %v", err)
    }
}

func (s *echoServer) healthCheck(c echo.Context) error {
    return c.String(http.StatusOK, "OK")
}