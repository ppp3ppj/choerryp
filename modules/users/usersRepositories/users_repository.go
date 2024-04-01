package usersRepositories

import (
	"context"
	"fmt"
	"time"

	"github.com/ppp3ppj/choerryp/modules/users"
	"github.com/ppp3ppj/choerryp/pkg/databases"
)

type IUserRepository interface {
    InsertUser(u *users.UserRegisterReq) error
}

type userRepository struct {
    db databases.Database
}

func UsersRepository(db databases.Database) IUserRepository {
    return &userRepository{
        db: db,
    }
}

func (r *userRepository) InsertUser(u *users.UserRegisterReq) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := u.BcryptPassword(); err != nil {
        return err
    }

    query := `INSERT INTO users (email, password, username, first_name, last_name) VALUES ($1, $2, $3, $4, $5) RETURNING "id"`
    if err := r.db.Connect().QueryRowContext(
        ctx,
        query,
        u.Email,
        u.Password,
        u.Username,
        u.Firstname,
        u.Lastname,
    ).Scan(&u.Id); err != nil { 
      switch err.Error() {
            case "ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)":
                return fmt.Errorf("username has been used")
            case "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)":
                return fmt.Errorf("email has been used")
            default:
                return fmt.Errorf("insert user failed: %v", err)
        }
    }

    return nil
}
