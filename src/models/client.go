package model

type Client struct {
	Id          int
	Name        string
	CreditLimit int `json: "limite"`
	CreditUsed  int `json: "saldo"`
}
