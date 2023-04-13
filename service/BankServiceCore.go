package service

//aka use case (work flow)
//integration testng for service not unit: advance

import (
	"bank/domain"
	"errors"
	"log"
	"time"
)

//CQRS => Command and Query Responsibility segregation

type IBankServiceCore interface {

	//Commands : write or update operations and return error
	OpenAnAccount(CustomerName string, CustomerEmail string, CustomerDOB time.Time, depositAmount float64) error

	DepositFundInAccount(AccNumber int, depositAmount float64) error

	WithdrawFundFromAccount(AccNumber int, withdrawAmount float64) error

	TransferFund(AccNumSource int, AccNumDest int, transferAmount float64) error

	//Query: read operation and return data
	CheckBankBalance(AccNumSource int) float64
}

type BankServiceCore struct {
	customers []*domain.Customer
	bank      *domain.Bank
}

// builder func
func NewBankServiceCore() (BankServiceCore, error) {

	bankObj, err := domain.NewBank("P&B Bank")
	if err != nil {
		return BankServiceCore{}, err
	}

	bsc := BankServiceCore{customers: make([]*domain.Customer, 0), bank: &bankObj}
	return bsc, nil
}

// Command
func (bsc *BankServiceCore) OpenAnAccount(CustomerName string, CustomerEmail string, CustomerDOB time.Time, depositAmount float64) error {

	customer, err := domain.NewCustomer(CustomerName, CustomerEmail, CustomerDOB)
	if err != nil {
		return err
	}

	bank, err1 := domain.NewBank("P&B Bank")
	if err1 != nil {
		return err1
	}

	// open account
	err2 := bank.OpenAccount(&customer, depositAmount)
	if err2 != nil {
		return err2
	}

	log.Println(customer.Account.AccountNumber)

	bsc.customers = append(bsc.customers, &customer)

	return nil

	// and many more services can be called here like sending an sms, email ...
}

func (bsc *BankServiceCore) DepositFundInAccount(AccNumber int, depositAmount float64) error {

	// get customer from acc no

	// log.Println(bsc.customers)

	for _, customer := range bsc.customers {
		if customer.Account.AccountNumber == AccNumber {
			// fmt.Println(customer.Account.AccountNumber)
			err := bsc.bank.DepositMoney(customer, depositAmount)
			if err != nil {
				return err
			}
			log.Println(customer.Account)
			return nil
		}
	}

	return errors.New("invalid account number")
}

func (bsc *BankServiceCore) WithdrawFundFromAccount(AccNumber int, WithdrawAmount float64) error {

	// fmt.Println(bsc.customers)

	for _, customer := range bsc.customers {
		if customer.Account.AccountNumber == AccNumber {
			err := bsc.bank.WithdrawMoney(customer, WithdrawAmount)
			if err != nil {
				return err
			}
			log.Println(customer.Account)
			return nil
		}
	}
	return errors.New("invalid account number")
}

func (bsc *BankServiceCore) TransferFund(AccNumSource int, AccNumDest int, transferAmount float64) error {

	var custSource *domain.Customer
	var custDest *domain.Customer

	for _, cust := range bsc.customers {
		if cust.Account.AccountNumber == AccNumSource {
			custSource = cust
		}
		if cust.Account.AccountNumber == AccNumDest {
			custDest = cust
		}
	}

	err := bsc.bank.TransferFund(custSource, custDest, transferAmount)
	if err != nil {
		return err
	}

	log.Println(custSource.Account)
	log.Println(custDest.Account)

	return nil
}

func (bsc *BankServiceCore) CheckBankBalance(AccNumber int) (float64, error) {

	var balance float64
	for _, cust := range bsc.customers {
		if cust.Account.AccountNumber == AccNumber {
			balance, err := bsc.bank.CheckCustomerBalance(cust)
			if err != nil {
				return balance, err
			}
		}
	}

	return balance, nil
}
