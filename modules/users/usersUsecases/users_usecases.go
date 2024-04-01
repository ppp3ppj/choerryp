package usersUsecases

import (
	"github.com/ppp3ppj/choerryp/modules/users"
	"github.com/ppp3ppj/choerryp/modules/users/usersRepositories"
)

type IUserUsecase interface {
    InsertUser(u *users.UserRegisterReq) error
}

type userUsecase struct {
    userRepo usersRepositories.IUserRepository
}

func UsersUsecase(userRepo usersRepositories.IUserRepository) IUserUsecase {
    return &userUsecase{
        userRepo: userRepo,
    }
}

func (u *userUsecase) InsertUser(userRepo *users.UserRegisterReq) error {
    return u.userRepo.InsertUser(userRepo)
}
