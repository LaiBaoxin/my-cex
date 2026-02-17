package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/wwater/my-cex/app/wallet/rpc/internal/model"
	"github.com/wwater/my-cex/app/wallet/rpc/internal/svc"
	"github.com/wwater/my-cex/app/wallet/rpc/wallet"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"math/big"
	"strconv"

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
	amountFloat, err := strconv.ParseFloat(in.Amount, 64)
	if err != nil || amountFloat <= 0 {
		return nil, errors.New("无效的提现金额")
	}

	// 计算链上大数 (Wei): amount * 10^18
	amountInt := new(big.Int)
	famount := new(big.Float).Mul(big.NewFloat(amountFloat), big.NewFloat(math.Pow10(18)))
	famount.Int(amountInt)

	var txHash string

	err = l.svcCtx.DB.Transaction(func(dbTx *gorm.DB) error {
		var exists int64
		dbTx.Model(&model.UserTransaction{}).Where("tx_id = ?", in.TxId).Count(&exists)
		if exists > 0 {
			return nil
		}

		// 锁定并查询余额
		var asset model.UserAsset
		if err = dbTx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("uid = ? AND currency = ?", in.Uid, "ETH").
			First(&asset).Error; err != nil {
			return errors.New("用户钱包不存在")
		}

		if asset.Balance < amountFloat {
			return errors.New("余额不足")
		}

		privateKey, _ := crypto.HexToECDSA(l.svcCtx.Config.Blockchain.AdminPrivateKey)
		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(l.svcCtx.Config.Blockchain.ChainID))
		if err != nil {
			return fmt.Errorf("构造签名器失败: %v", err)
		}

		// 调用合约的 Transfer 函数
		targetAddr := common.HexToAddress(asset.Address)
		blockchainTx, err := l.svcCtx.TokenSvc.Transfer(auth, targetAddr, amountInt)
		if err != nil {
			return fmt.Errorf("区块链转账失败: %v", err)
		}
		txHash = blockchainTx.Hash().Hex()

		// 扣除数据库余额
		asset.Balance -= amountFloat
		if err := dbTx.Save(&asset).Error; err != nil {
			return err
		}

		// 创建流水并保存真实的链上 Hash
		newTx := model.UserTransaction{
			TxId:   in.TxId,
			Uid:    in.Uid,
			Amount: amountFloat,
			Type:   2, // 2: Withdraw
			Status: 2, // 成功
		}
		if err := dbTx.Create(&newTx).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		l.Logger.Errorf("提现流程失败: %v", err)
		return &wallet.WithdrawResp{Success: false}, err
	}

	return &wallet.WithdrawResp{
		Success: true,
		TxHash:  txHash,
	}, nil
}
