package repository

import (
	"bank/domain"
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IBankRepo interface {
	GetByID(int) (domain.Bank, error)

	GetAll() ([]domain.Bank, error)

	Save(domain.Bank) error

	DeleteByID(int) error
}

type BankRepo struct {
	db *gorm.DB
}

// BankDTO: DATA TRANSFER OBJECT
type BankDTO struct {
	gorm.Model
	ID   int
	Name string
}

// constructor
func NewBankRepo(DSN string) (BankRepo, error) {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		return BankRepo{}, err
	}

	db.AutoMigrate(&BankDTO{})

	return BankRepo{db: db}, nil
}

func (br *BankRepo) GetByID(bid int) (domain.Bank, error) {
	var singleBank domain.Bank

	err := br.db.Where("id=?", bid).First(&singleBank).Error
	if err != nil {
		return domain.Bank{}, err
	}

	return singleBank, nil
}

func (br *BankRepo) GetAll() ([]domain.Bank, error) {

	var BankList []domain.Bank

	err := br.db.Find(&BankList).Error
	if err != nil {
		return []domain.Bank{}, err
	}

	return BankList, nil
}

func (br *BankRepo) Save(b domain.Bank) error {

	// Insert in bank table

	if b.Name == "" {
		return errors.New("bank name can not be empty")
	}

	newBank := &BankDTO{
		Name: b.Name,
		ID:   b.ID, //upsert
	}

	err := br.db.Save(newBank).Error
	if err != nil {
		return err
	}

	return nil
}

func (br *BankRepo) DeleteByID(bid int) error {

	err := br.db.Where("id = ?", bid).Delete(&BankDTO{}).Error
	if err != nil {
		return err
	}
	return nil
}
