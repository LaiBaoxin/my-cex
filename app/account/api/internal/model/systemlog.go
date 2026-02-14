package model

import "time"

type SystemLog struct {
	Id        int64     `gorm:"primaryKey"`
	Uid       int64     `gorm:"index:idx_uid"`
	OpType    string    `gorm:"type:varchar(20)"` // DEPOSIT, WITHDRAW
	Content   string    `gorm:"type:text"`
	Amount    float64   `gorm:"type:decimal(30,18)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (SystemLog) TableName() string {
	return "system_logs"
}
