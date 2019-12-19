package main

import (
	"echo-api-starter/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func init()  {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {

	e :=routes.New()

	e.Logger.Fatal(e.Start(":1323"))
}
