package api

type GetTransactionInfoByBlockNumContactResponse struct {
	Log []struct {
		Address string   `json:"address"`
		Topics  []string `json:"topics"`
	} `json:"log"`
	BlockNumber    int      `json:"blockNumber"`
	ContractResult []string `json:"contractResult"`
	BlockTimeStamp int64    `json:"blockTimeStamp"`
	Receipt        struct {
		Result           string `json:"result"`
		EnergyUsage      int    `json:"energy_usage"`
		EnergyUsageTotal int    `json:"energy_usage_total"`
		NetUsage         int    `json:"net_usage"`
	} `json:"receipt"`
	ID                   string `json:"id"`
	ContractAddress      string `json:"contract_address"`
	InternalTransactions []struct {
		CallerAddress     string `json:"caller_address"`
		Note              string `json:"note"`
		TransferToAddress string `json:"transferTo_address"`
		CallValueInfo     []struct {
			CallValue int64 `json:"callValue"`
		} `json:"callValueInfo"`
		Hash string `json:"hash"`
	} `json:"internal_transactions"`
}

type GetTransactionInfoByBlockNumTransferResponse struct {
	BlockNumber    int      `json:"blockNumber"`
	ContractResult []string `json:"contractResult"`
	BlockTimeStamp int64    `json:"blockTimeStamp"`
	Receipt        struct {
		NetUsage int `json:"net_usage"`
	} `json:"receipt"`
	ID string `json:"id"`
}
