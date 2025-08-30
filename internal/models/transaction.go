package models

type DepositTransactionModel struct {
	BaseModel
	UserId ID
	Amount float64
	Status DepositTransactionStatus
}

func newDepositTransactionModel(userId ID, amount float64, status DepositTransactionStatus) (*DepositTransactionModel, error) {
	return &DepositTransactionModel{
		UserId: userId,
		Amount: amount,
		Status: status,
	}, nil
}

func NewSuccessDepositTransactionModel(userId ID, amount float64) (*DepositTransactionModel, error) {
	return newDepositTransactionModel(userId, amount, DepositTransactionSuccess)
}

type DepositTransactionStatus string

const (
	DepositTransactionSuccess DepositTransactionStatus = "success"
	DepositTransactionFailed  DepositTransactionStatus = "failed"
)
