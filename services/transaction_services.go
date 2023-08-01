package services

import (
	"database/sql"
	"transactions/models"
)

type TransactionService struct {
	db *sql.DB
}

func NewTransactionService(db *sql.DB) *TransactionService {
	return &TransactionService{db: db}
}

func (s *TransactionService) Insert(amount float64, category string) (int64, error) {
	result, err := s.db.Exec("INSERT INTO transactions (amount, category) VALUES (?, ?)", amount, category)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (s *TransactionService) GetByID(id int64) (*models.Transaction, error) {
	row := s.db.QueryRow("SELECT id, amount, category FROM transactions WHERE id = ?", id)
	transaction := &models.Transaction{}
	err := row.Scan(&transaction.ID, &transaction.Amount, &transaction.Category)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *TransactionService) Update(id int64, amount float64, category string) error {
	_, err := s.db.Exec("UPDATE transactions SET amount = ?, category = ? WHERE id = ?", amount, category, id)
	return err
}

func (s *TransactionService) GetAll() ([]models.Transaction, error) {
	rows, err := s.db.Query("SELECT id, amount, category FROM transactions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []models.Transaction{}
	for rows.Next() {
		transaction := models.Transaction{}
		err := rows.Scan(&transaction.ID, &transaction.Amount, &transaction.Category)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (s *TransactionService) Delete(id int64) error {
	_, err := s.db.Exec("DELETE FROM transactions WHERE id = ?", id)
	return err
}
