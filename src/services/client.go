package services

import (
	"context"
	"database/sql"
	"golang-extract-api/src/configuration/database"
	model "golang-extract-api/src/models"
	"golang-extract-api/src/repositories"
)

type ClientService struct {
	repository     *repositories.ClientRepository
	extractService *ExtractService
}

func NewClientService(repository *repositories.ClientRepository, extractService *ExtractService) *ClientService {
	return &ClientService{
		repository,
		extractService,
	}
}

func (cs *ClientService) FindById(clientId int, ctx context.Context) (model.Client, error) {
	return cs.repository.FindById(clientId, ctx)
}

func (cs *ClientService) UpdateCreditUsed(clientId int, createExtract model.CreateExtract) (model.Client, error) {
	transactionCtx := context.Background()
	var client model.Client

	err := database.UseTransaction(transactionCtx, func(tx *sql.Tx) error {
		multiplier := 1
		if createExtract.Operation == "d" {
			multiplier = -1
		}

		amountUsed := createExtract.Amount * multiplier

		if err := cs.repository.UpdateCreditUsed(clientId, amountUsed, transactionCtx); err != nil {
			return err
		}

		createdItemId, err := cs.extractService.CreateItem(clientId, createExtract, transactionCtx)

		if err != nil || createdItemId <= 0 {
			return err
		}

		client, err = cs.repository.FindById(clientId, transactionCtx)

		if err != nil {
			return err
		}

		return nil
	})

	return client, err
}
