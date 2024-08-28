package model

type CreateExtract struct {
	Amount      int    `json:"valor"`
	Operation   string `json:"tipo" fieldcontains:"c"`
	Description string `json:"descricao" fieldcontains:"Bla"`
}
