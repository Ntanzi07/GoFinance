package handlers

import (
	"github.com/Ntanzi07/gofinance/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type UsersHandler struct {
	Repo *repository.UsersRepository
}

func NewUsersHandler(repo *repository.UsersRepository) *UsersHandler {
	return &UsersHandler{Repo: repo}
}

func (h *UsersHandler) GetAllUserHandler(c *fiber.Ctx) error {
	users, err := h.Repo.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving users")
	}
	return c.JSON(users)
}
