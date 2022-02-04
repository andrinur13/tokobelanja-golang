package service

import (
	"errors"
	"tokobelanja-golang/model/entity"
	"tokobelanja-golang/model/input"
	"tokobelanja-golang/repository"
)

type TransactionHistoryService interface {
	CreateTransaction(transactionInput input.InputTransaction) (entity.TransactionHistory, error)
	GetMyTransaction(IDUser int) ([]entity.TransactionHistory, error)
}

type transactionHistoryService struct {
	transactionHistoryRepository repository.TransactionHistoryRepository
}

func NewTransactionHistoryService(transactionHistoryRepository repository.TransactionHistoryRepository) *transactionHistoryService {
	return &transactionHistoryService{transactionHistoryRepository}
}

func (s *transactionHistoryService) CreateTransaction(transactionInput input.InputTransaction) (entity.TransactionHistory, error) {
	newTransactionHistory := entity.TransactionHistory{}

	newTransactionHistory.ProductID = transactionInput.ProductID
	newTransactionHistory.Quantity = transactionInput.Quantity

	transactionCreated, err := s.transactionHistoryRepository.Save(newTransactionHistory)

	if err != nil {
		return entity.TransactionHistory{}, err
	}

	if transactionCreated.ID == 0 {
		return entity.TransactionHistory{}, err
	}

	return transactionCreated, nil
}

func (s *transactionHistoryService) GetMyTransaction(IDUser int) ([]entity.TransactionHistory, error) {
	myTransaction, err := s.transactionHistoryRepository.GetTransactionByIDUser(IDUser)

	if err != nil {
		return []entity.TransactionHistory{}, err
	}

	if len(myTransaction) < 1 {
		return []entity.TransactionHistory{}, err
	}

	return myTransaction, nil
}

func (s *transactionHistoryService) GetUserTransaction(IDUser int, levelUser string) ([]entity.TransactionHistory, error) {

	if levelUser != "admin" {
		return []entity.TransactionHistory{}, errors.New("Unauthorized User")
	}

	myTransaction, err := s.transactionHistoryRepository.GetTransactionByIDUser(IDUser)

	if err != nil {
		return []entity.TransactionHistory{}, err
	}

	if len(myTransaction) < 1 {
		return []entity.TransactionHistory{}, err
	}

	return myTransaction, nil
}
