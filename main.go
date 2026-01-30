package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Raghunandan-79/auth-service/database"
	"github.com/Raghunandan-79/auth-service/models"
	"github.com/Raghunandan-79/auth-service/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	database.DB.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
	)

	routes.RegisterRoutes(r)

	// Custom server
	srv := &http.Server{
		Addr: "localhost:8082",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen %v\n", err)
		}
	}()

	log.Println("Server started on localhost:8082")

	// listen for SIGINT / SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server")

	// Graceful shutdown (max wait 5s)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}


}