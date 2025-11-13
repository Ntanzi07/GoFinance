package handlers

import (
	"github.com/Ntanzi07/gofinance/internal/config"
	"github.com/Ntanzi07/gofinance/internal/models"
	"github.com/Ntanzi07/gofinance/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func (h *UsersHandler) GetUserByIdHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	user, err := h.Repo.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving user")
	}

	return c.JSON(user)
}

func (h *UsersHandler) GetUserByNameHandler(c *fiber.Ctx) error {
	name := c.Params("name")

	user, err := h.Repo.GetUserByName(name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving user")
	}

	return c.JSON(user)
}

func (h *UsersHandler) CreateUserHandler(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	if err := h.Repo.CreateUser(user.Name, user.Email, user.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating user")
	}

	return c.SendString("User created successfully")
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

func (h *UsersHandler) LoginUserHandler(c *fiber.Ctx) error {
	var creds models.UserLogin

	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("JSON invalido!")
	}

	user, err := h.Repo.UserLogin(creds.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("email not founded")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		// "exp": time.Now().Add(time.Hour * 1).Unix(), // 👈 Descomente isso para adicionar expiração (1 hora)
	})

	tokenString, err := token.SignedString(config.LoadJwt())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao gerar token"})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}
