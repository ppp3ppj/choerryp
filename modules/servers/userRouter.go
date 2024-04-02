package servers

import (
	"github.com/ppp3ppj/choerryp/modules/users/usersHandlers"
	"github.com/ppp3ppj/choerryp/modules/users/usersRepositories"
	"github.com/ppp3ppj/choerryp/modules/users/usersUsecases"
)

func (s *echoServer) initUserRouter() {
	userRepo := usersRepositories.UsersRepository(s.db)
	userUsecase := usersUsecases.UsersUsecase(userRepo)
    userHandler := usersHandlers.UsersHandler(s.app, userUsecase)

    router := s.app.Group("/users")

    router.POST("/signup", userHandler.Signup)
}
