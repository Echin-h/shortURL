package model

import (
	"gorm.io/gorm"
	"time"
)

// gorm.Model 的定义
type Base struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 时间要表示成时间戳
func (m Base) CreateTime() int64 {
	return m.CreatedAt.UnixMilli()
}

func (m Base) UpdateTime() int64 {
	return m.UpdatedAt.UnixMilli()
}
