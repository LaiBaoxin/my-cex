package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/wwater/my-cex/app/wallet/rpc/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"

	"github.com/wwater/my-cex/app/wallet/rpc/internal/svc"
	"github.com/wwater/my-cex/app/wallet/rpc/wallet"

	"github.com/zeromicro/go-zero/core/logx"
)

type WithdrawLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWithdrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawLogic {
	return &WithdrawLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Withdraw 提现
func (l *WithdrawLogic) Withdraw(in *wallet.WithdrawReq) (*wallet.WithdrawResp, error) {
	amount, err := strconv.ParseFloat(in.Amount, 64)
	if err != nil || amount <= 0 {
		return nil, errors.New("无效的金额")
	}

	// 模拟生成一个链上 Hash
	fakeTxHash := fmt.Sprintf("0x%d_fake_hash", time.Now().UnixNano())

	// 开启进行提现事务
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {

		var exists int64
		tx.Model(&model.UserTransaction{}).Where("tx_id = ?", in.TxId).Count(&exists)
		if exists > 0 { // 已经处理过了，视为成功
			return nil
		}

		// 悲观锁锁住这个用户的余额行，防止别人同时修改
		var asset model.UserAsset
		if err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("uid = ? AND currency = ?", in.Uid, "ETH").
			First(&asset).Error; err != nil {
			return errors.New("用户钱包不存在")
		}

		// 检查余额是否充足
		if asset.Balance < amount {
			return errors.New("余额不足")
		}

		// 扣除余额
		asset.Balance -= amount
		if err = tx.Save(&asset).Error; err != nil {
			return err
		}

		// 插入提现流水记录
		newTx := model.UserTransaction{
			TxId:   in.TxId,
			Uid:    in.Uid,
			Amount: amount,
			Type:   2, // 2: Withdraw
			Status: 2, // Success
		}
		if err = tx.Create(&newTx).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		l.Logger.Errorf("Withdraw failed: %v", err)
		return &wallet.WithdrawResp{Success: false}, err
	}

	return &wallet.WithdrawResp{
		Success: true,
		TxHash:  fakeTxHash,
	}, nil
}
