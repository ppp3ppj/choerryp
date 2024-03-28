package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/ppp3ppj/choerryp/internal/config"

    _ "github.com/jackc/pgx/v5/stdlib" // Standard library bindings for pgx
)

type postgresDatabase struct {
    *sqlx.DB
}

var (
    postgresDatabaseInstance *postgresDatabase
    once sync.Once
)

func NewPostgresDatabase(conf *config.Database) *postgresDatabase {
    once.Do(func() {
        dsn := fmt.Sprintf(
            "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
            conf.Host,
            conf.Port,
            conf.User,
            conf.Password,
            conf.DBName,
            conf.SSLMode,
        )
        fmt.Println(dsn)

        conn, err := sqlx.Connect("pgx", dsn)
        if err != nil {
            panic(err)
        }

        log.Printf("Connected to postgres database %s successfully", conf.DBName)

        postgresDatabaseInstance = &postgresDatabase{conn}
    })

    return postgresDatabaseInstance
}

func (db *postgresDatabase) Connect() *sqlx.DB {
    return db.DB
}