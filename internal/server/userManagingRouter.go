package server

import (
	"github.com/ppp3ppj/choerryp/pkg/users/usersHandlers"
	"github.com/ppp3ppj/choerryp/pkg/users/usersRepositories"
	"github.com/ppp3ppj/choerryp/pkg/users/usersUsecases"
)

func (s *echoServer) initUserManagingRouter() {
    router := s.app.Group("/users")
    userRepo := usersRepositories.UsersRepository(s.db.Connect())
    userUsecase := usersUsecases.UsersUsecase(userRepo)
    uH := usersHandlers.UsersHandler(s.app, userUsecase)

    router.GET("/signup", uH.Signup)


}
