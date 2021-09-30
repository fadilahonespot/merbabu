package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type OperationalHours struct {
	ID         string `json:"id"`
	MerchantID string `json:"merchantId" validate:"required"`
	Day        string `json:"day" validate:"required"`
	OpenHour   string `json:"openHour" validate:"required"`
	CloseHour  string `json:"closeHour" validate:"required"`
}

type OperationalHoursUpdate struct {
	Day       string `json:"day" validate:"required"`
	OpenHour  string `json:"openHour" validate:"required"`
	CloseHour string `json:"closeHour" validate:"required"`
}

func (e *OperationalHours) TableName() string {
	return "operational_hours"
}

func (e *OperationalHours) BeforeCreate(scope *gorm.DB) (err error) {
	unixId := fmt.Sprintf("%v", time.Now().UnixNano())
	e.ID = "oph-" + unixId[len(unixId)-10:]
	return
}
