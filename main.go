package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Account struct (Model)
type Account struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Balance  float32   `json:"balance"`
	Customer *Customer `json:"customer"`
}

//Customer the account owner
type Customer struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {
	//Initialize router
	router := mux.NewRouter()
	//Route Handlers /EndPoint
	router.HandleFunc("/api/accounts", getAccounts).Methods("GET")
	router.HandleFunc("/api/accounts/{id}", getAccount).Methods("GET")
	router.HandleFunc("/api/accounts", createAccount).Methods("POST")
	router.HandleFunc("/api/accounts/{id}", updateAccount).Methods("PUT")
	router.HandleFunc("/api/accounts/{id}", deleteAccount).Methods("DELETE")

	//start server
	log.Fatal(http.ListenAndServe(":8000", router))

}

func getAccounts(w http.ResponseWriter, r *http.Request) {

}
func getAccount(w http.ResponseWriter, r *http.Request) {

}
func createAccount(w http.ResponseWriter, r *http.Request) {

}
func updateAccount(w http.ResponseWriter, r *http.Request) {

}
func deleteAccount(w http.ResponseWriter, r *http.Request) {

}
