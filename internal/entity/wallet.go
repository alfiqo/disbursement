package entity

import (
	"sync"

	"gorm.io/gorm"
)

type Wallet struct {
	RWMutext sync.RWMutex `gorm:"-"`
	gorm.Model
	UserID uint
	User   User
	Notes  string
	Debit  float64
	Credit float64
}

func (w *Wallet) AddCredit(amount float64) {
	w.RWMutext.Lock()
	w.Credit = w.Credit + amount
	w.RWMutext.Unlock()
}

func (w *Wallet) GetCredit() float64 {
	w.RWMutext.Lock()
	amount := w.Credit
	w.RWMutext.Unlock()

	return amount
}

func (w *Wallet) AddDebit(amount float64) {
	w.RWMutext.Lock()
	w.Debit = w.Debit + amount
	w.RWMutext.Unlock()
}

func (w *Wallet) GetDebit() float64 {
	w.RWMutext.Lock()
	amount := w.Debit
	w.RWMutext.Unlock()

	return amount
}
