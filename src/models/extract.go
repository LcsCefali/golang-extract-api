package model

import (
	"time"
)

type Extract struct {
	Id          int       `json:"-"`
	ClientId    int       `json:"-"`
	Amount      int       `json:"valor"`
	Operation   string    `json:"tipo" `
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizada_em"`
	UpdatedAt   time.Time `json:"-"`
}
