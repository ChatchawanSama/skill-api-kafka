package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"skill-api-kafka/database"
	"skill-api-kafka/skill"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	database.ConnectDB()
	db := database.DB
	defer database.DB.Close()
	database.CreateTable()

	repo := skill.NewSkillRepo(db)

	skillHandler := skill.NewSkillHandler(repo)

	r := gin.Default()

	r.GET("/api/v1/skills", skillHandler.GetSkills)
	r.GET("/api/v1/skills/:key", skillHandler.GetSkillByKey)
	r.POST("/api/v1/skills", skillHandler.PostSkillByKey)
	// r.PUT("/api/v1/skills/:key", putSkillByKey)
	// r.DELETE("/api/v1/skills/:key", deleteSkillByKey)
	// r.PATCH("/api/v1/skills/:key/actions/name", patchSkillName)
	// r.PATCH("/api/v1/skills/:key/actions/description", patchSkillDescription)
	// r.PATCH("/api/v1/skills/:key/actions/logo", patchSkillLogo)
	// r.PATCH("/api/v1/skills/:key/actions/tags", patchSkillTags)

	port := os.Getenv("HOST")

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	serverErrors := make(chan error, 1)

	// Start the service listening for requests
	go func() {
		log.Printf("Listening on port %s", port)
		serverErrors <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		log.Println("Received shutdown signal, gracefully shutting down...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf("Graceful shutdown failed: %v", err)
		}

	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v", err)
		}
	}

	log.Println("Server stopped")
}
