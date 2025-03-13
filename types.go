package main

import (
	"math/rand"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	AccountNumber int64  `json:"accountNumber"`
	Password      string `json:"password"`
}

type LoginResponse struct {
	AccountNumber int64  `json:"accountNumber"`
	Token         string `json:"token"`
}

type BankJWTClaims struct {
	AccountNumber int64 `json:"accountNumber"`
	jwt.RegisteredClaims
}

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	AccountNumber     int64     `json:"accountNumber"`
	EncryptedPassword string    `json:"-"`
	Balance           int64     `json:"balance"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (acc *Account) ValidatePassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(acc.EncryptedPassword), []byte(pw)) == nil
}

func NewAccount(firstName, lastName, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
		FirstName:         firstName,
		LastName:          lastName,
		AccountNumber:     int64(rand.Intn(1_000_000)),
		EncryptedPassword: string(encpw),
		CreatedAt:         time.Now().UTC(),
	}, nil
}
