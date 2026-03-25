package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"backend/internal/database"
	"backend/internal/modules/auth"
)

func main() {
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			authGroup := v1.Group("/auth")
			{
				authGroup.POST("/login", authHandler.Login)
			}
		}
	}

	r.Run(":8080")
}
