package svc

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wwater/my-cex/common/contracts/mytoken"
	"github.com/wwater/my-cex/common/db"
	"gorm.io/gorm"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/wwater/my-cex/app/wallet/rpc/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	KeyStore *keystore.KeyStore
	DB       *gorm.DB
	TokenSvc *mytoken.MyToken
	EthCli   *ethclient.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 私钥存放目录
	ksDir := "./keystore_data"
	if _, err := os.Stat(ksDir); os.IsNotExist(err) {
		_ = os.Mkdir(ksDir, 0700)
	}

	// 初始化 Geth Keystore
	ks := keystore.NewKeyStore(ksDir, keystore.StandardScryptN, keystore.StandardScryptP)

	// 初始化 MySQL 连接
	dbConn := db.InitMysql(c.DataSource)

	// 初始化 ETH 客户端
	client, err := ethclient.Dial(c.Blockchain.RpcUrl)
	if err != nil {
		panic("连接区块链失败: " + err.Error())
	}

	// 实例化 MyToken 合约
	contractAddr := common.HexToAddress(c.Blockchain.ContractAddress)
	tokenInstance, err := mytoken.NewMyToken(contractAddr, client)
	if err != nil {
		panic("合约实例化失败: " + err.Error())
	}

	return &ServiceContext{
		Config:   c,
		KeyStore: ks,
		DB:       dbConn,
		TokenSvc: tokenInstance,
		EthCli:   client,
	}
}
