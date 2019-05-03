package main

import (
	"CanYouGetTo20_REST-API/pkg/database"
	"CanYouGetTo20_REST-API/pkg/server"
	"flag"
	"os"
)

func main() {
	port := flag.Int("port", 8000, "port to start server on")

	scoreRepo := database.NewScoreRepository(createDbConfig())
	s := server.NewServer(scoreRepo)

	s.Run(*port)
}

func createDbConfig() database.DbConfig {
	return database.DbConfig{
		os.Getenv("SERVER"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("PASSWD"),
		os.Getenv("DB"),
	}
}
