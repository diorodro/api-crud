package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

func configurarServidor() {
	configurarRotas()

	fmt.Println("Servidor rodando na porta 1337")
	log.Fatal(http.ListenAndServe(":1337", nil))
}

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sejam todos Bem-vindos")
}

func listarLivros(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Livros)
}

func cadastrarLivro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var livro Livro

	resp, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Print("error:", err)
	}

	err1 := json.Unmarshal(resp, &livro)
	livro.Id = len(Livros) + 1

	if err1 == nil {
		Livros = append(Livros, livro)
	} else {
		fmt.Println(err1.Error())
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(livro)
}

func rotearLivros(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		listarLivros(w, r)
	} else if r.Method == "POST" {
		cadastrarLivro(w, r)
	}
}

func buscarLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	caminho := r.URL.Path
	zp := regexp.MustCompile(` */ *`)
	reqNumId := zp.Split(caminho, -1)

	numId, err := strconv.Atoi(reqNumId[2])

	if err != nil {
		fmt.Println(err.Error())
	}

	if numId <= len(Livros) && numId > 0 {
		numId--
		w.WriteHeader(http.StatusFound)

		encoder := json.NewEncoder(w)
		encoder.Encode(Livros[numId])

		return

	} else {
		w.WriteHeader(http.StatusNotFound)

		encoder := json.NewEncoder(w)
		encoder.Encode("ERROR: id not found!")

		return
	}
}

func configurarRotas() {
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/livros", rotearLivros)

	http.HandleFunc("/livros/", buscarLivros)
}

func main() {
	configurarServidor()
}
