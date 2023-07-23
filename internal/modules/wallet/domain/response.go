package domain

import (
	"time"

	"github.com/alfiqo/disbursement/internal/entity"
)

type WalletResponse struct {
	ID        uint    `json:"id"`
	Notes     string  `json:"notes"`
	Credit    float64 `json:"credit"`
	Debit     float64 `json:"debit"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

func NewCreateResponse(wallet *entity.Wallet) WalletResponse {

	return WalletResponse{
		ID:        wallet.ID,
		Notes:     wallet.Notes,
		Debit:     wallet.Debit,
		Credit:    wallet.Credit,
		Balance:   wallet.User.GetBalance(),
		CreatedAt: wallet.CreatedAt.Format(time.DateTime),
		UpdatedAt: wallet.UpdatedAt.Format(time.DateTime),
	}
}
