package entity

import (
	uid "github.com/google/uuid"
	"gorm.io/gorm"
)

type MerchantPaymentAccount struct {
	ID                  string `json:"id"`
	MerchantID          string `json:"merchantId"`
	MerchantAccountType string `json:"merchantAccountType"`
	AccountNumber       string `json:"accountNumber"`
}

type MerchantPaymentAccountListView struct {
	PaymentAccountID string `json:"paymentAccountId"`
	AccountNumber    string `json:"accountNumber"`
	ProviderName     string `json:"providerName"`
	Type             string `json:"type"`
}

func (e *MerchantPaymentAccount) TableName() string {
	return "merchant_payment_account"
}

func (e *MerchantPaymentAccount) BeforeCreate(scope *gorm.DB) (err error) {
	e.ID = uid.New().String()
	return
}
