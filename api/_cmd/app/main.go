package main

import (
	"log"
	"tgo/api/internal/config"
	"tgo/api/pkg/cluster"
	"tgo/api/pkg/http"
)

func main() {
	config.LoadEnv()
	cluster.SetMaxProcs()

	app := http.Setup()

	port := config.GetEnv("APP_PORT", "3333")

	log.Fatal(app.Listen(":" + port))
}
