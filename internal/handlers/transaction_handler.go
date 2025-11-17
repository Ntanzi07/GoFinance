package handlers

import (
	"github.com/Ntanzi07/gofinance/internal/models"
	"github.com/Ntanzi07/gofinance/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	Repo *repository.TransactionRepository
}

func NewTransactionHandler(repo *repository.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{Repo: repo}
}

// GetAllTransactionsHandler godoc
// @Summary get all transactions
// @Description retrieve all transaction records
// @Tags Transactions
// @Accept json
// @Produce json
// @Success 200 {array} models.Transaction
// @Failure 500 {string} string "Internal Server Error"
// @Router /transactions [get]
func (h *TransactionHandler) GetAllTransactionsHandler(c *fiber.Ctx) error {
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
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetTransactionByIdHandler(c *fiber.Ctx) error {
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

// CreateTransactionHandler godoc
// @Summary create a new transaction
// @Description create a new transaction record
// @Tags Transactions
// @Accept json
// @Produce json
// @Param transaction body models.Transaction true "Transaction Data"
// @Success 200 {string} string "Transaction created successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /transactions [post]
func (h *TransactionHandler) CreateTransactionHandler(c *fiber.Ctx) error {
	var transaction models.Transaction

	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := h.Repo.CreateTransaction(
		transaction.UserID,
		transaction.Type,
		transaction.Amount,
		transaction.Description,
		transaction.Date,
	); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("Transaction created successfully")
}

// DeleteTransacionHandler godoc
// @Summary delete a transaction
// @Description delete a transaction record by its ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {string} string "Transaction deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /transactions/{id} [delete]
func (h *TransactionHandler) DeleteTransacionHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := h.Repo.DeleteTransaction(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("Transaction deleted successfully")
}
