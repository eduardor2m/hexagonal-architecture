package routes

import (
	"github.com/eduardor2m/hexagonal-architecture/src/internal/domain/models"
	"github.com/eduardor2m/hexagonal-architecture/src/internal/domain/usecases"
	"github.com/labstack/echo/v4"
)

func NewUserRoutes(e *echo.Echo, userUseCase *usecases.UserUseCase) {

	e.POST("/users", func(c echo.Context) error {
		user := models.User{}
		c.Bind(&user)

		err := userUseCase.CreateUser(user.Name, user.Email)

		if err != nil {
			return err
		}

		return c.JSON(200, map[string]string{
			"message": "User created",
		})
	})
}
