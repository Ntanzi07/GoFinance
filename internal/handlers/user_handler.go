package handlers

import (
	"fmt"

	"github.com/Ntanzi07/gofinance/internal/models"
	"github.com/Ntanzi07/gofinance/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	Repo *repository.UsersRepository
}

func NewUsersHandler(repo *repository.UsersRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

// verifyJwt verifies the JWT token and checks if the user has permissio
func (h *UserHandler) verifyJwt(c *fiber.Ctx) (models.User, error) {
	name := c.Params("name")

	user, err := h.Repo.GetUserByName(name)
	if err != nil {
		return models.User{}, fiber.NewError(fiber.StatusInternalServerError, "Erro ao buscar usuário")
	}

	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	isAdmin := claims["isAdmin"].(bool)

	fmt.Println(user.Email != email, !isAdmin)
	fmt.Println(user.Email, email)

	if user.Email != email && !isAdmin {
		return models.User{}, fiber.NewError(fiber.StatusForbidden, "Você não tem permissão")
	}

	return user, nil
}

// GetUserByNameHandler godoc
// @Summary get user by name
// @Description retrieve user details using their name
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name path string true "User Name"
// @Success 200 {object} models.User
// @Failure 403 {string} string "Forbidden"
// @Failure 500 {string} string "Internal Server Error"
// @Router /{name}/infos [get]
func (h *UserHandler) GetUserByNameHandler(c *fiber.Ctx) error {
	user, err := h.verifyJwt(c)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

// GetUserTransactions godoc
// @Summary get user transactions
// @Description retrieve all transactions for a specific user
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name path string true "User Name"
// @Success 200 {array} models.Transaction
// @Failure 403 {string} string "Forbidden"
// @Failure 500 {string} string "Internal Server Error"
// @Router /{name} [get]
func (h *UserHandler) GetUserTransactions(c *fiber.Ctx) error {
	name := c.Params("name")

	_, err := h.verifyJwt(c)
	if err != nil {
		return err
	}

	transactions, err := h.Repo.GetAllUserTransactions(name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving transactions")
	}

	return c.JSON(transactions)
}

/*
func (h *UsersHandler) GetAllUserHandler(c *fiber.Ctx) error {
	users, err := h.Repo.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving users")
	}
	return c.JSON(users)
}


func (h *UsersHandler) DeleteUserHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	if err := h.Repo.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting user")
	}

	return c.SendString("User deleted successfully")
}
*/
