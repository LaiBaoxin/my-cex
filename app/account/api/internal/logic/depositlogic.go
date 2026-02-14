// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/wwater/my-cex/app/account/api/internal/model"
	"github.com/wwater/my-cex/common/websocket"
	"math/big"
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
	// 获取用户地址
	var asset struct {
		Address string
	}
	if err = l.svcCtx.DB.Table("user_assets").Where("uid = ?", req.UserId).First(&asset).Error; err != nil {
		return nil, fmt.Errorf("用户未开通钱包: %v", err)
	}

	// 准备链上签名对象
	privateKey, _ := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	// 获取 ChainID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))
	if err != nil {
		return nil, fmt.Errorf("构造签名器失败: %v", err)
	}

	// 计算金额 (18 位精度)
	amountInt := new(big.Int)
	amountInt.SetString(req.Amount, 10)
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil) // 10^18
	finalAmount := new(big.Int).Mul(amountInt, exp)

	// 调用合约 Mint 铸币
	targetAddr := common.HexToAddress(asset.Address)
	tx, err := l.svcCtx.TokenSvc.Mint(auth, targetAddr, finalAmount)
	if err != nil {
		l.Logger.Errorf("链上铸币失败: %v", err)
		return nil, err
	}

	// 记录持久化日志
	logEntry := model.SystemLog{
		Uid:     req.UserId,
		OpType:  "DEPOSIT",
		Content: fmt.Sprintf("成功充值 %s MCT, Hash: %s", req.Amount, tx.Hash().Hex()[:10]+"..."),
	}
	l.svcCtx.DB.Create(&logEntry)

	// WebSocket 推送
	websocket.GlobalHub.Broadcast(map[string]interface{}{
		"type":    "DEPOSIT",
		"userId":  req.UserId,
		"message": fmt.Sprintf("UID:%d 充值成功", req.UserId),
		"time":    time.Now().Format("15:04:05"),
	})

	return &types.DepositResp{Ok: true}, nil
}
