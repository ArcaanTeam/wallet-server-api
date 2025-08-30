package models

import "gorm.io/gorm"

type Wallet struct {
	*gorm.Model
	UserId ID
	Amount float64
}

func (w *Wallet) AddAmount(amount float64) {
	w.Amount += amount
}
