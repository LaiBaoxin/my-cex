package svc

import (
	"github.com/wwater/my-cex/common/db"
	"gorm.io/gorm"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/wwater/my-cex/app/wallet/rpc/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	KeyStore *keystore.KeyStore
	DB       *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 私钥存放目录
	ksDir := "./keystore_data"
	if _, err := os.Stat(ksDir); os.IsNotExist(err) {
		_ = os.Mkdir(ksDir, 0700)
	}

	// 初始化 Geth Keystore
	ks := keystore.NewKeyStore(ksDir, keystore.StandardScryptN, keystore.StandardScryptP)

	// 初始化 MySQL 连接
	dbConn := db.InitMysql(c.DataSource)

	return &ServiceContext{
		Config:   c,
		KeyStore: ks,
		DB:       dbConn,
	}
}
