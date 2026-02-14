package logic

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wwater/my-cex/app/wallet/rpc/internal/model"
	"math/big"

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
	// 获取数据库中的资产信息
	var asset model.UserAsset
	err := l.svcCtx.DB.Where("uid = ? AND currency = ?", in.Uid, in.Currency).First(&asset).Error
	if err != nil {
		return nil, fmt.Errorf("资产记录不存在")
	}

	// 从链上合约查询实时余额
	userAddr := common.HexToAddress(asset.Address)
	chainBalance, err := l.svcCtx.TokenSvc.BalanceOf(nil, userAddr)
	if err != nil {
		l.Logger.Errorf("链上余额查询失败: %v", err)
		// 容错处理：链上查不到则返回数据库金额
		return &wallet.GetBalanceResp{
			Balance: fmt.Sprintf("%.4f", asset.Balance),
			Address: asset.Address,
		}, nil
	}

	fbalance := new(big.Float).SetInt(chainBalance)
	divisor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	ethValue := new(big.Float).Quo(fbalance, divisor)

	resultStr := ethValue.Text('f', 4)

	return &wallet.GetBalanceResp{
		Balance: resultStr,
		Address: asset.Address,
	}, nil
}
