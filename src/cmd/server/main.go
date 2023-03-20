package main

import (
	http "github.com/eduardor2m/hexagonal-architecture/src/internal/adapters/http"
	"github.com/eduardor2m/hexagonal-architecture/src/internal/domain/repositories"
	"github.com/eduardor2m/hexagonal-architecture/src/internal/domain/usecases"
	"github.com/eduardor2m/hexagonal-architecture/src/internal/infrastructure/database"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := database.Connect()

	userRepo := repositories.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := http.NewUserHandler(userUseCase)

	u := e.Group("/users")

	u.POST("", userHandler.CreateUser)
	u.GET("/:email", userHandler.GetUser)

	e.Logger.Fatal(e.Start(":8080"))
}
