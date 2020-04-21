package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olá. Bem vindo a minha página!")
	})

	http.ListenAndServe(":8080", nil)
}