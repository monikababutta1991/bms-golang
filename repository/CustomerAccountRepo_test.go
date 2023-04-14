package repository

import (
	"bank/domain"
	"testing"
	"time"
)

// const DSNA string = "root:monika@tcp(127.0.0.1:3307)/bms?charset=utf8mb4&parseTime=True&loc=Local"

func Test_CreateCustomerInDB(t *testing.T) {

	custAccRepo, err := NewCustomerAccountRepo(DSN)

	if err != nil {
		t.Fail()
	}

	date1 := time.Date(2000, time.January, 1, 12, 10, 0, 0, time.UTC)
	custObj, custErr := domain.NewCustomer("Ajay", "ajay@abx.com", date1)
	if custErr != nil {
		t.Fail()
	}

	bankObj, errBank := domain.NewBank("ABC Bank")
	if errBank != nil {
		t.Fail()
	}

	errAccOpen := bankObj.OpenAccount(&custObj, 1000000)
	if errAccOpen != nil {
		t.Fail()
	}

	errCustRepo := custAccRepo.Save(custObj)
	if errCustRepo != nil {
		t.Fail()
	}
}

func Test_GetCustomerByID(t *testing.T) {
	CustAccRepo, err := NewCustomerAccountRepo(DSN)
	if err != nil {
		t.Fail()
	}

	_, errCust := CustAccRepo.GetCustomerByID(1)
	if errCust != nil {
		t.Fail()
	}
}

func Test_GetAllCustomers(t *testing.T) {
	custAccRepo, err := NewCustomerAccountRepo(DSN)
	if err != nil {
		t.Fail()
	}

	_, errCust := custAccRepo.GetAllCustomers()
	if errCust != nil {
		t.Fail()
	}
}

func Test_DeleteByID(t *testing.T) {
	CustAccRepo, err := NewCustomerAccountRepo(DSN)

	if err != nil {
		t.Fail()
	}

	errDelCust := CustAccRepo.DeleteByID(1)
	if errDelCust != nil {
		t.Fail()
	}
}

func Test_GetCustomerByAccountNumber(t *testing.T) {
	CustAccRepo, err := NewCustomerAccountRepo(DSN)
	if err != nil {
		t.Fail()
	}
	_, errCust := CustAccRepo.GetCustomerByAccountNumber(3)
	if errCust != nil {
		t.Fail()
	}
}
