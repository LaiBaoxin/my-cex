package main

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wwater/my-cex/common/contracts/mytoken"
)

func TestContractClient(t *testing.T) {
	// 连接到本地 Anvil 节点
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		t.Fatalf("无法连接到 Anvil: %v", err)
	}

	// 部署合约地址Contract Address: 0x5FbDB...
	address := common.HexToAddress("0x5FbDB")

	// 创建合约实例
	instance, err := mytoken.NewMyToken(address, client)
	if err != nil {
		t.Fatalf("合约实例化失败: %v", err)
	}

	// 只读调用合约
	name, err := instance.Name(nil)
	if err != nil {
		t.Fatalf("查询名称失败: %v", err)
	}

	fmt.Printf("连接合约成功！代币名称为: %s\n", name)
}
