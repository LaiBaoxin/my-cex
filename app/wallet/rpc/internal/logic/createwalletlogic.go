package logic

import (
	"context"
	"fmt"
	"github.com/wwater/my-cex/app/wallet/rpc/internal/model"

	"github.com/wwater/my-cex/app/wallet/rpc/internal/svc"
	"github.com/wwater/my-cex/app/wallet/rpc/wallet"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWalletLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWalletLogic {
	return &CreateWalletLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateWalletLogic) CreateWallet(in *wallet.CreateWalletReq) (*wallet.CreateWalletResp, error) {
	// 先查数据库，看这个人有没有钱包
	var asset model.UserAsset
	// 查询是否存在
	err := l.svcCtx.DB.Where("uid = ? AND currency = ?", in.Uid, "ETH").First(&asset).Error

	// 如果存在就直接返回
	if err == nil {
		return &wallet.CreateWalletResp{
			Address: asset.Address,
		}, nil
	}

	// 如果没找到，生成新钱包
	password := fmt.Sprintf("wwater_%d", in.Uid)

	// Keystore 生成新账户
	account, err := l.svcCtx.KeyStore.NewAccount(password)
	if err != nil {
		l.Logger.Errorf("创建钱包失败: %v", err)
		return nil, err
	}

	// 获取生成的0x地址
	address := account.Address.Hex()

	// 创建存储的数据
	data := model.UserAsset{
		Uid:      in.Uid,
		Address:  address,
		Currency: "ETH",
		Balance:  0,
		Status:   1,
	}

	// 保存
	if err = l.svcCtx.DB.Create(&data).Error; err != nil {
		l.Logger.Errorf("[CreateWallet] 保存用户钱包数据库失败: %v", err)
		return nil, err
	}

	return &wallet.CreateWalletResp{
		Address: address,
	}, nil
}
