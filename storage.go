package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	// UpdateAccountBalance(int, int) (*Account, error)
	GetAccounts() ([]*Account, error)
	GetAccountById(int) (*Account, error)
	GetAccountByAccountNumber(int) (*Account, error)
	TransferAccountBalance(int, int, int) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=golangbank sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		account_number SERIAL,
		encrypted_password VARCHAR(100),
		balance SERIAL,
		created_at TIMESTAMP
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `INSERT INTO account 
		(first_name, last_name, account_number, encrypted_password, balance, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.AccountNumber,
		acc.EncryptedPassword,
		acc.Balance,
		acc.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query("DELETE FROM account WHERE id = $1", id)

	return err
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := scanAccount(rows)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PostgresStore) GetAccountById(id int) (*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) GetAccountByAccountNumber(accNum int) (*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account WHERE account_number = $1", accNum)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanAccount(rows)
	}

	return nil, fmt.Errorf("account number [%d] not found", accNum)
}

func (s *PostgresStore) TransferAccountBalance(fromAccountId, toAccountNum, amount int) error {
	// start transaction
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		UPDATE account 
		SET balance = balance - $1 
		WHERE id = $2
		AND balance >= $1`, amount, fromAccountId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to deduct funds: %v", err)
	}

	result, err := tx.Exec(`
		UPDATE account 
		SET balance = balance + $1 
		WHERE account_number = $2`, amount, toAccountNum)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to add funds: %v", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("account number [%d] not found", toAccountNum)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

// helper funcs
func scanAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.AccountNumber,
		&account.EncryptedPassword,
		&account.Balance,
		&account.CreatedAt,
	)

	return account, err
}
