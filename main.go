package main

import (
	"context"
	"fmt"
	"time"
	"tron-go-sdk/httpapi"
)

var (
	lastBlockNum = 0
)

func CheckBlockTrans(ctx context.Context, blockNum int, txIDs []string) {
	client := httpapi.NewClient(httpapi.NetworkShasta)
	for i := range txIDs {
		transactionResp, err := client.GetTransactionByID(ctx, httpapi.GetTransactionByIDRequest{Value: txIDs[i]})
		if err != nil {
			fmt.Printf("GetTransactionByID err:%s\n", err)
			continue
		}
		isTransfer := false
		ownerAddress := ""
		toAddress := ""
		amount := 0
		for j := range transactionResp.RawData.Contract {
			if transactionResp.RawData.Contract[j].Type == "TransferContract" {
				isTransfer = true
				ownerAddress = transactionResp.RawData.Contract[j].Parameter.Value.OwnerAddress
				toAddress = transactionResp.RawData.Contract[j].Parameter.Value.ToAddress
				amount = transactionResp.RawData.Contract[j].Parameter.Value.Amount
			}
		}
		if isTransfer {
			fmt.Printf("block:%d,owner_address:%s,to_address:%s,amount:%d\n", blockNum, ownerAddress, toAddress, amount)
		}
	}
}

func checkBlock(ctx context.Context, blockNum int) {
	client := httpapi.NewClient(httpapi.NetworkShasta)
	resp, err := client.GetBlockByNum(ctx, httpapi.GetBlockByNumRequest{Num: blockNum})
	if err != nil {
		fmt.Printf("GetBlockByNum err:%s\n", err)
		return
	}
	isTransfer := false
	ownerAddress := ""
	toAddress := ""
	amount := 0
	for i := range resp.Transactions {
		for j := range resp.Transactions[i].RawData.Contract {
			if resp.Transactions[i].RawData.Contract[j].Type == "TransferContract" {
				isTransfer = true
				ownerAddress = resp.Transactions[i].RawData.Contract[j].Parameter.Value.OwnerAddress
				toAddress = resp.Transactions[i].RawData.Contract[j].Parameter.Value.ToAddress
				amount = resp.Transactions[i].RawData.Contract[j].Parameter.Value.Amount
			}
		}
	}
	if isTransfer {
		fmt.Printf("block:%d,owner_address:%s,to_address:%s,amount:%d\n", blockNum, ownerAddress, toAddress, amount)
	}
}

func main() {
	ch := make(chan struct{}, 5)
	ctx := context.Background()
	client := httpapi.NewClient(httpapi.NetworkShasta)
	for {
		resp, err := client.GetNowBlock(ctx)
		if err != nil {
			fmt.Printf("GetNowBlock err:%s\n", err)
			continue
		}
		if lastBlockNum > 0 && resp.BlockHeader.RawData.Number == lastBlockNum {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if resp.BlockHeader.RawData.Number > lastBlockNum+1 && lastBlockNum > 0 { // 非自增漏块情况
			for missingBlockNum := lastBlockNum + 1; missingBlockNum < resp.BlockHeader.RawData.Number; missingBlockNum++ {
				fmt.Printf("missing block:%d\n", missingBlockNum)
				ch <- struct{}{}
				go func(bn int) {
					checkBlock(ctx, bn)
					<-ch
				}(missingBlockNum)
			}
		}
		lastBlockNum = resp.BlockHeader.RawData.Number
		fmt.Printf("new block:%d\n", lastBlockNum)

		txIDs := make([]string, 0, len(resp.Transactions))
		for _, tx := range resp.Transactions {
			txIDs = append(txIDs, tx.TxID)
		}
		if len(txIDs) > 0 {
			CheckBlockTrans(ctx, lastBlockNum, txIDs)
		}
	}
}
