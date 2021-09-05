package models

import (
	"fmt"
	"gorm.io/gorm"
	u "pawoon/utils"
)

type Transaction struct {
	gorm.Model
	UUID   string `json:"uuid"`
	UserID   uint `json:"user_id"`
	DeviceTimestamp   string `json:"device_timestamp"`
	TotalAmount   string `json:"total_amount"`
	PaidAmount  string `json:"paid_amount"`
	ChangeAmount  string `json:"change_amount"`
	PaymentMethod  string `json:"payment_method"`
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (transaction *Transaction) Validate() (map[string]interface{}, bool) {
	if transaction.TotalAmount  == "" {
		return u.Message(false, "TotalAmount should be on the payload"), false
	}
	if len(transaction.PaidAmount) < 1 {
		return u.Message(false, "PaidAmount is required"), false
	}
	if len(transaction.ChangeAmount) < 1 {
		return u.Message(false, "ChangeAmount is required"), false
	}


	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (transaction *Transaction) Create() (map[string]interface{}) {

	if resp, ok := transaction.Validate(); !ok {
		return resp
	}

	GetDB().Create(transaction)

	resp := u.Message(true, "success")
	resp["transaction"] = transaction
	return resp
}

func GetTransaction(id uint) (*Transaction) {

	transaction := &Transaction{}
	err := GetDB().Table("transactions").Where("id = ?", id).First(transaction).Error
	if err != nil {
		return nil
	}
	return transaction
}

func GetTransactions(UserID uint) ([]*Transaction) {

	transaction := make([]*Transaction, 0)
	err := GetDB().Table("transactions").Where("user_id = ?", UserID).Find(&transaction).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return transaction
}

func GetTransactionDetail(id int, UserID uint) ([]*Transaction) {
	transaction := make([]*Transaction, 0)
	err := GetDB().Where("id = ? and user_id = ?", id, UserID).Find(&transaction).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return transaction
}
func DeleteTransaction(id int, UserID uint) ([]*Transaction) {
	transaction := make([]*Transaction, 0)
	err := GetDB().Where("id = ? and user_id = ?", id, UserID).Delete(&transaction).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return transaction
}
func (transaction *Transaction) UpdateTransaction(id int, UserID uint) (map[string]interface{}) {

	if resp, ok := transaction.Validate(); !ok {
		return resp
	}
	err := GetDB().Where("id = ? and user_id = ?", id, UserID).Updates(Transaction{UUID: transaction.UUID,UserID: transaction.UserID,DeviceTimestamp: transaction.DeviceTimestamp, TotalAmount: transaction.TotalAmount, PaidAmount: transaction.PaidAmount, ChangeAmount: transaction.ChangeAmount, PaymentMethod: transaction.PaymentMethod}).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resp := u.Message(true, "Update Transaction success")
	resp["transaction"] = transaction
	return resp
}