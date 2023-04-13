package domain

import (
	"testing"
	"time"
)

func Test_NewCustomer(t *testing.T) {

	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	_, err := NewCustomer("John Doe", "johndoe@xyz.com", date1)
	if err != nil {
		t.Fail()
	}
}

func Test_NewCustomerIn(t *testing.T) {

	date1 := time.Date(2006, time.January, 1, 12, 10, 0, 0, time.UTC)
	_, err := NewCustomer("John Doe", "johndoe@xyz.com", date1)
	if err == nil {
		t.Fail()
	}
}

func Test_AssignAccount(t *testing.T) {
	//create bank
	BankObj, errB := NewBank("test bank name")
	if errB != nil {
		t.Fail()
	}

	// create customer
	date1 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	CustomerSObj, errC := NewCustomer("John Doe", "johndoe@xyz.com", date1)
	if errC != nil {
		t.Fail()
	}

	// open account
	erra1 := BankObj.OpenAccount(&CustomerSObj, 100000)
	if erra1 != nil {
		t.Fail()
	}

	errAssignAcc := CustomerSObj.AssignAccount(CustomerSObj.Account)
	if errAssignAcc != nil {
		t.Fail()
	}
}
