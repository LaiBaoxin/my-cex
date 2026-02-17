# My-CEX 资产管理与链上审计系统

本项目是一个打通 Web2 业务逻辑与 Web3 链上结算的全栈演示系统。包含智能合约开发、Go 微服务后端、以及 Vue3 实时监控前端。

## 技术栈
- **Smart Contract**: Solidity, Foundry (Forge/Cast)
- **Backend**: Go, Go-Zero, Go-Ethereum, GORM
- **Database**: MySQL, ClickHouse, Etcd
- **Frontend**: Vue3, Element Plus, WebSocket


## 快速启动指南

### 1. 环境准备
确保本地已安装：
- [Foundry](https://book.getfoundry.sh/getting-started/installation)
- [Go](https://go.dev/dl/) (1.20+)
- [Docker & Docker Compose](https://www.docker.com/)

### 2. 启动本地私链
```bash
# 启动 Anvil (模拟以太坊环境) foundry
anvil
```

### 3. 智能合约部署与 ABI 生成
```bash
# 部署合约到本地
forge create src/MyToken.sol:MyToken --rpc-url [http://127.0.0.1:8545](http://127.0.0.1:8545) --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# 提取并生成 Go 绑定代码 (abigen)
# 先从 out/MyToken.sol/MyToken.json 提取 abi 字段
cat out/MyToken.sol/MyToken.json | jq .abi > MyToken.abi
abigen --abi MyToken.abi --pkg mytoken --type MyToken --out ../common/contracts/mytoken/MyToken.go
```

### 4. 后端代码生成 (go-zero)
```bash
# Account API 生成
cd app/account/api
goctl api go -api account.api -dir . -style go_style

# Wallet RPC 生成
cd app/wallet/rpc
goctl rpc protoc wallet.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. -style go_style

```