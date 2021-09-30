package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Merchant struct {
	ID                 string  `json:"id"`
	MerchantName       string  `json:"merchantName"`
	MerchantFee        float64 `json:"merchantFee"`
	MerchantFeeType    string  `json:"merchantFeeType" validate:"required,eq=FIX|eq=PERCENTAGE"`
	OperationalStatus  string  `json:"operationalStatus" validate:"required,eq=OPEN|eq=CLOSE"`
	Email              string  `json:"email" validate:"required,email"`
	Phone              string  `json:"phone" validate:"min=10,numeric"`
	MerchantUniqueCode string  `json:"merchantUniqueCode" validate:"required,alphanum"`
	Description        string  `json:"description"`
	Address            string  `json:"address"`
	Lat                float64 `json:"lat" validate:"latitude"`
	Long               float64 `json:"long" validate:"longitude"`
	CoverageArea       float64 `json:"coverageArea"`
	AverageRate        float64 `json:"averageRate"`
	RateCount          float64 `json:"rateCount"`
	CreatedAt          string  `json:"createdAt"`
	// UpdatedAt          time.Time      `json:"updatedAt"`
	// DeletedAt          gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (e *Merchant) BeforeCreate(scope *gorm.DB) (err error) {
	t := time.Now() 
	unixId := fmt.Sprintf("%v", time.Now().UnixNano())
	e.ID = unixId[len(unixId)-14:]
	e.CreatedAt = t.Format("2006-01-02 15:04:05")
	return
}

func (e *Merchant) TableName() string {
	return "merchant"
}
