package api

var (
	// TransactionAPI 交易相关接口
	TransactionAPI = transaction{
		GetTransactionInfoByBlockNum: "/walletsolidity/gettransactioninfobyblocknum", // 获取区块内的交易信息
	}

	// BlockAPI 区块相关接口
	BlockAPI = block{
		GetNowBlock: "/walletsolidity/getnowblock", // 获取当前区块高度
	}

	// AccountAPI 账户相关接口
	AccountAPI = account{
		GetAccount: "/walletsolidity/getaccount", // 获取区块账户
	}
)

type account struct {
	GetAccount string
}

type transaction struct {
	GetTransactionInfoByBlockNum string
}

type block struct {
	GetNowBlock string
}
