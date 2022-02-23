package main

import (
	"github.com/zul9an9/webapi-with-go/database"
	"github.com/zul9an9/webapi-with-go/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
