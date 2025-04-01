package models

import "time"

type CurrentPosition struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	AccountKey   string    `gorm:"column:account_key" json:"account_key"`
	Symbol       string    `gorm:"column:symbol" json:"symbol"`
	Ticket       int       `gorm:"column:ticket" json:"ticket"`
	Time         time.Time `gorm:"column:time" json:"time"`
	Type         string    `gorm:"column:type" json:"type"`
	Volume       float64   `gorm:"column:volume" json:"volume"`
	Price        float64   `gorm:"column:price" json:"price"`
	SL           float64   `gorm:"column:sl" json:"sl"`
	TP           float64   `gorm:"column:tp" json:"tp"`
	Value        float64   `gorm:"column:value" json:"value"`
	Swap         float64   `gorm:"column:swap" json:"swap"`
	Profit       float64   `gorm:"column:profit" json:"profit"`
	PriceChange  float64   `gorm:"column:price_change" json:"price_change"`
	MagicNumber  int       `gorm:"column:magic_number" json:"magic_number"`
	Comment      string    `gorm:"column:comment" json:"comment"`
}
