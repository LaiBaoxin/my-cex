package model

import "time"

type UserTransaction struct {
	Id        int64     `gorm:"primaryKey"`
	TxId      string    `gorm:"uniqueIndex:idx_tx_id;type:varchar(64)"`
	Uid       int64     `gorm:"index:idx_uid"`
	Amount    float64   `gorm:"type:decimal(30,18)"`
	Type      int       `gorm:"default:1"` // 1: Deposit
	Status    int       `gorm:"default:1"` // 1: Pending, 2: Success
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (UserTransaction) TableName() string {
	return "user_transactions"
}
