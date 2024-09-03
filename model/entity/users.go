package entity

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        int    `gorm:"type:int,primary_key"`
	Username  string `gorm:"type:varchar(255),uniqueIndex,not null"`
	Email     string `gorm:"uniqueIndex,not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
