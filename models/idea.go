package models


import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"time"
)




type Idea struct {
	ID	uint64		`gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time  	`gorm:"not null"`
	Content  string		`gorm:"size:400"`
}
