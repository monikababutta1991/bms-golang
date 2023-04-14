package domain

import (
	"errors"
	"time"
)

type Customer struct {
	ID      int
	Name    string
	Email   string
	dob     time.Time
	Account *Account
}

const CUSTOMER_MIN_AGE = 18

var lastCustomerID int

// domain constructor | builder function
func NewCustomer(namein string, emailin string, dobin time.Time) (Customer, error) {

	// check customer age
	checkAgeLimit := checkCustomerAge(dobin)
	if !checkAgeLimit {
		return Customer{}, errors.New("Customer is not allowed to operate as age is less than allowed limit")
	}

	lastCustomerID++
	NewCustomer := Customer{ID: lastCustomerID, Name: namein, Email: emailin, dob: dobin}
	return NewCustomer, nil
}

func (c *Customer) GetDOB() time.Time {
	return c.dob
}

func checkCustomerAge(dob time.Time) bool {
	now := time.Now()
	years := now.Year() - dob.Year()
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		years--
	}
	return years >= CUSTOMER_MIN_AGE
}

// Assign an account to a customer
func (c *Customer) AssignAccount(a *Account) error {
	c.Account = a
	return nil
}
