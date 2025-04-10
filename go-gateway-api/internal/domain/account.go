package domain

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/google/uuid"
	"sync"
	"time"
)

type Account struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
	mu        sync.RWMutex // read-write mutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Generates a random API key using crypto/rand
func generateAPIKey() string {
	key := make([]byte, 16)
	rand.Read(key)
	return hex.EncodeToString(key)
}

// NewAccount creates a new account with unique ID, API key, and initial balance
func NewAccount(name, email string) *Account {
	account := &Account{
		ID:      uuid.New().String(),
		Name:    name,
		Email:   email,
		Balance: 0,
		APIKey:  generateAPIKey(),

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}

// Deposit adds the specified amount to the account's balance
func (a *Account) Deposit(amount float64) {
	a.mu.Lock()         // trava para a operação ser segura
	defer a.mu.Unlock() // libera o mutex ao final da operação
	a.Balance += amount
	a.UpdatedAt = time.Now()
}
