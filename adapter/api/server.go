package api

import (
	"FullCycle/payments/entity"
	"FullCycle/payments/usecase/process_transaction"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WebServer struct {
	Repository entity.TransactionRepository
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.POST("/transaction", w.Process)
	e.Logger.Fatal(e.Start(":8586"))
}

func (w WebServer) Process(c echo.Context) error {
	transactionDto := process_transaction.TransactionDtoInput{}
	//Qundo dou um Bind é para dizer que estamos recebendo a mesma extrutura, estou passando o conteudo de c para transactionDto
	c.Bind(transactionDto)
	usecase := process_transaction.NewProcessTransaction(w.Repository)
	output, _ := usecase.Execute(transactionDto)
	return c.JSON(http.StatusOK, output)

}
