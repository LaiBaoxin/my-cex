// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/wwater/my-cex/app/account/api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/wwater/my-cex/app/wallet/rpc/walletclient"
)

type ServiceContext struct {
	Config    config.Config
	WalletRpc walletclient.Wallet
	DB        *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	return &ServiceContext{
		Config:    c,
		WalletRpc: walletclient.NewWallet(zrpc.MustNewClient(c.WalletRpc)),
		DB:        db,
	}
}
