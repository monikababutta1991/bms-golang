package repository

import (
	"bank/domain"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID            int
	CustomerId    int
	BankId        int
	AccountNumber int
	OpeningDate   time.Time
	Status        string
	Balance       float64
}

type IManageAccount interface {
	OpenNewAccount(domain.Customer, domain.Account, domain.Bank) error
}

type ManageAccount struct {
	db *gorm.DB
}

func ManageNewAccount(dsn string) IManageAccount {
	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	dbConnection.AutoMigrate(&Account{})

	return &ManageAccount{db: dbConnection}
}

func (ma *ManageAccount) OpenNewAccount(c domain.Customer, a domain.Account, b domain.Bank) error {

	accountDetails := Account{
		ID:            a.ID,
		CustomerId:    c.ID,
		BankId:        b.ID,
		AccountNumber: a.AccountNumber,
		OpeningDate:   time.Now(),
		Status:        a.Status,
		Balance:       a.Balance,
	}

	err := ma.db.Where(Account{ID: accountDetails.ID}).Assign(accountDetails).FirstOrCreate(&accountDetails).Error

	if err != nil {
		return err
	}
	return nil
}
