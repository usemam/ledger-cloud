package datacontract

const (
	// TransactionTypeNewAccount - 'New Account' transaction type
	TransactionTypeNewAccount = iota
	// TransactionTypeTransfer - 'Transfer' transaction type
	TransactionTypeTransfer = iota
	// TransactionTypeDebit - 'Debit' transaction type
	TransactionTypeDebit = iota
	// TransactionTypeCredit - 'Credit' transaction type
	TransactionTypeCredit = iota
)

// Transaction - represents individual transaction and its attributes
type Transaction struct {
	UserID      string
	Date        int64
	Type        int
	Amount      float32
	Origin      string
	Destination string
}

// Account - represents named account and its attribute
type Account struct {
	UserID      string
	Name        string
	Balance     float32
	MaxCredit   float32
	LastUpdated int64
}

// Snapshot - represents total balance for given category within given time interval
type Snapshot struct {
	UserID       string
	CategoryName string
	Balance      float32
	DateFrom     int64
	DateTo       int64
}
