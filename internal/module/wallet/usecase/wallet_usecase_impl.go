package usecase

import (
	"github.com/alfiqo/disbursement/internal/models"
	"github.com/alfiqo/disbursement/internal/module/wallet/domain"
	"github.com/gofiber/fiber/v2"
)

type WalletUsecaseImpl struct{}

func NewUsecaseWallet() WalletUsecase {
	return &WalletUsecaseImpl{}
}

func (uc *WalletUsecaseImpl) Disbursement(ctx *fiber.Ctx) error {
	var request domain.WalletRequest
	var wallet models.Wallet

	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	return ctx.JSON(wallet)
}
