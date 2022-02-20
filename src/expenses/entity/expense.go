package entity

import "time"

type Expense struct {
	Uuid        string    `gorm:"primaryKey" json:"uuid"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Tags        string    `json:"tags"`
	Value       float32   `json:"value"`
	Description string    `json:"description"`
	User        string    `json:"user"`
	CreatedAt   time.Time `json:"created_at"`
}
