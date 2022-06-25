package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"uts.com/umd/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.Home).Methods("GET")
	router.HandleFunc("/search/{name}", handlers.Search).Methods("GET")

	fmt.Println("[UMD SERVER STARTED ON http://localhost:8080]")
	log.Fatal(http.ListenAndServe(":8080", router))
}
