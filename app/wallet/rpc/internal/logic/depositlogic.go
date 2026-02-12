package logic

import (
	"context"
	"errors"
	"github.com/wwater/my-cex/app/wallet/rpc/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"

	"github.com/wwater/my-cex/app/wallet/rpc/internal/svc"
	"github.com/wwater/my-cex/app/wallet/rpc/wallet"

	"github.com/zeromicro/go-zero/core/logx"
)

type DepositLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDepositLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepositLogic {
	return &DepositLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Deposit 实现资金充值
func (l *DepositLogic) Deposit(in *wallet.DepositReq) (*wallet.DepositResp, error) {
	amount, err := strconv.ParseFloat(in.Amount, 64)
	if err != nil || amount <= 0 {
		return nil, errors.New("invalid amount")
	}

	// 开启事务处理资金的充值
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {

		// 检查流水号是否已存在
		var exists int64
		tx.Model(&model.UserTransaction{}).Where("tx_id = ?", in.TxId).Count(&exists)
		if exists > 0 { // 已经处理过了，直接返回成功
			return nil
		}

		// 悲观锁查询资产
		var asset model.UserAsset
		if err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("uid = ? AND currency = ?", in.Uid, "ETH").
			First(&asset).Error; err != nil {
			return errors.New("钱包不存在")
		}

		// 增加余额
		asset.Balance += amount
		if err = tx.Save(&asset).Error; err != nil {
			return err
		}

		// 插入流水记录
		newTx := model.UserTransaction{
			TxId:   in.TxId,
			Uid:    in.Uid,
			Amount: amount,
			Type:   1, // Deposit
			Status: 2, // Success
		}
		if err = tx.Create(&newTx).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		l.Logger.Errorf("[Deposit] 充值失败: %v", err)
		return &wallet.DepositResp{Success: false}, err
	}

	return &wallet.DepositResp{Success: true}, nil
}
