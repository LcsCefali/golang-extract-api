package repositories

import (
	"context"
	"database/sql"
	"fmt"
	model "golang-extract-api/src/models"
)

type ExtractRepository struct {
	database *sql.DB
}

func NewExtractRepository(database *sql.DB) *ExtractRepository {
	return &ExtractRepository{
		database,
	}
}

func (er *ExtractRepository) FindAllBy(clientId int, ctx context.Context) ([]model.Extract, error) {
	rows, err := er.database.QueryContext(
		ctx,
		`SELECT id, amount, operation, created_at, updated_at FROM extracts WHERE client_id = $1 ORDER BY created_at DESC LIMIT 10`,
		clientId,
	)

	if err != nil {
		return []model.Extract{}, fmt.Errorf("FindAllBy %d: Não foi possível buscar o extrato desse cliente", clientId)
	}

	extracts := []model.Extract{}

	for rows.Next() {
		extract := model.Extract{}
		rows.Scan(&extract.Id, &extract.Amount, &extract.Operation, &extract.CreatedAt, &extract.UpdatedAt)
		extracts = append(extracts, extract)
	}

	return extracts, nil
}

func (er *ExtractRepository) CreateItem(clientId, amount int, operation, description string, ctx context.Context) (int64, error) {
	fail := func(err error) (int64, error) {
		return 0, fmt.Errorf("CreateItem: %v", err)
	}

	var createdItemId int64

	err := er.database.QueryRow(`
		INSERT INTO extracts (client_id, amount, operation, description) VALUES ($1, $2, $3, $4) RETURNING id
		`,
		clientId,
		amount,
		operation,
		description,
	).Scan(&createdItemId)

	if err != nil {
		return fail(err)
	}

	return createdItemId, nil
}
