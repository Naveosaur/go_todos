package main

import (
	"todo/database"

	"github.com/labstack/echo"
)

func main() {
	db := database.InitDB()
	 err := db.Ping()

	 if err != nil {
		panic(err)
	 }
	e := echo.New()

	e.Start(":8080")
}