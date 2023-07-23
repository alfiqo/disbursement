package domain

import (
	"time"

	"github.com/alfiqo/disbursement/internal/entity"
)

type WalletResponse struct {
	ID        uint      `json:"id"`
	Notes     string    `json:"notes"`
	Credit    float64   `json:"credit"`
	Debit     float64   `json:"debit"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCreateResponse(wallet *entity.Wallet) WalletResponse {
	return WalletResponse{
		ID:        wallet.ID,
		Notes:     wallet.Notes,
		Debit:     wallet.Debit,
		Credit:    wallet.Credit,
		CreatedAt: wallet.CreatedAt.UTC().Local(),
		UpdatedAt: wallet.UpdatedAt.Local(),
	}
}
