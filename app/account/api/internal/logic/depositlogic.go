// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"github.com/wwater/my-cex/app/account/api/internal/model"
	"github.com/wwater/my-cex/app/wallet/rpc/walletclient"
	"github.com/wwater/my-cex/common/websocket"
	"strconv"
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

	amount, _ := strconv.ParseFloat(req.Amount, 64)
	logEntry := model.SystemLog{
		Uid:     req.UserId,
		OpType:  "DEPOSIT",
		Amount:  amount,
		Content: fmt.Sprintf("用户 %d 成功充值 %s ETH", req.UserId, req.Amount),
	}

	if err = l.svcCtx.DB.Create(&logEntry).Error; err != nil {
		l.Logger.Errorf("记录审计日志失败: %v", err)
	}

	websocket.GlobalHub.Broadcast(map[string]interface{}{
		"type":    "DEPOSIT",
		"userId":  req.UserId,
		"amount":  req.Amount,
		"time":    time.Now().Format("15:04:05"),
		"message": fmt.Sprintf("UID:%d 充值成功，金额: %s ETH", req.UserId, req.Amount),
	})

	return &types.DepositResp{Ok: true}, nil
}
