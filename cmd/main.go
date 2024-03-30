package main

import (
	"log"

	"github.com/ppp3ppj/choerryp/internal/config"
	"github.com/ppp3ppj/choerryp/internal/databases"
	"github.com/ppp3ppj/choerryp/internal/server"
)

func main() {
    conf := config.ConfigGetting()
    db := databases.NewPostgresDatabase(conf.Database)
    defer func() {
        if err := db.Close(); err != nil {
            log.Fatalf("Failed to close database connection: %v", err)
        }
    }()

    server := server.NewEchoServer(conf, db)
    server.Start()
}

