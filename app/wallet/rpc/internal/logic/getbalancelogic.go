package logic

import (
	"context"
	"fmt"
	"github.com/wwater/my-cex/app/wallet/rpc/internal/model"

	"github.com/wwater/my-cex/app/wallet/rpc/internal/svc"
	"github.com/wwater/my-cex/app/wallet/rpc/wallet"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBalanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBalanceLogic {
	return &GetBalanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBalanceLogic) GetBalance(in *wallet.GetBalanceReq) (*wallet.GetBalanceResp, error) {
	var asset model.UserAsset

	// 查库
	err := l.svcCtx.DB.Where("uid = ? AND currency = ?", in.Uid, in.Currency).First(&asset).Error

	if err != nil {
		// 如果没找到，返回余额 0
		return &wallet.GetBalanceResp{
			Balance:  "0.0000",
			Currency: in.Currency,
		}, nil
	}

	// 返回结果
	return &wallet.GetBalanceResp{
		Balance:  fmt.Sprintf("%.4f", asset.Balance), // 保留4位小数
		Currency: asset.Currency,
	}, nil
}
