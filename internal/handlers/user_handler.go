package handlers

import (
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

// CreateUserTransaction godoc
// @Summary create a new transaction for a user
// @Description create a new transaction associated with a specific user
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name path string true "User Name"
// @Param transaction body models.TransactionCreate true "Transaction Data"
// @Success 200 {object} models.Transaction
// @Failure 400 {string} string "Bad Request"
// @Failure 403 {string} string "Forbidden"
// @Failure 500 {string} string "Internal Server Error"
// @Router /{name} [post]
func (h *UserHandler) CreateUserTransaction(c *fiber.Ctx) error {
	name := c.Params("name")

	_, err := h.verifyJwt(c)
	if err != nil {
		return err
	}
	var transaction models.TransactionCreate
	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	if err := h.Repo.CreateUserTransaction(name, transaction.Type, transaction.Amount, transaction.Description, transaction.Date); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating transaction")
	}
	return c.JSON(transaction)
}

// DeleteUserTransaction godoc
// @Summary delete a user transaction
// @Description delete a specific transaction for a user
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name path string true "User Name"
// @Param id path int true "Transaction ID"
// @Success 200 {string} string "Transaction deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 403 {string} string "Forbidden"
// @Failure 500 {string} string "Internal Server Error"
// @Router /{name}/transactions/{id} [delete]
func (h *UserHandler) DeleteUserTransaction(c *fiber.Ctx) error {
	name := c.Params("name")

	transactionID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid transaction ID")
	}

	_, err = h.verifyJwt(c)
	if err != nil {
		return err
	}

	if err := h.Repo.DeleteUserTransaction(name, transactionID); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting transaction")
	}

	return c.SendString("Transaction deleted successfully")
}
