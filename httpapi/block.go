package httpapi

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	GetNowBlockUrl = "/walletsolidity/getnowblock" // 获取当前区块
)

type GetNowBlockResponse struct {
	BlockID     string        `json:"block_id"`
	BlockHeader BlockHeader   `json:"block_header"` // 区块头数据
	Transaction []Transaction `json:"transaction"`  // 交易数据
}

type BlockHeader struct {
	RawData struct {
		Number         int    `json:"number"`
		TxTrieRoot     string `json:"txTrieRoot"`
		WitnessAddress string `json:"witness_address"`
		ParentHash     string `json:"parentHash"`
		Version        int    `json:"version"`
		Timestamp      int64  `json:"timestamp"`
	} `json:"raw_data"`
	WitnessSignature string `json:"witness_signature"`
}

type Transaction struct {
	Ret []struct {
		ContractRet string `json:"contractRet"`
	} `json:"ret"`
	Signature []string `json:"signature"`
	TxID      string   `json:"txID"`
	RawData   struct {
		Contract []struct {
			Parameter struct {
				Value struct {
					Amount       int    `json:"amount"`
					OwnerAddress string `json:"owner_address"`
					ToAddress    string `json:"to_address"`
				} `json:"value"`
				TypeUrl string `json:"type_url"`
			} `json:"parameter"`
			Type string `json:"type"` // TransferContract 为转账合约
		} `json:"contract"`
		RefBlockBytes string `json:"ref_block_bytes"`
		RefBlockHash  string `json:"ref_block_hash"`
		Expiration    int64  `json:"expiration"`
		Timestamp     int64  `json:"timestamp"`
	} `json:"raw_data"`
	RawDataHex string `json:"raw_data_hex"`
}

// GetNowBlock 获取当前区块高度
func (c *Client) GetNowBlock(ctx context.Context) (resp GetNowBlockResponse, err error) {
	parseUrl, err := url.Parse(c.HTTPApiHost)
	if err != nil {
		return resp, err
	}
	parseUrl.Path = GetNowBlockUrl
	response, err := http.Post(parseUrl.String(), "application/json", bytes.NewBuffer([]byte{}))
	if err != nil {
		return resp, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return resp, err
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		return resp, err
	}
	return resp, nil
}
