package repository

import (
	"github.com/alfiqo/disbursement/internal/entity"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: DB,
	}
}

func (r *UserRepositoryImpl) Update(ctx *fiber.Ctx, id uint, user *entity.User) (*entity.User, error) {

	return user, nil
}
