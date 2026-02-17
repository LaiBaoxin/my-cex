package model

import "time"

type AccountSystemLog struct {
	Id        int64     `gorm:"column:id;primaryKey"`
	Uid       int64     `gorm:"column:uid"`
	OpType    string    `gorm:"column:op_type"`
	Amount    float64   `gorm:"column:amount"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (AccountSystemLog) TableName() string {
	return "system_logs"
}
