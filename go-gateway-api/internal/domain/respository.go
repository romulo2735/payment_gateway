package domain

type AccountRepository interface {
	Create(account *Account)
	FindByAPIKey(apiKey string) (*Account, error)
	FindById(id string) (*Account, error)
	Update(account *Account) error
}
