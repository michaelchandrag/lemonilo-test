package main

import (
	"os"
	"fmt"

	database "github.com/michaelchandrag/lemonilo-test/database"
	router "github.com/michaelchandrag/lemonilo-test/module/router"
)

func main () {
	err := database.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	fmt.Println("Database connected")

	r := router.SetupRouter()

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}