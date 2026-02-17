package config

// BlockchainConfig 链相关的通用配置
type BlockchainConfig struct {
	RpcUrl          string
	ContractAddress string
	AdminPrivateKey string
	ChainID         int64
}
