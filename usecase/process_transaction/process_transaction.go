package process_transaction

import "FullCycle/payments/entity"

type ProcessTransaction struct {
	Repository entity.TransactionRepository
}

func NewProcessTransaction(repository entity.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Ammount = input.Ammount
	invalidTransaction := transaction.IsValid()
	if invalidTransaction != nil {
		err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Ammount, "rejected", invalidTransaction.Error())
		if err != nil {
			return TransactionDtoOutput{}, err
		}
		output := TransactionDtoOutput{
			ID:           transaction.ID,
			Status:       "rejected",
			ErrorMessage: invalidTransaction.Error(),
		}
		return output, nil
	}
	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Ammount, "approved", "")
	if err != nil {
		return TransactionDtoOutput{}, err
	}
	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       "approved",
		ErrorMessage: "",
	}
	return output, nil
}
