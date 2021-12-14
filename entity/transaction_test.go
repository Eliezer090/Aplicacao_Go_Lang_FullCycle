package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionWithAmmountGreaterThen100(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Ammount = 2000
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "you dont have limit for execute this transaction", err.Error())
}

func TestTransactionWithAmmountGreaterLassThen1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Ammount = 0
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "the ammount must be greater than 1", err.Error())
}
