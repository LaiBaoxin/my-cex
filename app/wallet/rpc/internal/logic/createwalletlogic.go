package logic

import (
	"context"
	"fmt"

	"github.com/wwater/my-cex/app/wallet/rpc/internal/svc"
	"github.com/wwater/my-cex/app/wallet/rpc/wallet"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWalletLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWalletLogic {
	return &CreateWalletLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateWalletLogic) CreateWallet(in *wallet.CreateWalletReq) (*wallet.CreateWalletResp, error) {
	// 设定账户密码
	password := fmt.Sprintf("wwater_%d", in.Uid)

	// Keystore 生成新账户
	account, err := l.svcCtx.KeyStore.NewAccount(password)
	if err != nil {
		l.Logger.Errorf("创建钱包失败: %v", err)
		return nil, err
	}

	// 获取生成的0x地址
	address := account.Address.Hex()
	l.Logger.Infof("[Wallet RPC] 用户ID: %d 生成地址: %s", in.Uid, address)

	return &wallet.CreateWalletResp{
		Address: address,
	}, nil
}
