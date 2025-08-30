package services

import (
	"context"
	"fmt"
	"wallet-api/internal/models"
)

type IWalletRepository interface {
	Deposit() (*models.Wallet, error)
	FindByIdAndLock(models.ID, func(context.Context, *models.Wallet) error) error
	Save(context.Context, *models.Wallet) error
}

type IDepositTransactionRepository interface {
	Save(context.Context, *models.DepositTransactionModel) error
}

type WalletService struct {
	walletRepository IWalletRepository
	depositTrxRepo   IDepositTransactionRepository
}

func NewWalletService(wRepo IWalletRepository, dTrxRepo IDepositTransactionRepository) *WalletService {
	return &WalletService{walletRepository: wRepo, depositTrxRepo: dTrxRepo}

}

func (s *WalletService) Deposit(userID models.ID, walletID models.ID, amount float64) error {
	return s.walletRepository.FindByIdAndLock(walletID, func(ctx context.Context, w *models.Wallet) error {
		if w.UserId != userID {
			return fmt.Errorf("")
		}
		w.AddAmount(amount)
		trx, err := models.NewSuccessDepositTransactionModel(userID, amount)
		if err != nil {
			return err
		}
		s.depositTrxRepo.Save(ctx, trx)
		return s.walletRepository.Save(ctx, w)

	})
}
