package api

import (
	"encoding/json"
	"net/http"

	"validador-cpf/models"
	"validador-cpf/services"
)

func ValidadorCpfHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var entrada models.CpfEntrada
	err := json.NewDecoder(r.Body).Decode(&entrada)
	if err != nil {
		http.Error(w, "Erro ao ler entrada: "+err.Error(), http.StatusBadRequest)
		return
	}

	resultado, err := services.ValidadorCpf(entrada)
	if err != nil {
		http.Error(w, "Erro na validação: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultado)
}