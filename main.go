package main

import (
	"database/sql"
	"fmt"
	"log"
	"test-gin/api"
	db "test-gin/db/sqlc"
	"test-gin/util"
)

func main() {
	fmt.Println("Start processing..")
	defer fmt.Println("End processing.")

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
