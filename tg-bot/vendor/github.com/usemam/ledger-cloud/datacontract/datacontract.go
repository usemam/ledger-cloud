package datacontract

import (
	m "github.com/usemam/ledger-cloud/money"
)

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
	Amount      m.Money
	Origin      string
	Destination string
}

// Account - represents named account and its attribute
type Account struct {
	UserID      string
	Name        string
	Balance     m.Money
	MaxCredit   m.Money
	LastUpdated int64
}

// Snapshot - represents total balance for given category within given time interval
type Snapshot struct {
	UserID       string
	CategoryName string
	Balance      m.Money
	DateFrom     int64
	DateTo       int64
}
