package users

import (
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
    Id string `db:"id" json:"id" form:"id"`
    Email string `db:"email" json:"email" form:"email"`
    Password string `db:"password" json:"password" form:"password"`
    Username string `db:"username" json:"username" form:"username"`
    Firstname string `db:"first_name" json:"first_name" form:"first_name"`
    Lastname string `db:"last_name" json:"last_name" form:"last_name"`
}

func (obj *UserRegisterReq) BcryptPassword() error {
    hashPassword, err := bcrypt.GenerateFromPassword([]byte(obj.Password), 10)
    if err != nil {
        return err
    }
    obj.Password = string(hashPassword)
    return nil
}
