package process_transaction

import (
	mock_entity "FullCycle/payments/entity/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransactionVali(t *testing.T) {
	input := TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Ammount:   200,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       "approved",
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	//a palavra defer vai esperar acabar de rodar para dar o finish na conexao com o banco de dados
	defer ctrl.Finish()
	repositoryMock := mock_entity.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Ammount, "approved", "").Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
