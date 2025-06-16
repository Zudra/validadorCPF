package main

import (
	"log"
	"net/http"

	"validador-cpf/api"
)

func main() {
	http.HandleFunc("/validarcpf", api.ValidadorCpfHandler)

	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
