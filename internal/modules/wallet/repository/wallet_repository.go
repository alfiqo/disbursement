package repository

import (
	"github.com/alfiqo/disbursement/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type WalletRepository interface {
	Disbursement(ctx *fiber.Ctx, wallet *entity.Wallet) (*entity.Wallet, error)
}
