package repository

import (
	"bank/domain"
	"log"
	"testing"
	"time"
)

func Test_OpenNewAccount(t *testing.T) {
	var DSN string = "root:monika@tcp(127.0.0.1:3307)/bank_customer?parseTime=True"
	accMgr := ManageNewAccount(DSN)

	date1 := time.Date(2000, time.January, 1, 12, 10, 0, 0, time.UTC)
	customerObj, _ := domain.NewCustomer(1, "Monika", "monika@abc.abcd", date1)

	customerAccObj := domain.Account{}
	bankObj := domain.Bank{}

	err := accMgr.OpenNewAccount(customerObj, customerAccObj, bankObj)
	if err != nil {
		t.Fail()
	}
	log.Println("Account Created")

}
