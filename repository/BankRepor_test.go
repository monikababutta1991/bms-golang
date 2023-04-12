package repository

import (
	"bank/domain"
	"testing"
)

const DSN string = "root:monika@tcp(127.0.0.1:3307)/bms?charset=utf8mb4&parseTime=True&loc=Local"

// test case 1: valid data
func Test_CreateBankInDB(t *testing.T) {
	BankObj, err := domain.NewBank("HDFC Bank")
	if err != nil {
		t.Fail()
	}

	BankRepoObj, errBnk := NewBankRepo(DSN)
	if errBnk != nil {
		t.Fail()
	}

	errCreateBnk := BankRepoObj.Save(BankObj)
	if errCreateBnk != nil {
		t.Fail()
	}
}

// test case 2: invalid / empty data
func Test_CreateBankInDbWithInvalidData(t *testing.T) {
	BankObj, err := domain.NewBank("")
	if err == nil {
		t.Fail()
	}

	BankRepoObj, errBnk := NewBankRepo(DSN)
	if errBnk != nil {
		t.Fail()
	}

	errCreateBnk := BankRepoObj.Save(BankObj)
	if errCreateBnk == nil {
		t.Fail()
	}
}
