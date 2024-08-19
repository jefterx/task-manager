package main

import (
	"log"
	"net/http"
	"taskmanager/router"
)

func main() {
	r := router.Router()
	log.Println("Servidor iniciado na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
