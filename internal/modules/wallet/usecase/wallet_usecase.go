package usecase

import "github.com/gofiber/fiber/v2"

type WalletUsecase interface {
	Disbursement(ctx *fiber.Ctx) error
	Topup(ctx *fiber.Ctx) error
}
