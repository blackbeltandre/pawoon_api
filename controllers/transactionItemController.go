package controllers

import (
	"encoding/json"
	"pawoon/models"
	u "pawoon/utils"
	"net/http"
	"strconv"
	"fmt"
)

var CreateTransactionItem = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	transactionItem := &models.TransactionItem{}

	err := json.NewDecoder(r.Body).Decode(transactionItem)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	transactionItem.UserID = user
	resp := transactionItem.Create()
	u.Respond(w, resp)
}

var GetTransactionItem = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetTransactionsItem(id)
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetTransactionDetailItem = func(w http.ResponseWriter, r *http.Request) {
	idnya := r.URL.Query().Get("id")
	id,_ := strconv.Atoi(idnya)
	fmt.Println(id)
	UserID := r.Context().Value("user").(uint)
	data := models.GetTransactionItemDetail(id,UserID)
	resp := u.Message(true, "Display Detail Transaction Item Success")
	resp["data"] = data
	u.Respond(w, resp)
}


var DeleteTransactionItem = func(w http.ResponseWriter, r *http.Request) {
	idnya := r.URL.Query().Get("id")
	id,_ := strconv.Atoi(idnya)
	fmt.Println(id)
	UserID := r.Context().Value("user").(uint)
	data := models.DeleteTransactionItem(id,UserID)
	resp := u.Message(true, "Delete Transaction Item Success")
	resp["data"] = data
	u.Respond(w, resp)
}


var UpdateTransactionItem = func(w http.ResponseWriter, r *http.Request) {
	//body, err := ioutil.ReadAll(r.Body)
	transactionItem := &models.TransactionItem{}
	err := json.NewDecoder(r.Body).Decode(transactionItem)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request Transaction body"))
		return
	}
	idnya := r.URL.Query().Get("id")
	id,_ := strconv.Atoi(idnya)
	fmt.Println(id)
	UserID := r.Context().Value("user").(uint)
	data := transactionItem.UpdateTransactionItem(id,UserID)
	resp := u.Message(true, "Update Transaction Item Success")
	resp["data"] = data
	u.Respond(w, resp)
}


