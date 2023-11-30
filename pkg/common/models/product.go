package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model          // adds ID, created_at etc.
	Name         string `json:"name"`
	Description  string `json:"description"`
	Amount       int    `json:"amount"`
	CurrentPrice int    `json:"current_price"`
	CostPrice    int    `json:"cost_price"`
}
