package models

import "time"

type Image struct {
	ID        string `gorm:"primary_key"`
	Data      string
	CreatedAt time.Time
}
