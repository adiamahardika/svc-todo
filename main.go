package main

import (
	"fmt"
	"svc-todo/database"
	"svc-todo/router"
)

func main() {

	db, error := database.ConnectPostgres()
	database.MigrationPostgres(db)

	if error != nil {
		fmt.Println("Connection to db has been error!")
	} else {
		fmt.Println("Connection to db succeed!")
	}

	router.Router(db)
}
