package models

import (
	"time"

	_ "github.com/lib/pq"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int       `gorm:"type:int;primary_key;identity(1,1)"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `json:"Password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}