package main

import "errors"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(customerPtr *customer, trans transaction) error {
	if customerPtr.balance < trans.amount {
		return errors.New("insufficient funds")
	}
	
	if trans.transactionType == transactionDeposit {
		customerPtr.balance += trans.amount
		return nil
	}
	if trans.transactionType == transactionWithdrawal {
		customerPtr.balance -= trans.amount
		return nil
	}
	return errors.New("unknown transaction type")
} 
