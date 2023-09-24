package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env.secret")
	if envErr != nil {
		log.Fatal(envErr)
	}
	postgres_user := os.Getenv("POSTGRES_USER")
	postgres_pass := os.Getenv("POSTGRES_PASS")
	postgres_db := os.Getenv("POSTGRES_DB")
	api_port := ":" + os.Getenv("API_PORT")
	repository, err := NewPostgresRepository("5432", "weatherdb", postgres_db, postgres_user, postgres_pass)
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(api_port, repository)
	server.Run()
}
