package handlers

import (
	"github.com/Ntanzi07/gofinance/internal/config"
	"github.com/Ntanzi07/gofinance/internal/models"
	"github.com/Ntanzi07/gofinance/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Repo *repository.UsersRepository
}

func NewAuthHandler(repo *repository.UsersRepository) *AuthHandler {
	return &AuthHandler{Repo: repo}
}

// LoginUserHandler handles user login and JWT token generation.
func (h *AuthHandler) LoginUserHandler(c *fiber.Ctx) error {
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

func (h *AuthHandler) SingupUserHandler(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("JSON invalido!")
	}

	if err := h.Repo.CreateUser(user.Name, user.Email, user.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("User not created :/")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})

	tokenString, err := token.SignedString(config.LoadJwt())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao gerar token"})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}
