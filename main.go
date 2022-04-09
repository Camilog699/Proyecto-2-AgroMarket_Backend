package main

import (
	"fmt"
	"os"

	"habilitacion_backend/app"
	"habilitacion_backend/http/routes"
)

func main() {
	// AWS.Init()
	// fmt.Println(AWS.ListarBuckets())
	fmt.Println("AgroMarket back running")
	app.LoadEnv()
	server := app.NewServer()
	router := routes.CreateRouter()
	server.Initialize(os.Getenv("APP_URL"), router)
	server.Run()
}
