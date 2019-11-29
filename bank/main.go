package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type customer struct {
	CustomerID  int    `form:"customerId" json:"customerId"`
	FirstName   string `form:"firstname" json:"firstname"`
	LastName    string `form:"lastname" json:"lastname"`
	PhoneNumber string `form:"phonenumber" json:"phonenumber"`
	Gender      string `form:"gender" json:"gender"`
	Salary      int64  `form:"salary" json:"salary"`
}

type customerAccount struct {
	AccountID int   `form:"accountId" json:"accountId"`
	Pin       int   `form:"pin" json:"pin"`
	Balance   int64 `form:"balance" json:"balance"`
	Customer  customer
}

type bulkPayment struct {
	Escrow   int64 `form:"balance" json:"balance"`
}

func (c *customer) changePhoneNumber(newPhoneNumber string) error {
	fmt.Println("changing phone number for")
	c.PhoneNumber = newPhoneNumber
	return nil
}

func (c *customerAccount) changePin(newPin int) error {
	fmt.Println("changing pin")
	c.Pin = newPin
	return nil
}

func (c *customerAccount) deposit(amount int64) string {
	c.Balance += amount
	return fmt.Sprintf("deposited %v to account %v", amount, c.AccountID)
}

func (c *customerAccount) withdraw(amount int64) error {
	if c.Balance < amount {
		return errors.New("insufficient funds")
	}
	c.Balance -= amount
	fmt.Sprintf("wiithdrew %v to account %v", amount, c.AccountID)
	return nil
}

func (c *customerAccount) checkBalance(accountNumber int) string {
	if accountNumber == c.AccountID {
		return fmt.Sprintf("balance for account %v is: %v", c.AccountID, c.Balance)
	}
	return "Transaction failure: not authorized"
}

// TransferFunds moves account balance from one account to another
func TransferFunds(from, to *customerAccount, amount int64) error {
	if err := from.withdraw(amount); err != nil {
		return errors.New("insufficient funds to complete transaction")
	}
	to.deposit(amount)
	return nil
}

// creditAccountByID moves funds to an account
func (c *customerAccount) creditAccountByID(accountNumber int, amount int64) (customerAccount, error) {
	if accountNumber == c.AccountID {
		c.Balance += amount
		return customerAccount{}, nil
	}
	return customerAccount{}, errors.New("Transaction failure: Account cannot be found")
}

func main() {
	cust1 := &customerAccount{
		AccountID: 1,
		Pin:       1234,
		Balance:   10000,
		Customer: customer{
			CustomerID:  1,
			FirstName:   "danny",
			LastName:    "lwe",
			PhoneNumber: "0772504991",
			Gender:      "MALE",
			Salary: 4000,
		},
	}

	cust2 := &customerAccount{
		AccountID: 2,
		Pin:       1234,
		Balance:   20000,
		Customer: customer{
			CustomerID:  2,
			FirstName:   "moses",
			LastName:    "kamira",
			PhoneNumber: "0772504441",
			Gender:      "MALE",
			Salary: 9000,
		},
	}

	escrow := &customerAccount{
		AccountID: 3,
		Pin:       1234,
		Balance:   60000,
		Customer: customer{
			CustomerID:  3,
			FirstName:   "escrow",
			LastName:    "escrow",
			PhoneNumber: "0772504441",
			Gender:      "MALE",
		},
	}

	cust1Json, _ := json.Marshal(cust1)
	fmt.Println(cust1, cust2)
	fmt.Println(string(cust1Json))
	cust1.changePin(4400)
	cust1JsonC, _ := json.Marshal(cust1)
	fmt.Println(string(cust1JsonC))

	cust1.deposit(5000)
	deposit1, _ := json.Marshal(cust1)
	fmt.Println(string(deposit1))

	checkbalance1 := cust1.checkBalance(1)
	fmt.Println(checkbalance1)

	if err := TransferFunds(cust1, cust2, 3000); err != nil {
		fmt.Println(err)
	} else {
		fmt.Sprintf("transfering funds from %v to %v", cust1.AccountID, cust2.AccountID)
		fmt.Println(cust1.checkBalance(1))
		fmt.Println(cust2.checkBalance(2))
	}

	
	b := [...]customerAccount{*cust1, *cust2}
	for _, account := range b {
		
		s := &account
		fmt.Println(s)

		escrow.withdraw(s.Customer.Salary)
		s.deposit(s.Customer.Salary)

		newer,_ := json.Marshal(s)
		account := newer
		fmt.Println(string(account))
	}
	
	fmt.Println(cust1.checkBalance(1))
	fmt.Println(escrow.checkBalance(3))
	//send email/s
}

