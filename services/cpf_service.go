package services

import (
	"errors"
	"regexp"

	"validador-cpf/models"
)

func ValidadorCpf(e models.CpfEntrada) (models.CpfResultado, error) {
	err := verifyFormat(e.Cpf)
	if err != nil {
		return models.CpfResultado{}, err
	}

	cpf, err := removeCharacters(e.Cpf)
	if err != nil {
		return models.CpfResultado{}, err
	}

	PrimeiraOrdemValidade, err := calculateValidationDigit(cpf[0:9], cpf[9:], 1)
	if err != nil {
		return models.CpfResultado{}, err
	}
	SegundaOrdemValidade, err := calculateValidationDigit(cpf[0:10], cpf[10:], 2)
	if err != nil {
		return models.CpfResultado{}, err
	}

	if PrimeiraOrdemValidade && SegundaOrdemValidade {
		return models.CpfResultado{
			Validade: "Válido",
		}, nil
	}

	return models.CpfResultado{}, errors.New("cpf inválido")
}

func verifyFormat(cpf string) error {
	re := regexp.MustCompile(`([0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2})`)
	validade := re.MatchString(cpf)

	if validade {
		return nil
	}

	return errors.New("documento com formatação inválida")
}

func removeCharacters(cpf string) (string, error) {
	re := regexp.MustCompile("[0-9]+")
	digits := re.FindAllString(cpf, -1)
	cpfValido := ""
	
	for _, d := range digits {
		cpfValido += d
	}

	if len(cpfValido) != 11 {
		return "", errors.New("cpf com erro pós formatação")
	}

	return cpfValido, nil
}

func calculateValidationDigit(digitos string, validador string, ordem int) (bool, error) {
	var somaDigitos int
	var subOrdem int

	switch {
	case ordem == 1:
		subOrdem = 10
	case ordem == 2:
		subOrdem = 11
	}

	for index, value := range digitos {
		somaDigitos += (int(value-'0') * (subOrdem - index))
	}

	resto := somaDigitos % 11
	digitoValidadorCalculado := 11 - resto

	digitoValidador := int(validador[0] - '0')

	switch {
	case digitoValidadorCalculado >= 10:
		if digitoValidador == 0 {
			return true, nil
		}
		return false, errors.New("cpf inválido")
	default:
		if digitoValidadorCalculado == digitoValidador {
			return true, nil
		}
		return false, errors.New("cpf inválido")
	}
}
