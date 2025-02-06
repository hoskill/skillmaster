package main

import (
	"log"
	"skillmaster/internal/handler"
	"skillmaster/internal/repository"
	"skillmaster/internal/service"
	"skillmaster/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	//todo mv to env
	dsn := "host=localhost user=pguser password=pgpass dbname=skillmaster port=5432 sslmode=disable"
	db, err := database.InitDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	r := gin.Default()

	// todo init routes in func
	r.GET("/materials", handler.GetMaterials)
	r.GET("/materials/:id", handler.GetMaterial)
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	if err = r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
