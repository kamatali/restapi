package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

//CustomResponse message on failure
type CustomResponse struct {
	Message string `json:"message"`
}

//Collections of accounts
var accounts []Account

func main() {
	//Some dummy data tart with
	accounts = append(accounts, Account{ID: "100001", Title: "John Doe account", Balance: 234.00,
		Customer: &Customer{ID: "cu-01", FirstName: "John", LastName: "Doe"}})
	accounts = append(accounts, Account{ID: "100002", Title: "Jane Smith - CA", Balance: 4054.87,
		Customer: &Customer{ID: "cu-02", FirstName: "Jane", LastName: "Smith"}})

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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //getting params
	for _, account := range accounts {
		if account.ID == params["id"] {
			json.NewEncoder(w).Encode(account)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&CustomResponse{Message: "Account with " + params["id"] + " not found"})

}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var account Account
	_ = json.NewDecoder(r.Body).Decode(&account)
	account.ID = strconv.Itoa(rand.Intn(1000000)) //generate unsafe interger and cast it to str
	accounts = append(accounts, account)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, account := range accounts {
		if account.ID == params["id"] {
			accounts = append(accounts[:index], accounts[index+1:]...)
			var account Account
			_ = json.NewDecoder(r.Body).Decode(&account)
			account.ID = params["id"]
			accounts = append(accounts, account)
			json.NewEncoder(w).Encode(account)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&CustomResponse{Message: "Account with " + params["id"] + " not found"})
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, account := range accounts {
		if account.ID == params["id"] {
			accounts = append(accounts[:index], accounts[index+1:]...)
			json.NewEncoder(w).Encode(&CustomResponse{Message: "Account with " + params["id"] + " has been deleted"})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&CustomResponse{Message: "Account with " + params["id"] + " not found"})
}
