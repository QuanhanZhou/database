package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Product struct {
	Id          int				`json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Price       decimal.Decimal `json:"price,omitempty"`
}

type ProductInput struct { //为了只接收id
	Id int `json:"id"`
}
