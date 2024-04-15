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
	GetTransactionByID = "/walletsolidity/gettransactionbyid" // 根据交易 ID 获取交易信息
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
