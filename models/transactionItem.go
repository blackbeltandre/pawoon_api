package models

import (
	"fmt"
	"gorm.io/gorm"
	u "pawoon/utils"
)

type TransactionItem struct {
	gorm.Model
	UserID   uint `json:"user_id"`
	UUID   string `json:"uuid"`
	TransactionID   string `json:"transaction_id"`
	Title   string `json:"title"`
	Qty   string `json:"qty"`
	Price  string `json:"price"`
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (transactionItem *TransactionItem) Validate() (map[string]interface{}, bool) {
	if transactionItem.Qty  == "" {
		return u.Message(false, "Qty should be on the payload"), false
	}
	if len(transactionItem.Price) < 1 {
		return u.Message(false, "Price is required"), false
	}
	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (transactionItem *TransactionItem) Create() (map[string]interface{}) {

	if resp, ok := transactionItem.Validate(); !ok {
		return resp
	}

	GetDB().Create(transactionItem)

	resp := u.Message(true, "success")
	resp["transactionItem"] = transactionItem
	return resp
}

func GetTransactionItem(id uint) (*TransactionItem) {

	transactionItem := &TransactionItem{}
	err := GetDB().Table("transaction_items").Where("id = ?", id).First(transactionItem).Error
	if err != nil {
		return nil
	}
	return transactionItem
}

func GetTransactionsItem(UserID uint) ([]*TransactionItem) {

	transactionItem := make([]*TransactionItem, 0)
	err := GetDB().Table("transaction_items").Where("user_id = ?", UserID).Find(&transactionItem).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return transactionItem
}

func GetTransactionItemDetail(id int, UserID uint) ([]*TransactionItem) {
	transactionItem := make([]*TransactionItem, 0)
	err := GetDB().Where("id = ? and user_id = ?", id, UserID).Find(&transactionItem).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return transactionItem
}
func DeleteTransactionItem(id int, UserID uint) ([]*TransactionItem) {
	transactionItem := make([]*TransactionItem, 0)
	err := GetDB().Where("id = ? and user_id = ?", id, UserID).Delete(&transactionItem).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return transactionItem
}
func (transactionItem *TransactionItem) UpdateTransactionItem(id int, UserID uint) (map[string]interface{}) {

	if resp, ok := transactionItem.Validate(); !ok {
		return resp
	}
	err := GetDB().Where("id = ? and user_id = ?", id, UserID).Updates(TransactionItem{UserID: transactionItem.UserID,UUID: transactionItem.UUID,TransactionID: transactionItem.TransactionID,Title: transactionItem.Title, Qty: transactionItem.Qty, Price: transactionItem.Price}).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resp := u.Message(true, "Update Transaction Item success")
	resp["transactionItem"] = transactionItem
	return resp
}