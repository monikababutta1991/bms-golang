package main

import (
	"bank/service"
	"log"
	"time"
)

func main() {

	// fmt.Println("Jai Guru Dev")

	// create new customer
	date1 := time.Date(2000, time.January, 1, 12, 10, 0, 0, time.UTC)
	// monika, err1 := domain.NewCustomer("Monika", "monika@abc.abcd", date1)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	// date2 := time.Date(1991, time.January, 1, 12, 10, 0, 0, time.UTC)
	// rohit, errn := domain.NewCustomer("ROhit", "rohit@abc.abcd", date2)
	// if errn != nil {
	// 	fmt.Println(errn)
	// 	return
	// }

	// // create new bank
	// bankAllahabad, err := domain.NewBank("Allahabad Bank")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// request bank to open an account and pass customer's details along w/deposit amount
	// err2 := bankAllahabad.OpenAccount(&monika, 11000)
	// fmt.Println(err2, monika.Account)

	// errb := bankAllahabad.OpenAccount(&rohit, 11000)
	// fmt.Println(errb, monika.Account)

	// deposit 1000 in customer's account and check new balance
	// bankAllahabad.DepositMoney(monika.Account, 1000.00)
	// err3 := bankAllahabad.DepositMoney(&monika, 2000)
	// fmt.Println(err3)

	// err4 := bankAllahabad.WithdrawMoney(&monika, 50)
	// fmt.Println(err4)

	// fmt.Println(monika.Account.Balance)

	// erra := bankAllahabad.TransferFund(&monika, &rohit, 10)
	// fmt.Println(erra)

	// balance1 := bankAllahabad.CheckCustomerBalance(&monika)
	// fmt.Println(balance1)

	// balance2 := bankAllahabad.CheckCustomerBalance(&rohit)
	// fmt.Println(balance2)

	bsc, _ := service.NewBankServiceCore()

	err := bsc.OpenAnAccount("Deepak", "deepak@qualsights.com", date1, 20000)
	if err != nil {
		log.Println(err)
	}

	errA := bsc.OpenAnAccount("Monika", "monika@qualsights.com", date1, 20000)
	if errA != nil {
		log.Println(errA)
	}

	errnew := bsc.DepositFundInAccount(4, 500)
	if errnew != nil {
		log.Println(errnew)
	}

	//
	errWith := bsc.WithdrawFundFromAccount(4, 5)
	if errWith != nil {
		log.Println(errWith)
	}

	errTrans := bsc.TransferFund(4, 7, 2)
	if errTrans != nil {
		log.Println(errTrans)
	}

	balance := bsc.CheckBankBalance(4)
	log.Println(balance)
}
