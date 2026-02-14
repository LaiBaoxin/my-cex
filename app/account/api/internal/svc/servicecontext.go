// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wwater/my-cex/app/account/api/internal/config"
	"github.com/wwater/my-cex/common/contracts/mytoken"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/wwater/my-cex/app/wallet/rpc/walletclient"
)

type ServiceContext struct {
	Config    config.Config
	WalletRpc walletclient.Wallet
	DB        *gorm.DB
	TokenSvc  *mytoken.MyToken
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}
	// ETH 连接
	client, _ := ethclient.Dial("http://127.0.0.1:8545")
	// 部署得到的合约地址
	contractAddr := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	instance, _ := mytoken.NewMyToken(contractAddr, client)

	return &ServiceContext{
		Config:    c,
		WalletRpc: walletclient.NewWallet(zrpc.MustNewClient(c.WalletRpc)),
		DB:        db,
		TokenSvc:  instance,
	}
}
