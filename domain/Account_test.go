package domain

import (
	"testing"
	"time"
)

func Test_NewAccount(t *testing.T) {

	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	CustomerObj, errC := NewCustomer("John Doe", "johndoe@xyz.com", date1)
	if errC != nil {
		t.Fail()
	}

	BankObj, errB := NewBank("test bank name")
	if errB != nil {
		t.Fail()
	}

	_, err := NewAccount(CustomerObj.ID, BankObj.ID, 100000)
	if err != nil {
		t.Fail()
	}
}

func Test_NewAccountIn(t *testing.T) {

	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	CustomerObj, errC := NewCustomer("John Doe", "johndoe@xyz.com", date1)
	if errC != nil {
		t.Fail()
	}

	BankObj, errB := NewBank("test bank name")
	if errB != nil {
		t.Fail()
	}

	_, err := NewAccount(CustomerObj.ID, BankObj.ID, 9999)
	if err == nil {
		t.Fail()
	}
}

func Test_DepositMoneyInAccount(t *testing.T) {

	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	CustomerObj, errC := NewCustomer("John Doe", "johndoe@xyz.com", date1)
	if errC != nil {
		t.Fail()
	}

	BankObj, errB := NewBank("test bank name")
	if errB != nil {
		t.Fail()
	}

	errAcc := BankObj.OpenAccount(&CustomerObj, 100000)
	if errAcc != nil {
		t.Fail()
	}

	errDeposit := DepositMoneyInAccount(CustomerObj.Account, 200)
	if errDeposit != nil {
		t.Fail()
	}

}

// func Test_DepositMoneyInAccountIN() {

// }

func Test_WithdarwMoneyFromAccount(t *testing.T) {
	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	CustomerObj, errC := NewCustomer("John Doe", "johndoe@xyz.com", date1)
	if errC != nil {
		t.Fail()
	}

	BankObj, errB := NewBank("test bank name")
	if errB != nil {
		t.Fail()
	}

	errAcc := BankObj.OpenAccount(&CustomerObj, 100000)
	if errAcc != nil {
		t.Fail()
	}

	errDeposit := WithdarwMoneyFromAccount(CustomerObj.Account, 200)
	if errDeposit != nil {
		t.Fail()
	}
}

// func Test_WithdarwMoneyFromAccountIn() {

// }

func Test_getAccountBalance(t *testing.T) {
	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	CustomerObj, errC := NewCustomer("John Doe", "johndoe@xyz.com", date1)
	if errC != nil {
		t.Fail()
	}

	BankObj, errB := NewBank("test bank name")
	if errB != nil {
		t.Fail()
	}

	errAcc := BankObj.OpenAccount(&CustomerObj, 100000)
	if errAcc != nil {
		t.Fail()
	}

	_, errAccountB := getAccountBalance(CustomerObj.Account)
	if errAccountB != nil {
		t.Fail()
	}
}

func Test_TranserFundsSourceDestination(t *testing.T) {

	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)

	//create bank
	BankObj, errB := NewBank("test bank name")
	if errB != nil {
		t.Fail()
	}

	// create customers
	CustomerSObj, errC1 := NewCustomer("John Doe", "johndoe@xyz.com", date1)
	if errC1 != nil {
		t.Fail()
	}

	CustomerDObj, errC2 := NewCustomer("Lorem Ipsum", "loremipsum@xyz.com", date1)
	if errC2 != nil {
		t.Fail()
	}

	// open accounts
	erra1 := BankObj.OpenAccount(&CustomerSObj, 100000)
	if erra1 != nil {
		t.Fail()
	}

	erra2 := BankObj.OpenAccount(&CustomerDObj, 100000)
	if erra2 != nil {
		t.Fail()
	}

	errTransfer := TranserFundsSourceDestination(CustomerSObj.Account, CustomerDObj.Account, 200)
	if errTransfer != nil {
		t.Fail()
	}
}

// func Test_TranserFundsSourceDestinationIn() {

// }
