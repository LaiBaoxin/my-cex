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

type WithdrawLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWithdrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawLogic {
	return &WithdrawLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Withdraw 提现
func (l *WithdrawLogic) Withdraw(req *types.WithdrawReq) (resp *types.WithdrawResp, err error) {
	// 生成流水号
	txId := fmt.Sprintf("wd_%d_%d", req.UserId, time.Now().UnixNano())

	rpcResp, err := l.svcCtx.WalletRpc.Withdraw(l.ctx, &walletclient.WithdrawReq{
		Uid:    req.UserId,
		Amount: req.Amount,
		TxId:   txId,
	})

	if err != nil {
		return nil, err
	}

	return &types.WithdrawResp{
		Ok:     rpcResp.Success,
		TxHash: rpcResp.TxHash,
	}, nil
}
