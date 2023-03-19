package http

import (
	"github.com/eduardor2m/hexagonal-architecture/src/internal/domain/models"
	"github.com/eduardor2m/hexagonal-architecture/src/internal/domain/usecases"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUseCase usecases.UserUseCase
}

func NewUserHandler(userUseCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: userUseCase,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(400, err)
	}

	if err := h.UserUseCase.CreateUser(user.Email, user.Name); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(201, user)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	email := c.Param("email")

	user, err := h.UserUseCase.GetUserByEmail(email)

	if err != nil {
		return c.JSON(404, err)
	}

	return c.JSON(200, user)
}
