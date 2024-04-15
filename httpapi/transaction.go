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
	GetTransactionByID       = "/walletsolidity/gettransactionbyid"           // 根据交易 ID 获取交易信息
	GetTransactionByBlockNum = "/walletsolidity/gettransactioninfobyblocknum" // 根据区块高度获取交易信息
)

type GetTransactionByIDRequest struct {
	Value string `json:"value"` // 交易 ID
}

type GetTransactionByIDResponse struct {
	Transaction
}

func (c *Client) GetTransactionByID(ctx context.Context, req GetTransactionByIDRequest) (resp GetTransactionByIDResponse, err error) {
	parseUrl, err := url.Parse(c.HTTPApiHost)
	if err != nil {
		return resp, err
	}
	parseUrl.Path = GetTransactionByID

	reqBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}
	response, err := http.Post(parseUrl.String(), "application/json", bytes.NewBuffer(reqBody))
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

type GetTransactionByBlockNumRequest struct {
	Num uint64 `json:"num"` // 区块高度
}

type GetTransactionByBlockNumResponseItem struct {
	Fee uint64 `json:"fee"`
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
	Id                   string `json:"id"`
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

// GetTransactionByBlockNum 根据区块高度获取交易，如果需要获取更详细的交易信息，则需要 GetTransactionByID
func (c *Client) GetTransactionByBlockNum(ctx context.Context, req GetTransactionByBlockNumRequest) (res []GetTransactionByBlockNumResponseItem, err error) {
	parseUrl, err := url.Parse(c.HTTPApiHost)
	if err != nil {
		return res, err
	}
	parseUrl.Path = GetTransactionByBlockNum
	data, err := json.Marshal(req)
	if err != nil {
		return res, err
	}
	response, err := http.Post(parseUrl.String(), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return res, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return res, err
	}
	if err = json.Unmarshal(body, &res); err != nil {
		return res, err
	}
	return res, nil
}
