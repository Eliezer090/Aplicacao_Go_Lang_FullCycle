package entity

import "errors"

type Transaction struct {
	ID           string
	AccountID    string
	Ammount      float64
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

//Método, Atacha uma func a uma struct e vira um método
func (t *Transaction) IsValid() error {
	if t.Ammount > 1000 {
		return errors.New("you dont have limit for execute this transaction")
	}
	if t.Ammount < 1 {
		return errors.New("the ammount must be greater than 1")
	}
	return nil
}
