package entity

import (
	"sync"

	"gorm.io/gorm"
)

type User struct {
	RWMutext sync.RWMutex `gorm:"-"`
	gorm.Model
	Username string
	Password string
	Balance  float64
}

func (u *User) AddBalance(amount float64) {
	u.RWMutext.Lock()
	u.Balance = u.Balance + amount
	u.RWMutext.Unlock()
}

func (u *User) GetBalance() float64 {
	u.RWMutext.Lock()
	balance := u.Balance
	u.RWMutext.Unlock()

	return balance
}

func (u *User) DeductBalance(amount float64) {
	u.RWMutext.Lock()
	u.Balance = u.Balance - amount
	u.RWMutext.Unlock()
}
