package controllers

import (
	"fmt"
	model "golang-extract-api/src/models"
	"golang-extract-api/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClientController struct {
	clientService  *services.ClientService
	extractService *services.ExtractService
}

func NewClientController(clientService *services.ClientService, extractService *services.ExtractService) *ClientController {
	return &ClientController{
		clientService,
		extractService,
	}
}

func (cc *ClientController) GetExtract(ctx *gin.Context) {
	clientIdString := ctx.Param("id")
	clientId, err := strconv.Atoi(clientIdString)

	if err != nil {
		ctx.String(http.StatusBadRequest, "Falha ao buscar o id do usuário")
		ctx.Abort()
		return
	}

	client, err := cc.clientService.FindById(clientId, ctx)

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}

	clientLastTransactions, err := cc.extractService.GetLastTransactionsBy(clientId, ctx)

	if err != nil {
		ctx.String(http.StatusBadRequest, "Falha ao procurar as transacoes do usuário")
		ctx.Abort()
		return
	}

	lastItemInsertedExtract := clientLastTransactions[0]

	ctx.JSON(http.StatusOK, gin.H{
		"saldo": gin.H{
			"total":        client.CreditUsed,
			"data_extrato": lastItemInsertedExtract.CreatedAt,
			"limite":       client.CreditLimit,
		},
		"ultimas_transacoes": clientLastTransactions,
	})
}

func (cc *ClientController) UpdateCreditUsed(ctx *gin.Context) {
	clientIdString := ctx.Param("id")
	clientId, err := strconv.Atoi(clientIdString)

	if err != nil {
		ctx.String(http.StatusBadRequest, "Falha ao buscar o id do usuário")
		ctx.Abort()
		return
	}

	var createExtract model.CreateExtract

	if err := ctx.BindJSON(&createExtract); err != nil {
		ctx.String(http.StatusInternalServerError, "Falha ao fazer parse do request.body")
		ctx.Abort()
		return
	}

	client, err := cc.clientService.UpdateCreditUsed(clientId, createExtract)

	if err != nil {
		fmt.Printf("err: %v", err)
		ctx.String(http.StatusUnprocessableEntity, err.Error())
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, client)
}
