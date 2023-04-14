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

func Test_GetByID(t *testing.T) {
	BankRepo, err := NewBankRepo(DSN)
	if err != nil {
		t.Fail()
	}

	_, errBank := BankRepo.GetByID(1)
	if errBank != nil {
		t.Fail()
	}
}

func Test_GetAll(t *testing.T) {
	BankRepo, err := NewBankRepo(DSN)
	if err != nil {
		t.Fail()
	}

	_, errBank := BankRepo.GetAll()
	if errBank != nil {
		t.Fail()
	}
}

func Test_DeleteBankByID(t *testing.T) {
	BankRepo, err := NewBankRepo(DSN)
	if err != nil {
		t.Fail()
	}

	errDelete := BankRepo.DeleteByID(1)
	if errDelete != nil {
		t.Fail()
	}
}
