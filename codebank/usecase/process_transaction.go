package usecase

import (
	"github.com/codeedu/codebank/domain"
	"github.com/codeedu/codebank/dto"
)

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
}

func NewUseCaseTransaction(transactionRepository domain.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository}
}

func (u UseCaseTransaction) ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error) {

}
