// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"github.com/wwater/my-cex/app/wallet/rpc/walletclient"
	"time"

	"github.com/wwater/my-cex/app/account/api/internal/svc"
	"github.com/wwater/my-cex/app/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DepositLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDepositLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepositLogic {
	return &DepositLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DepositLogic) Deposit(req *types.DepositReq) (resp *types.DepositResp, err error) {
	// 生成流水号
	txId := fmt.Sprintf("wwater_%d_%d", req.UserId, time.Now().UnixNano())

	_, err = l.svcCtx.WalletRpc.Deposit(l.ctx, &walletclient.DepositReq{
		Uid:    req.UserId,
		Amount: req.Amount,
		TxId:   txId,
	})

	if err != nil {
		return nil, err
	}

	return &types.DepositResp{Ok: true}, nil
}
