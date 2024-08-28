package services

import (
	"context"
	model "golang-extract-api/src/models"
	"golang-extract-api/src/repositories"
)

type ExtractService struct {
	repository *repositories.ExtractRepository
}

func NewExtractService(repository *repositories.ExtractRepository) *ExtractService {
	return &ExtractService{
		repository,
	}
}

func (es *ExtractService) GetLastTransactionsBy(clientId int, ctx context.Context) ([]model.Extract, error) {
	return es.repository.FindAllBy(clientId, ctx)
}

func (es *ExtractService) CreateItem(clientId int, createExtract model.CreateExtract, ctx context.Context) (int64, error) {
	createdItemId, err := es.repository.CreateItem(
		clientId,
		createExtract.Amount,
		createExtract.Operation,
		createExtract.Description,
		ctx,
	)

	if err != nil {
		return 0, err
	}

	return createdItemId, nil
}
