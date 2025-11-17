package handlers

import (
	"github.com/Ntanzi07/gofinance/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TransactionHandler struct {
	Repo *repository.TransactionRepository
}

func NewTransactionHandler(repo *repository.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{Repo: repo}
}

// verifyJwt verifies the JWT token and checks if the user has permissio
func (h *TransactionHandler) verifyJwt(c *fiber.Ctx) error {

	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	isAdmin := claims["isAdmin"].(bool)

	if !isAdmin {
		return fiber.NewError(fiber.StatusForbidden, "Você não tem permissão")
	}

	return nil
}

// GetAllTransactionsHandler godoc
// @Summary get all transactions
// @Description retrieve all transaction records
// @Tags Transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Transaction
// @Failure 500 {string} string "Internal Server Error"
// @Router /transactions [get]
func (h *TransactionHandler) GetAllTransactionsHandler(c *fiber.Ctx) error {
	if err := h.verifyJwt(c); err != nil {
		return err
	}

	transactions, err := h.Repo.GetAllTransactions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(transactions)
}

// GetTransactionByIdHandler godoc
// @Summary get transaction by id
// @Description retrieve a specific transaction using its ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetTransactionByIdHandler(c *fiber.Ctx) error {
	if err := h.verifyJwt(c); err != nil {
		return err
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	transaction, err := h.Repo.GetTransactionByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(transaction)
}
