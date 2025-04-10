package dto

import (
	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
	"time"
)

// CreateAccountInput represents account data in API requests
type CreateAccountInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// AccountOutput represents account data in API responses
type AccountOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	APIKey    string    `json:"api_key,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToAccount converts CreateAccountInput to domain.Account
func ToAccount(input CreateAccountInput) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}

// FromAccount convert domain.Account to AccountOutput
func FromAccount(account *domain.Account) AccountOutput {
	return AccountOutput{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		Balance:   account.Balance,
		APIKey:    account.APIKey,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}
