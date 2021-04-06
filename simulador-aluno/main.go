package main

import (
	"fmt"
	route2 "github.com/codeedu/imersaofsfc2-simulator/aplication/route"
	"github/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	route := route2.Route{
		Id: "1",
		ClientId: "1",
	}
	route.LoadPositions()
	stringjson,_ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}