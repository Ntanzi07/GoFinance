package handlers

import (
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
