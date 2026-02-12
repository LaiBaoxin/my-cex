package model

import "time"

// UserAsset 对应数据库的 user_assets 表
type UserAsset struct {
	Id        int64     `gorm:"primaryKey"`
	Uid       int64     `gorm:"uniqueIndex:idx_uid_currency"`
	Address   string    `gorm:"uniqueIndex:idx_address;type:varchar(64)"`
	Currency  string    `gorm:"uniqueIndex:idx_uid_currency;type:varchar(10);default:ETH"`
	Balance   float64   `gorm:"type:decimal(30,18);default:0"`
	Status    int       `gorm:"default:1"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// TableName 指定表名
func (UserAsset) TableName() string {
	return "user_assets"
}
