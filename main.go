package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Livro struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

var Livros []Livro = []Livro{
	{
		Id:     1,
		Titulo: "Capitães de Areia",
		Autor:  "Jorge Amado",
	},
	{
		Id:     2,
		Titulo: "Dona Flor e seus dois Maridos",
		Autor:  "Jorge Amado",
	},
	{
		Id:     3,
		Titulo: "Porto dos Milagres",
		Autor:  "Jorge Amado",
	},
}

func configurarRotas() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sejam todos Bem-vindos")
	})
	http.HandleFunc("/livros", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(Livros)
	})
}

func configurarServidor() {
	configurarRotas()

	fmt.Println("Servidor rodando na porta 1337")
	log.Fatal(http.ListenAndServe(":1337", nil))
}

func main() {
	configurarServidor()
}
