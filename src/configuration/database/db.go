package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Connection struct {
	DB *sql.DB
}

func NewConnection() *Connection {
	database, err := Open("host.docker.internal", "admin", "admin")

	if err != nil {
		panic(err)
	}

	return &Connection{
		DB: database,
	}
}

func Open(host string, user string, password string) (*sql.DB, error) {
	//host.docker.internal
	return sql.Open("postgres", fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=golang-extract sslmode=disable", host, user, password))
}

func UseTransaction(ctx context.Context, callback func(*sql.Tx) error) error {
	connection := NewConnection()

	transaction, err := connection.DB.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		return err
	}

	defer transaction.Rollback()

	err = callback(transaction)

	if err != nil {
		return fmt.Errorf("transaction error: %v", err)
	}

	return transaction.Commit()
}
