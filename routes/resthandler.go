package routes

import (
	"github.com/alfiqo/disbursement/exceptions"
	walletRepo "github.com/alfiqo/disbursement/internal/modules/wallet/repository"
	"github.com/alfiqo/disbursement/internal/modules/wallet/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewResthandler(DB *gorm.DB) *fiber.App {
	router := fiber.New(fiber.Config{
		ErrorHandler: exceptions.ErorrHandler,
	})

	walletRepository := walletRepo.NewWalletRepository(DB)
	walletUsecase := usecase.NewUsecaseWallet(walletRepository)
	router.Post("/disbursement", walletUsecase.Disbursement)
	router.Post("/topup", walletUsecase.Topup)

	return router
}
