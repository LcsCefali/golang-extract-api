package repositories

import (
	"context"
	"database/sql"
	"fmt"
	model "golang-extract-api/src/models"
)

type ClientRepository struct {
	database *sql.DB
}

func NewClientRepository(database *sql.DB) *ClientRepository {
	return &ClientRepository{
		database: database,
	}
}

func (cr *ClientRepository) FindById(id int, ctx context.Context) (model.Client, error) {
	client := model.Client{}

	err := cr.database.QueryRowContext(
		ctx,
		"SELECT id, name, credit_limit, credit_used FROM clients WHERE id = $1",
		id,
	).Scan(&client.Id, &client.Name, &client.CreditLimit, &client.CreditUsed)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Client{}, fmt.Errorf("FindById %d: client not exists", id)
		}

		return model.Client{}, err
	}

	return client, nil
}

func (cr *ClientRepository) UpdateCreditUsed(clientId int, amountUsed int, ctx context.Context) error {
	_, err := cr.database.ExecContext(ctx, "UPDATE clients SET credit_used = credit_used + $1 WHERE id = $2", amountUsed, clientId)

	if err != nil {
		return fmt.Errorf("UpdateCreditUsed %v: ", err)
	}

	return nil
}
