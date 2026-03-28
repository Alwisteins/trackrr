package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"backend/internal/database"
	"backend/internal/modules/auth"
	"backend/internal/modules/client"
)

func main() {
	// Try multiple paths for .env
	envPaths := []string{
		".env",       // Current directory
		"../../.env", // If running from backend/cmd/api
		"../.env",    // If running from backend/
	}

	var loaded bool
	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			log.Println("✓ Loaded .env from:", path)
			loaded = true
			break
		}
	}

	if !loaded {
		log.Println("Warning: Could not load .env file from any path")
	}

	db, errDB := database.NewPostgres()
	if errDB != nil {
		log.Fatal(errDB)
	}

	// Initialize repositories, services, and handlers
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	clientRepo := client.NewRepository(db)
	clientService := client.NewService(clientRepo)
	clientHandler := client.NewHandler(clientService)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			authGroup := v1.Group("/auth")
			{
				authGroup.POST("/register", authHandler.Register)
				authGroup.POST("/login", authHandler.Login)
			}
			clientGroup := v1.Group("/clients")
			{
				clientGroup.POST("/", clientHandler.CreateClient)
				clientGroup.GET("/", clientHandler.FindAllClients)
				clientGroup.GET("/:uuid", clientHandler.FindClientByUUID)
				clientGroup.PUT("/:uuid", clientHandler.UpdateClient)
				clientGroup.DELETE("/:uuid", clientHandler.DeleteClient)
			}
		}
	}

	r.Run(":8080")
}
