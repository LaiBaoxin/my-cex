// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"github.com/wwater/my-cex/app/account/api/internal/svc"
	"github.com/wwater/my-cex/app/account/api/internal/types"
	"github.com/wwater/my-cex/app/wallet/rpc/walletclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAccountLogic {
	return &CreateAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAccountLogic) CreateAccount(req *types.CreateAccountReq) (resp *types.CreateAccountResp, err error) {
	rpcResp, err := l.svcCtx.WalletRpc.CreateWallet(l.ctx, &walletclient.CreateWalletReq{
		Uid: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateAccountResp{
		Address: rpcResp.Address,
	}, nil
}
