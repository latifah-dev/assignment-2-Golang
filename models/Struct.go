package models

import (
	"time"
)

type Order struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	Customer_Name string    `gorm:"not null; type:varchar(50)" json:"customerName"`
	Ordered_at    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"orderedAt"`
	Items         []Items   `grom:"constraint:OnUpdate:CASCADE,OnDelete: SET NULL" json:"items"`
}

type Items struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"lineItemId"`
	Item_Code   string `gorm:"not null;type:varchar(50)" json:"itemCode"`
	Description string `grom:"not null; type:TEXT" json:"description"`
	Quantity    uint   `gorm:"not null" json:"quantity"`
	OrderID     uint
}
