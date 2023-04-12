package domain

import (
	"testing"
	"time"
)

func Test_NewBank(t *testing.T) {
	_, err := NewBank("test bank name")
	if err != nil {
		t.Fail()
	}
}

func Test_NewBankIn(t *testing.T) {
	_, err := NewBank("")
	if err == nil {
		t.Fail()
	}
}

func Test_OpenAccount(t *testing.T) {

	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	CustomerObj, _ := NewCustomer("John Doe", "johndoe@xyz.com", date1)
	// if errc != nil {
	// 	t.Fail()
	// }

	BankObj, _ := NewBank("test bank name")
	// if errb != nil {
	// 	t.Fail()
	// }

	errAcc := BankObj.OpenAccount(&CustomerObj, 50000)
	if errAcc != nil {
		t.Fail()
	}
}

func Test_OpenAccountIn(t *testing.T) {

	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	CustomerObj, _ := NewCustomer("John Doe", "johndoe@xyz.com", date1)
	// if errc != nil {
	// 	t.Fail()
	// }

	BankObj, _ := NewBank("test bank name")
	// if errb != nil {
	// 	t.Fail()
	// }

	errAcc := BankObj.OpenAccount(&CustomerObj, 999)
	if errAcc == nil {
		t.Fail()
	}
}

func Test_DepositMoney(t *testing.T) {

	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	CustomerObj, _ := NewCustomer("John Doe", "johndoe@xyz.com", date1)

	BankObj, _ := NewBank("test bank name")

	err := BankObj.OpenAccount(&CustomerObj, 100000)
	if err != nil {
		t.Fail()
	}

	errDeposit := BankObj.DepositMoney(&CustomerObj, 900)
	if errDeposit != nil {
		t.Fail()
	}
	// fmt.Println(CustomerObj.Account)
}

func Test_DepositMoneyIn(t *testing.T) {

	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	CustomerObj, _ := NewCustomer("John Doe", "johndoe@xyz.com", date1)

	BankObj, _ := NewBank("test bank name")

	err := BankObj.OpenAccount(&CustomerObj, 100000)
	if err != nil {
		t.Fail()
	}

	errDeposit := BankObj.DepositMoney(&CustomerObj, 0)
	if errDeposit == nil {
		t.Fail()
	}
	// fmt.Println(CustomerObj.Account)
}

func Test_WithdrawMoney(t *testing.T) {

	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	CustomerObj, _ := NewCustomer("John Doe", "johndoe@xyz.com", date1)

	BankObj, _ := NewBank("test bank name")

	err := BankObj.OpenAccount(&CustomerObj, 100000)
	if err != nil {
		t.Fail()
	}

	errWith := BankObj.WithdrawMoney(&CustomerObj, 5000)
	if errWith != nil {
		t.Fail()
	}
}

func Test_WithdrawMoneyIn(t *testing.T) {

	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	CustomerObj, _ := NewCustomer("John Doe", "johndoe@xyz.com", date1)

	BankObj, _ := NewBank("test bank name")

	err := BankObj.OpenAccount(&CustomerObj, 100000)
	if err != nil {
		t.Fail()
	}

	errWith := BankObj.WithdrawMoney(&CustomerObj, 0)
	if errWith == nil {
		t.Fail()
	}
}

// INCOMPLETE
func Test_CheckCustomerBalance(t *testing.T) {

	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	CustomerObj, _ := NewCustomer("John Doe", "johndoe@xyz.com", date1)

	BankObj, _ := NewBank("test bank name")

	err := BankObj.OpenAccount(&CustomerObj, 100000)
	if err != nil {
		t.Fail()
	}

	// CustomerBalance := BankObj.CheckCustomerBalance(&CustomerObj)
	// if CustomerBalance
}

func Test_TransferFund(t *testing.T) {

	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)

	//create bank
	BankObj, _ := NewBank("test bank name")

	// create customers
	CustomerSObj, _ := NewCustomer("John Doe", "johndoe@xyz.com", date1)
	CustomerDObj, _ := NewCustomer("Lorem Ipsum", "loremipsum@xyz.com", date1)

	// open accounts
	erra1 := BankObj.OpenAccount(&CustomerSObj, 100000)
	if erra1 != nil {
		t.Fail()
	}

	erra2 := BankObj.OpenAccount(&CustomerDObj, 100000)
	if erra2 != nil {
		t.Fail()
	}

	// transfer funds
	errTransfer := BankObj.TransferFund(&CustomerSObj, &CustomerDObj, 5000)
	if errTransfer != nil {
		t.Fail()
	}
}

func Test_TransferFundIn(t *testing.T) {

	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)

	//create bank
	BankObj, _ := NewBank("test bank name")

	// create customers
	CustomerSObj, _ := NewCustomer("John Doe", "johndoe@xyz.com", date1)
	CustomerDObj, _ := NewCustomer("Lorem Ipsum", "loremipsum@xyz.com", date1)

	// open accounts
	erra1 := BankObj.OpenAccount(&CustomerSObj, 100000)
	if erra1 != nil {
		t.Fail()
	}

	erra2 := BankObj.OpenAccount(&CustomerDObj, 100000)
	if erra2 != nil {
		t.Fail()
	}

	// transfer funds
	errTransfer := BankObj.TransferFund(&CustomerSObj, &CustomerDObj, 0)
	if errTransfer == nil {
		t.Fail()
	}
}
