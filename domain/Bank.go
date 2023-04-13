package domain

import (
	"errors"
)

type Bank struct {
	ID   int
	Name string
}

var lastBankID int

// domain constructor | builder function
func NewBank(namein string) (Bank, error) {

	if namein == "" {
		return Bank{}, errors.New("bank name can not be empty")
	}

	lastBankID++
	newBank := Bank{ID: lastBankID, Name: namein}

	return newBank, nil
}

func (b *Bank) OpenAccount(c *Customer, depositAmount float64) error {

	acc, err := NewAccount(c.ID, b.ID, depositAmount)
	if err != nil {
		return err
	}

	c.AssignAccount(&acc)
	return nil
}

// Deposit money
func (b *Bank) DepositMoney(c *Customer, depositAmount float64) error {
	// a.Balance += depositAmount

	if depositAmount <= 0 {
		return errors.New("deposit amount can not be 0")
	}

	err := DepositMoneyInAccount(c.Account, depositAmount)
	if err != nil {
		return err
	}
	// log.Println(c.Account)
	return nil
}

// Withdraw money from customer's account
func (b *Bank) WithdrawMoney(c *Customer, withdrawAmount float64) error {
	// a.Balance += depositAmount

	if withdrawAmount <= 0 {
		return errors.New("withdraw amount can not be 0")
	}

	err := WithdarwMoneyFromAccount(c.Account, withdrawAmount)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bank) CheckCustomerBalance(c *Customer) (float64, error) {
	return getAccountBalance(c.Account)
}

// transfer funds from one account to another
func (b *Bank) TransferFund(s *Customer, d *Customer, amount float64) error {

	if amount <= 0 {
		return errors.New("transfer amount can not be 0")
	}

	return TranserFundsSourceDestination(s.Account, d.Account, amount)
}
