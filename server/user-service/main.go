package main

import (
	"fmt"
	"net/http"
	"user-service/config"
	"user-service/handler"
	"user-service/model"
	"user-service/repository"
	"user-service/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})

	repo := repository.NewUserRepository(db)
	service := service.NewAuthService(repo)
	handler := handler.NewAuthHandler(service)

	r := gin.Default()

	api := r.Group("/user")
	{
		api.GET("/health", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "User service is up!")
		})

		api.POST("/login", handler.Login)

		api.POST("/register", handler.Register)
	}

	addr := fmt.Sprintf(":%s", cfg.Port)
	r.Run(addr)
}
