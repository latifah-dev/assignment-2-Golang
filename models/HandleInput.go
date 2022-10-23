package models

import (
	"encoding/json"
	"time"
)

type ItemInput struct {
	// ID          uint        `json:"lineItemId"`
	Item_Code   json.Number `json:"itemCode" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Quantity    json.Number `json:"quantity" binding:"required"`
}

type Orderitems struct {
	Customer_Name string       `json:"customerName" binding:"required"`
	Ordered_at    time.Time    `json:"orderedAt"`
	Items         []*ItemInput `json:"items" binding:"required,dive"`
}
