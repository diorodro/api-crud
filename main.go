package main

import (
	"fmt"
	"net/http"
)

func configurarRotas() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sejam todos Bem-vindos")
	})
}

func configurarServidor() {
	configurarRotas()

	fmt.Println("Servidor rodando na porta 1337")
	http.ListenAndServe(":1337", nil)
}

func main() {
	configurarServidor()
}
