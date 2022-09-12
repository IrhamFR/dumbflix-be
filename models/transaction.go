package models

import "time"

type Transaction struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	StartDate time.Time `json:"-"`
	DueDate   time.Time `json:"-"`
	// UserID    User      `json:"user_id"`
	Attach int    `json:"attach" gorm:"type:varchar(25)"`
	Status string `json:"status"  gorm:"type:varchar(25)"`
}
