package service

import (
	"context"
	"fmt"
	"wallet-api/internal/models"
)

type IWalletRepo interface {
	Deposit() (*models.Wallet, error)
	FindByIdAndLock(models.ID, func(context.Context, *models.Wallet) error) error
	Save(context.Context, *models.Wallet) error
}

type IDepositTransactionRepo interface {
	Save(context.Context, *models.DepositTransactionModel) error
}

type WalletService struct {
	walletRepo     IWalletRepo
	depositTrxRepo IDepositTransactionRepo
}

func NewWalletService(wRepo IWalletRepo, dTrxRepo IDepositTransactionRepo) *WalletService {
	return &WalletService{walletRepo: wRepo, depositTrxRepo: dTrxRepo}

}

func (s *WalletService) Deposit(userID models.ID, walletID models.ID, amount float64) error {
	return s.walletRepo.FindByIdAndLock(walletID, func(ctx context.Context, w *models.Wallet) error {
		if w.UserId != userID {
			return fmt.Errorf("")
		}
		w.AddAmount(amount)
		trx, err := models.NewSuccessDepositTransactionModel(userID, amount)
		if err != nil {
			return err
		}
		s.depositTrxRepo.Save(ctx, trx)
		return s.walletRepo.Save(ctx, w)

	})
}
