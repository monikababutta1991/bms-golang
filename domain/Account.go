package domain

import (
	"errors"
	"time"
)

const defaultAccOpenStatus = "active"
const minAccountOpeningFundLimit = 10000.00

type Account struct {
	ID            int
	AccountNumber int
	CustomerId    int
	BankId        int
	OpeningDate   time.Time
	Status        string
	Balance       float64
}

var lastAccountID int

// domain constructor
func NewAccount(customerid int, bankid int, deposit float64) (Account, error) {

	// check if deposit amount is >= minAccountOpeningFundLimit
	if deposit < minAccountOpeningFundLimit {
		return Account{}, errors.New("minimum new account opening balance limit does not meet")
	}

	lastAccountID++

	newAccountNum := generateAndSetAccountNumber(lastAccountID, customerid, bankid)

	newAccount := Account{ID: lastAccountID, CustomerId: customerid, BankId: bankid, OpeningDate: time.Now(), Status: defaultAccOpenStatus, Balance: deposit, AccountNumber: newAccountNum}

	return newAccount, nil
}

func generateAndSetAccountNumber(id, custId, bankId int) int {
	return id + custId + bankId
}

// Deposit money in customer's account
func DepositMoneyInAccount(a *Account, depositAmouunt float64) error {
	a.Balance += depositAmouunt
	return nil
}

func WithdarwMoneyFromAccount(a *Account, withdrawAmount float64) error {

	// check if customer's account balance will be >=10k after withdrawing funds from account
	updatedBalance := a.Balance - withdrawAmount
	if updatedBalance < minAccountOpeningFundLimit {
		return errors.New("customer will not have sufficient balance in account after this transaction")
	}

	// withdraw fund from account
	a.Balance -= withdrawAmount
	return nil
}

func getAccountBalance(a *Account) float64 {
	return a.Balance
}

func TranserFundsSourceDestination(sa *Account, da *Account, amount float64) error {

	// check if source has sufficient balance to transfer
	if sa.Balance < amount {
		return errors.New("Customer does not have sufficient account balance to make this transfer")
	}

	if sa.Balance-amount < minAccountOpeningFundLimit {
		return errors.New("Customer will not have sufficient balance in account after this transer so can not proceed")
	}

	sa.Balance -= amount
	da.Balance += amount
	return nil
}
