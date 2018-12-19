package main

import (
	db "ontbet/databases"
)

func main() {
	defer db.SqlDB.Close()

	router := initRouter()
	router.Run(":8000")
}
