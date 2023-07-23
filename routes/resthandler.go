package routes

import (
	"github.com/alfiqo/disbursement/exceptions"
	userRepo "github.com/alfiqo/disbursement/internal/modules/user/repository"
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
	// router.Use(basicauth.New(basicauth.Config{
	// 	Users: map[string]string{
	// 		"alfiqo": "password",
	// 	},
	// }))
	router.Post("/disbursement", walletUsecase.Disbursement)
	router.Post("/topup", walletUsecase.Topup)

	userRepo.NewUserRepository(DB)

	return router
}
