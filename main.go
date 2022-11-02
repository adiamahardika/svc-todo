package main

import (
	"fmt"
	"svc-todo/db"
	"svc-todo/router"
)

func main() {

	db, error := db.ConnectPostgres()
	if error != nil {
		fmt.Println("Connection to db has been error!")
	} else {
		fmt.Println("Connection to db succeed!")
	}

	router.Router(db)
}
