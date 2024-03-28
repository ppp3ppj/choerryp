package main

import (
	"fmt"

	"github.com/ppp3ppj/choerryp/internal/config"
	"github.com/ppp3ppj/choerryp/internal/databases"
	"github.com/ppp3ppj/choerryp/internal/server"
)

func main() {
    conf := config.ConfigGetting()
    db := databases.NewPostgresDatabase(conf.Database)
    fmt.Println(db)
    fmt.Println(conf)
    fmt.Println(conf.Database.Password)

    server := server.NewEchoServer(conf)
    server.Start()
}

