// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"github.com/wwater/my-cex/app/wallet/rpc/walletclient"

	"github.com/wwater/my-cex/app/account/api/internal/svc"
	"github.com/wwater/my-cex/app/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBalanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBalanceLogic {
	return &GetBalanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBalanceLogic) GetBalance(req *types.GetBalanceReq) (resp *types.GetBalanceResp, err error) {
	rpcResp, err := l.svcCtx.WalletRpc.GetBalance(l.ctx, &walletclient.GetBalanceReq{
		Uid:      req.UserId,
		Currency: req.Currency,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetBalanceResp{
		Balance: rpcResp.Balance,
	}, nil
}
