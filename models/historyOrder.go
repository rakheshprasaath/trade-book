package models

import "time"

type HistoryOrder struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	AccountKey string    `gorm:"column:account_key" json:"account_key"`
	TimeEntry  time.Time `gorm:"column:time_entry" json:"time_entry"`
	Symbol     string    `gorm:"column:symbol" json:"symbol"`
	Ticket     int       `gorm:"column:ticket" json:"ticket"`
	Type       string    `gorm:"column:type" json:"type"`
	Volume     float64   `gorm:"column:volume" json:"volume"`
	Price      float64   `gorm:"column:price" json:"price"`
	Value      float64   `gorm:"column:value" json:"value"`
	SL         float64   `gorm:"column:sl" json:"sl"`
	TP         float64   `gorm:"column:tp" json:"tp"`
	Time       time.Time `gorm:"column:time" json:"time"`
	State      string    `gorm:"column:state" json:"state"`
	MagicNumber int      `gorm:"column:magic_number" json:"magic_number"`
	Comment    string    `gorm:"column:comment" json:"comment"`
}
