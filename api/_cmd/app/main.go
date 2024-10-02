package main

import (
	"log"
	"tgo/api/internal/config"
	"tgo/api/internal/modules/queue"
	"tgo/api/pkg/cluster"
	"tgo/api/pkg/http"
)

func main() {
	config.LoadEnv()
	cluster.SetMaxProcs()

	rabbitAdapter, err := queue.NewRabbitMQAdapter()
	if err != nil {
		log.Fatalf("Erro ao conectar ao RabbitMQ: %v", err)
	}

	defer rabbitAdapter.Close()

	app := http.Setup(rabbitAdapter)

	port := config.GetEnv("APP_PORT", "3333")

	log.Fatal(app.Listen(":" + port))
}
