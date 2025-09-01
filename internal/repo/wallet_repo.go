package repo

import (
	"context"
	"wallet-api/internal/models"

	"gorm.io/gorm"
)

type WalletRepo struct {
	db gorm.DB
}

func (w *WalletRepo) FindByIdAndLock(ctx context.Context,
	id models.ID,
	c func(context.Context, *models.Wallet) error) error {

	return w.db.Transaction(func(tx *gorm.DB) error {

		var foundedWallet *models.Wallet
		err := tx.First(foundedWallet, id).Error
		if err != nil {
			return err
		}
		c(ctx, foundedWallet)

		return nil
	})

}
