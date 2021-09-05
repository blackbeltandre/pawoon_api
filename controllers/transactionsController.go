package controllers

import (
	"encoding/json"
	"pawoon/models"
	u "pawoon/utils"
	"net/http"
	"strconv"
	"fmt"
)

var CreateTransaction = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	transaction := &models.Transaction{}

	err := json.NewDecoder(r.Body).Decode(transaction)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	transaction.UserID = user
	resp := transaction.Create()
	u.Respond(w, resp)
}

var GetTransaction = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetTransactions(id)
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetTransactionDetail = func(w http.ResponseWriter, r *http.Request) {
	idnya := r.URL.Query().Get("id")
	id,_ := strconv.Atoi(idnya)
	fmt.Println(id)
	UserID := r.Context().Value("user").(uint)
	data := models.GetTransactionDetail(id,UserID)
	resp := u.Message(true, "Display Detail Transaction Success")
	resp["data"] = data
	u.Respond(w, resp)
}


var DeleteTransaction = func(w http.ResponseWriter, r *http.Request) {
	idnya := r.URL.Query().Get("id")
	id,_ := strconv.Atoi(idnya)
	fmt.Println(id)
	UserID := r.Context().Value("user").(uint)
	data := models.DeleteTransaction(id,UserID)
	resp := u.Message(true, "Delete Transaction Success")
	resp["data"] = data
	u.Respond(w, resp)
}


var UpdateTransaction = func(w http.ResponseWriter, r *http.Request) {
	//body, err := ioutil.ReadAll(r.Body)
	transaction := &models.Transaction{}
	err := json.NewDecoder(r.Body).Decode(transaction)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request Transaction body"))
		return
	}
	idnya := r.URL.Query().Get("id")
	id,_ := strconv.Atoi(idnya)
	fmt.Println(id)
	UserID := r.Context().Value("user").(uint)
	data := transaction.UpdateTransaction(id,UserID)
	resp := u.Message(true, "Update Success")
	resp["data"] = data
	u.Respond(w, resp)
}


