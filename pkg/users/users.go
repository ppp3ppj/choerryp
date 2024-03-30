package users

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
    Id string
    Email string
    Password string
    Username string
    Firstname string
    Lastname string
}

type UserRegisterReq struct {
    Id string `db:"id"`
    Email string `db:"email"`
    Password string `db:"password"`
    Username string `db:"username"`
    Firstname string `db:"first_name"`
    Lastname string `db:"last_name"`
}

func (obj *UserRegisterReq) BcryptPassword() error {
    hashPassword, err := bcrypt.GenerateFromPassword([]byte(obj.Password), 10)
    if err != nil {
        return err
    }
    obj.Password = string(hashPassword)
    return nil
}
