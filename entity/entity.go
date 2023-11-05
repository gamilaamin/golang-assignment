package entity

import (
	"time"

	"gorm.io/gorm"
)

type Entity interface {
	TableName() string
}

type GMODEL struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
