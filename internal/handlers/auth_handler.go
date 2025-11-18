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

// LoginUserHandler godoc
// @Summary user login
// @Description verify user credentials and login, returning a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body models.UserCreds true "User Login Credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "JSON inválido"
// @Failure 401 {string} string "Senha inválida"
// @Failure 404 {string} string "Email não encontrado"
// @Router /login [post]
func (h *AuthHandler) LoginUserHandler(c *fiber.Ctx) error {
	var creds models.UserCreds

	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("JSON invalido!")
	}

	// Fetch user credentials (email, hashed password, isAdmin) from repository
	user, err := h.Repo.UserLogin(creds.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("email not founded")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid password")
	}

	// Create a JWT token with minimal claims (email and isAdmin).
	// Note: token expiration is commented out; consider adding `exp` claim.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   user.Email,
		"isAdmin": user.IsAdmin,
		// "exp": time.Now().Add(time.Hour * 1).Unix(), // 👈 Uncomment to add expiration (1 hour)
	})

	tokenString, err := token.SignedString(config.LoadJwt())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao gerar token"})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

// SingupUserHandler godoc
// @Summary user signup
// @Description signup for new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.UserSingUp true "New User Data"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "JSON inválido"
// @Failure 500 {string} string "User not created :/"
// @Router /signup [post]
func (h *AuthHandler) SingupUserHandler(c *fiber.Ctx) error {
	var user models.UserSingUp
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("JSON invalido!")
	}

	if err := h.Repo.CreateUser(user.Name, user.Email, user.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("User not created :/")
	}

	// After creating user, issue a token for the new account (not admin)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   user.Email,
		"isAdmin": false,
	})

	tokenString, err := token.SignedString(config.LoadJwt())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao gerar token"})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}
