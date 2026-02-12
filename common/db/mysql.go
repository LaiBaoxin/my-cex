package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitMysql 初始化 MySQL 连接
func InitMysql(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("[InitMySql] 连接数据库失败: %v", err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取数据库实例失败: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)  // 空闲连接数
	sqlDB.SetMaxOpenConns(100) // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("MySQL连接成功！")
	return db
}
