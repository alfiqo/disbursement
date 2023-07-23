package usecase

import (
	"errors"

	"github.com/alfiqo/disbursement/helper"
	"github.com/alfiqo/disbursement/helper/static"
	"github.com/alfiqo/disbursement/internal/entity"
	"github.com/alfiqo/disbursement/internal/modules/wallet/domain"
	"github.com/alfiqo/disbursement/internal/modules/wallet/repository"
	"github.com/gofiber/fiber/v2"
)

type WalletUsecaseImpl struct {
	Repository repository.WalletRepository
}

func NewUsecaseWallet(walletRepository repository.WalletRepository) WalletUsecase {
	return &WalletUsecaseImpl{
		Repository: walletRepository,
	}
}

func (uc *WalletUsecaseImpl) Disbursement(ctx *fiber.Ctx) error {
	var request domain.WalletRequest

	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	if request.Amount < 0 {
		return errors.New("amount is invalid")
	}

	var wallet entity.Wallet
	wallet.AddCredit(request.Amount)
	wallet.UserID = 1
	wallet.Notes = static.Disbursed

	data, err := uc.Repository.Disbursement(ctx, &wallet)
	if err != nil {
		return err
	}

	res := domain.NewCreateResponse(data)

	return ctx.JSON(helper.GlobalResponse(fiber.StatusOK, "OK", res))
}
