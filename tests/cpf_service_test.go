package tests

import (
	"testing"

	"validador-cpf/models"
	"validador-cpf/services"
)

func TestValidadorCpf(t *testing.T) {
	casos := []struct {
		entrada          models.CpfEntrada
		validadeEsperada string
		esperaErro       bool
	}{
		{models.CpfEntrada{Cpf: "076.029.819-01"}, "Válido", false},
		{models.CpfEntrada{Cpf: "076.029-81901"}, "Inválido", true},
		{models.CpfEntrada{Cpf: "07602981901"}, "Válido", false},
		{models.CpfEntrada{Cpf: "076.029819-01"}, "Válido", false},
		{models.CpfEntrada{Cpf: "076.029819-02"}, "Inválido", true},
	}

	for _, c := range casos {
		resultado, err := services.ValidadorCpf(c.entrada)

		if c.esperaErro && err == nil {
			t.Errorf("esperava erro, mas não houve para entrada: %+v", c.entrada)
		}

		if !c.esperaErro && err != nil {
			t.Errorf("erro inesperado para entrada %+v: %v", c.entrada, err)
		}

		if !c.esperaErro {
			if resultado.Validade != c.validadeEsperada {
				t.Errorf("válidade esperada: %s, obtido: %s", c.validadeEsperada, resultado.Validade)
			}
		}
	}
}
