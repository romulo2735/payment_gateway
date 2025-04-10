package domain

type AccountRepository interface {
	Create(account *Account)
	FindByAPIKey(apiKey string) (*Account, error)
	FindById(id string) (*Account, error)
	Update(account *Account) error
}

type InvoiceRepository interface {
	Create(invoice *Invoice)
	FindById(id string) (*Invoice, error)
	FindByAccountID(accountID string) ([]*Invoice, error)
	UpdateStatus(invoice *Invoice) error
}
