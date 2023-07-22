package routes

import (
	"github.com/alfiqo/disbursement/exceptions"
	"github.com/alfiqo/disbursement/internal/module/wallet/usecase"
	"github.com/gofiber/fiber/v2"
)

func NewResthandler() *fiber.App {
	router := fiber.New(fiber.Config{
		ErrorHandler: exceptions.ErorrHandler,
	})

	walletUsecase := usecase.NewUsecaseWallet()
	router.Post("/disbursement", walletUsecase.Disbursement)

	return router
}
