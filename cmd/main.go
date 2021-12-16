package main

import (
	"FullCycle/payments/adapter/api"
	"FullCycle/payments/adapter/repository"
	"FullCycle/payments/usecase/process_transaction"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	ExecTestsFinish(db)

}

func ExecTestsFinish(db *sql.DB) {
	repo := repository.NewTransactionRepositoryDb(db)
	webserver := api.NewWebServer()
	webserver.Repository = repo

	usecase := process_transaction.NewProcessTransaction(repo)

	fmt.Println(Execute(PassOkTest(), usecase))
	fmt.Println(Execute(ErrorTestDontHaveLimit(), usecase))
	fmt.Println(Execute(ErrorTestMustbeGreaterThan1(), usecase))

	webserver.Serve()
}

func Execute(input process_transaction.TransactionDtoInput, usecase *process_transaction.ProcessTransaction) string {
	output, err := usecase.Execute(input)
	if err != nil {
		fmt.Print(err.Error())
	}
	outputJson, _ := json.Marshal(output)
	return string(outputJson)
}

func PassOkTest() process_transaction.TransactionDtoInput {
	return process_transaction.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Ammount:   200,
	}
}

func ErrorTestDontHaveLimit() process_transaction.TransactionDtoInput {
	return process_transaction.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Ammount:   5000,
	}
}

func ErrorTestMustbeGreaterThan1() process_transaction.TransactionDtoInput {
	return process_transaction.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Ammount:   0,
	}
}
