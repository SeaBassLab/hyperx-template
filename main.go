package main

import (
	"log"
	"os"

	server "github.com/SeaBassLab/hyperx-server"
)

func main() {
	env := os.Getenv("HYPERX_ENV") // puede ser "dev" o "prod"
	if env == "" {
		env = "dev" // fallback por defecto
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "4001" // default si no está en env
	}
	err := server.StartServer(env, port)
	if err != nil {
		log.Fatal(err)
	}
}
