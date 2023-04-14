package repository

import (
	"bank/domain"
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ICustomerAccountRepo interface {

	// Command
	Save(domain.Customer) error

	// Query
	GetCustomerByID(int) (domain.Customer, error)

	// Query
	GetAllCustomers() ([]domain.Customer, error)

	// Command
	DeleteByID(int) error

	// Query
	GetCustomerByAccountNumber(int) (domain.Customer, error)
}

type CustomerAccountRepo struct {
	db *gorm.DB
}

type CustomerDTO struct {
	gorm.Model
	ID    int
	Name  string
	Email string
	dob   time.Time
	// AccountID int
	Account AccountDTO `gorm:"constraint:OnDelete:CASCADE;foreignKey:CustomerID"` //OnUpdate:CASCADE,
}
type AccountDTO struct {
	gorm.Model
	ID            int
	AccountNumber int `gorm:"unique"`
	CustomerID    int
	BankId        int
	OpeningDate   time.Time
	Status        string
	Balance       float64
}

func (cdto *CustomerDTO) TableName() string {
	return "customers"
}
func (adto *AccountDTO) TableName() string {
	return "accounts"
}

func NewCustomerAccountRepo(DSN string) (CustomerAccountRepo, error) {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		return CustomerAccountRepo{}, err
	}

	db.AutoMigrate(&CustomerDTO{})
	db.AutoMigrate(&AccountDTO{})

	custAccRepo := CustomerAccountRepo{db: db}
	return custAccRepo, nil
}

func (car *CustomerAccountRepo) Save(cust domain.Customer) error {

	if cust.Name == "" {
		return errors.New("customer name can not be empty")
	}

	// ... checks

	newCustomer := CustomerDTO{
		ID:    cust.ID,
		Name:  cust.Name,
		Email: cust.Email,
		dob:   cust.GetDOB(),
		Account: AccountDTO{
			ID:            cust.Account.ID,
			AccountNumber: cust.Account.AccountNumber,
			BankId:        cust.Account.BankId,
			OpeningDate:   cust.Account.OpeningDate,
			Status:        cust.Account.Status,
			Balance:       cust.Account.Balance,
		},
	}

	err := car.db.Save([]CustomerDTO{newCustomer}).Error
	if err != nil {
		return err
	}
	return nil
}

func (car *CustomerAccountRepo) GetCustomerByID(cID int) (domain.Customer, error) {
	var singleCustomer domain.Customer

	errCustomer := car.db.Where("id=?", cID).First(&singleCustomer).Error
	if errCustomer != nil {
		return domain.Customer{}, errCustomer
	}

	return singleCustomer, nil
}

func (car *CustomerAccountRepo) GetAllCustomers() ([]domain.Customer, error) {
	var customers []domain.Customer

	errCustomers := car.db.Find(&customers).Error

	if errCustomers != nil {
		return []domain.Customer{}, errCustomers
	}
	return customers, nil
}

func (car *CustomerAccountRepo) DeleteByID(cID int) error {

	errDelete := car.db.Delete(&CustomerDTO{}, cID).Error
	if errDelete != nil {
		return errDelete
	}

	return nil
}

func (car *CustomerAccountRepo) GetCustomerByAccountNumber(accNum int) (domain.Customer, error) {
	var singleCustomer domain.Customer

	// errCustomer := car.db.Preload("Accounts").Where("account_number=?", accNum).First(&singleCustomer).Error

	// db.Preload("Orders").Find(&users)

	errCustomer := car.db.Joins("customers").Joins("Account").Find(&singleCustomer, "Account.account_number = ?", accNum).Error

	if errCustomer != nil {
		return domain.Customer{}, errCustomer
	}

	return singleCustomer, nil
}
