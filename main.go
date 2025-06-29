package main

import (
	"go_confess_space-project/config"
	"go_confess_space-project/config/websocket"
	"go_confess_space-project/model"
	"go_confess_space-project/router"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load it, skipping...")
	}

	db := config.DatabaseConnection()
	defer config.CloseDB()
	if db == nil {
		log.Fatal("Database connection failed")
	}

	if err := model.Migration(db); err != nil {
		log.Fatal("Database migration failed:", err)
	}

	hub := websocket.NewHub()
	go hub.Run()

	r := router.SetupRouter(hub)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server running on port", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}
