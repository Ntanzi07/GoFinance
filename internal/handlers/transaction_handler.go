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

func (h *TransactionHandler) GetAllTransactionsHandler(c *fiber.Ctx) error {
	transactions, err := h.Repo.GetAllTransactions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching transactions")
	}
	return c.JSON(transactions)
}

func (h *TransactionHandler) GetTransactionByIdHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid transaction ID")
	}

	transaction, err := h.Repo.GetTransactionByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching transaction")
	}

	return c.JSON(transaction)
}

func (h *TransactionHandler) CreateTransactionHandler(c *fiber.Ctx) error {
	var transaction models.Transaction
	err := c.BodyParser(&transaction)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	if err := h.Repo.CreateTransaction(
		transaction.ID, transaction.Type,
		transaction.Amount,
		transaction.Description,
		transaction.Date,
	); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating Transaction")
	}

	return c.SendString("Transaction created successfully")
}

func (h *TransactionHandler) DeleteTransacionHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	if err := h.Repo.DeleteTransaction(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting Transaction")
	}

	return c.SendString("Transaction deleted successfully")
}
