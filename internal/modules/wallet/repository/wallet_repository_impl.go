package repository

import (
	"errors"

	"github.com/alfiqo/disbursement/internal/entity"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type WalletRepositoryImpl struct {
	DB *gorm.DB
}

func NewWalletRepository(DB *gorm.DB) WalletRepository {
	return &WalletRepositoryImpl{
		DB: DB,
	}
}

func (r *WalletRepositoryImpl) Topup(ctx *fiber.Ctx, wallet *entity.Wallet) (*entity.Wallet, error) {
	var user entity.User
	var transaction entity.Wallet

	eg, _ := errgroup.WithContext(ctx.Context())

	eg.Go(func() error {
		err := r.DB.First(&user, wallet.UserID).Error
		return err
	})
	eg.Go(func() error {
		err := r.DB.Table("wallets").Select("sum(debit) debit, sum(credit) credit").Where("user_id = ?", wallet.UserID).Find(&transaction).Error
		return err
	})
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	ok := isBalance(&transaction, user.GetBalance())
	if !ok {
		return nil, errors.New("cant topup, please contact our administrator")
	}

	user.AddBalance(wallet.GetDebit())

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&wallet).Error; err != nil {
			return err
		}

		if err := tx.Save(&user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (r *WalletRepositoryImpl) Disbursement(ctx *fiber.Ctx, wallet *entity.Wallet) (*entity.Wallet, error) {
	var user entity.User
	var transaction entity.Wallet

	eg, _ := errgroup.WithContext(ctx.Context())

	eg.Go(func() error {
		err := r.DB.First(&user, wallet.UserID).Error
		return err
	})
	eg.Go(func() error {
		err := r.DB.Table("wallets").Select("sum(debit) debit, sum(credit) credit").Where("user_id = ?", wallet.UserID).Find(&transaction).Error
		return err
	})
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	ok := isBalance(&transaction, user.GetBalance())
	if !ok {
		return nil, errors.New("cant disbursed, please contact our administrator")
	}

	user.DeductBalance(wallet.GetCredit())
	if user.GetBalance() < 0 {
		return nil, errors.New("your balance is insufficient")
	}

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&wallet).Error; err != nil {
			return err
		}

		if err := tx.Save(&user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func isBalance(trx *entity.Wallet, balance float64) bool {
	return (trx.GetDebit() - trx.GetCredit()) == balance
}
