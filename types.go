package main

import (
	"math/rand"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

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
}

type Account struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	AccountNumber int64     `json:"accountNumber"`
	Balance       int64     `json:"balance"`
	CreatedAt     time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		// ID:            rand.Intn(10_000),
		FirstName:     firstName,
		LastName:      lastName,
		AccountNumber: int64(rand.Intn(1_000_000)),
		CreatedAt:     time.Now().UTC(),
	}
}
