package main

import (
	"github.com/eduardor2m/hexagonal-architecture/src/internal/domain/repositories"
	"github.com/eduardor2m/hexagonal-architecture/src/internal/domain/usecases"
	"github.com/eduardor2m/hexagonal-architecture/src/internal/infrastructure/database"
	"github.com/eduardor2m/hexagonal-architecture/src/internal/infrastructure/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := database.Connect()

	userRepo := repositories.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepo)

	routes.NewUserRoutes(e, userUseCase)

	e.Logger.Fatal(e.Start(":8080"))
}
