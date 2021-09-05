package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"pawoon/app"
	"pawoon/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	
	router.HandleFunc("/api/transaction/new", controllers.CreateTransaction).Methods("POST")
	router.HandleFunc("/api/transaction", controllers.GetTransaction).Methods("GET")
	router.HandleFunc("/api/transaction/detail", controllers.GetTransactionDetail).Methods("GET")
	router.HandleFunc("/api/transaction/delete", controllers.DeleteTransaction).Methods("DELETE")
	router.HandleFunc("/api/transaction/update", controllers.UpdateTransaction).Methods("PATCH")

	router.HandleFunc("/api/transaction_item/new", controllers.CreateTransactionItem).Methods("POST")
	router.HandleFunc("/api/transaction_item", controllers.GetTransactionItem).Methods("GET")
	router.HandleFunc("/api/transaction_item/detail", controllers.GetTransactionDetailItem).Methods("GET")
	router.HandleFunc("/api/transaction_item/delete", controllers.DeleteTransactionItem).Methods("DELETE")
	router.HandleFunc("/api/transaction_item/update", controllers.UpdateTransactionItem).Methods("PATCH")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("port")
	if port == "" {
		port = "8002" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8002/api
	if err != nil {
		fmt.Print(err)
	}
}
