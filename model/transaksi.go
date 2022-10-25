package model

import "time"

type Transaksi struct {
	Id       string    `json:"id" binding:"required"`
	Menu	 string    `json:"menu" binding:"required"`
	Quantity int       `json:"qunatity" binding:"required"`
	Date     time.Time `json:"date"`
}