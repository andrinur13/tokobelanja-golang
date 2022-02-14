package service

import (
	"errors"
	"tokobelanja-golang/model/entity"
	"tokobelanja-golang/model/input"
	"tokobelanja-golang/repository"
)

type TransactionHistoryService interface {
	CreateTransaction(transactionInput input.InputTransaction, IDUser int) (entity.TransactionHistory, error)
	GetMyTransaction(IDUser int) ([]entity.TransactionHistory, error)
}

type transactionHistoryService struct {
	transactionHistoryRepository repository.TransactionHistoryRepository
	productRepository            repository.ProductRepository
	userRepository               repository.UserRepository
}

func NewTransactionHistoryService(transactionHistoryRepository repository.TransactionHistoryRepository, productRepository repository.ProductRepository, userRepository repository.UserRepository) *transactionHistoryService {
	return &transactionHistoryService{transactionHistoryRepository, productRepository, userRepository}
}

func (s *transactionHistoryService) CreateTransaction(transactionInput input.InputTransaction, IDUser int) (entity.TransactionHistory, error) {
	newTransactionHistory := entity.TransactionHistory{}

	newTransactionHistory.ProductID = transactionInput.ProductID
	newTransactionHistory.Quantity = transactionInput.Quantity

	// query product
	product, err := s.productRepository.GetProductByID(transactionInput.ProductID)

	if err != nil {
		return entity.TransactionHistory{}, err
	}

	if product.ID == 0 {
		return entity.TransactionHistory{}, err
	}

	// ketika jumlah tidak mencukupi
	if product.Stock < transactionInput.Quantity {
		return entity.TransactionHistory{}, errors.New("Jumlah tidak mencukupi")
	}

	// query user
	datauser, err := s.userRepository.GetByID(IDUser)

	if err != nil {
		return entity.TransactionHistory{}, err
	}

	if datauser.ID == 0 {
		return entity.TransactionHistory{}, err
	}

	// saldo tidak mencukupi
	if datauser.Balance < (product.Price * transactionInput.Quantity) {
		return entity.TransactionHistory{}, errors.New("Saldo tidak mencukupi")
	}

	// pastikan balance tersedia
	buyAmount := product.Price * transactionInput.Quantity
	newTransactionHistory.TotalPrice = buyAmount
	newTransactionHistory.UserID = IDUser

	// kurangi stock
	productUpdate := entity.Product{
		ID:    transactionInput.ProductID,
		Stock: product.Stock - transactionInput.Quantity,
	}

	_, err = s.productRepository.Update(productUpdate)

	if err != nil {
		return entity.TransactionHistory{}, err
	}

	// store data ke transactions history
	transactionCreated, err := s.transactionHistoryRepository.Save(newTransactionHistory)

	if err != nil {
		return entity.TransactionHistory{}, err
	}

	_, err = s.userRepository.UpdateSaldo(IDUser, buyAmount)

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
