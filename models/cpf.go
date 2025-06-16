package models

type CpfEntrada struct {
	Cpf string `json:"cpf"`
}

type CpfResultado struct {
	Validade string `json:"validade"`
}