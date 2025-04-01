package models

import "time"

type HistoryPosition struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	AccountKey  string    `gorm:"column:account_key" json:"account_key"`
	TimeEntry   time.Time `gorm:"column:time_entry" json:"time_entry"`
	Symbol      string    `gorm:"column:symbol" json:"symbol"`
	Ticket      int       `gorm:"column:ticket" json:"ticket"`
	Type        string    `gorm:"column:type" json:"type"`
	Volume      float64   `gorm:"column:volume" json:"volume"`
	PriceEntry  float64   `gorm:"column:price_entry" json:"price_entry"`
	SL          float64   `gorm:"column:sl" json:"sl"`
	TP          float64   `gorm:"column:tp" json:"tp"`
	TimeExit    time.Time `gorm:"column:time_exit" json:"time_exit"`
	PriceExit   float64   `gorm:"column:price_exit" json:"price_exit"`
	Value       float64   `gorm:"column:value" json:"value"`
	Commission  float64   `gorm:"column:commission" json:"commission"`
	Fee         float64   `gorm:"column:fee" json:"fee"`
	Swap        float64   `gorm:"column:swap" json:"swap"`
	Profit      float64   `gorm:"column:profit" json:"profit"`
	PriceChange float64   `gorm:"column:price_change" json:"price_change"`
	MagicNumber int       `gorm:"column:magic_number" json:"magic_number"`
	Comment     string    `gorm:"column:comment" json:"comment"`
}
