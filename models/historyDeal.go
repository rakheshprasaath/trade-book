package models

import "time"

type HistoryDeal struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	AccountKey  string    `gorm:"column:account_key" json:"account_key"`
	TimeEntry   time.Time `gorm:"column:time_entry" json:"time_entry"`
	Symbol      string    `gorm:"column:symbol" json:"symbol"`
	Deal        int       `gorm:"column:deal" json:"deal"`
	OrderID     int       `gorm:"column:order_id" json:"order_id"`
	Type        string    `gorm:"column:type" json:"type"`
	Direction   string    `gorm:"column:direction" json:"direction"`
	Volume      float64   `gorm:"column:volume" json:"volume"`
	Price       float64   `gorm:"column:price" json:"price"`
	Value       float64   `gorm:"column:value" json:"value"`
	SL          float64   `gorm:"column:sl" json:"sl"`
	TP          float64   `gorm:"column:tp" json:"tp"`
	Commission  float64   `gorm:"column:commission" json:"commission"`
	Fee         float64   `gorm:"column:fee" json:"fee"`
	Swap        float64   `gorm:"column:swap" json:"swap"`
	Profit      float64   `gorm:"column:profit" json:"profit"`
	PriceChange float64   `gorm:"column:price_change" json:"price_change"`
	MagicNumber int       `gorm:"column:magic_number" json:"magic_number"`
	Comment     string    `gorm:"column:comment" json:"comment"`
}
