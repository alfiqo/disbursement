package repository

import (
	"github.com/alfiqo/disbursement/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	Update(ctx *fiber.Ctx, id uint, user *entity.User) (*entity.User, error)
}
