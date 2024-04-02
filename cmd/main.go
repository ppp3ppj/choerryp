package main

import (
	"log"

	"github.com/ppp3ppj/choerryp/config"
	"github.com/ppp3ppj/choerryp/modules/servers"
	"github.com/ppp3ppj/choerryp/pkg/databases"
)

func main() {
    conf := config.ConfigGetting()
    db := databases.NewPostgresDatabase(conf.Database)

    defer func() {
        if err := db.Close(); err != nil {
            log.Fatalf("Failed to close database connection: %v", err)
        }
    }()

    server := servers.NewEchoServer(conf, db)
    server.Start()
}

