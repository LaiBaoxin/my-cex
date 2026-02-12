// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/wwater/my-cex/app/account/api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/wwater/my-cex/app/wallet/rpc/walletclient"
)

type ServiceContext struct {
	Config    config.Config
	WalletRpc walletclient.Wallet
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		WalletRpc: walletclient.NewWallet(zrpc.MustNewClient(c.WalletRpc)),
	}
}
