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

func CheckBlockTrans(ctx context.Context, txIDs []string) {
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
		for j := range transactionResp.RawData.Contract {
			if transactionResp.RawData.Contract[j].Type == "TransferContract" {
				isTransfer = true
				ownerAddress = transactionResp.RawData.Contract[j].Parameter.Value.OwnerAddress
				toAddress = transactionResp.RawData.Contract[j].Parameter.Value.ToAddress
			}
		}
		if isTransfer {
			fmt.Printf("owner_address:%s,to_address:%s\n", ownerAddress, toAddress)
		}
	}
}

func checkBlock(blockNum int) {
	//client := httpapi.NewClient(httpapi.NetworkShasta)

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
		if resp.BlockHeader.RawData.Number > lastBlockNum+1 { // 非自增漏块情况
			for missingBlockNum := resp.BlockHeader.RawData.Number + 1; missingBlockNum <= lastBlockNum; missingBlockNum++ {
				fmt.Printf("missing block:%d\n", missingBlockNum)
				ch <- struct{}{}
				go func(bn int) {
					checkBlock(bn)
					<-ch
				}(missingBlockNum)
			}
		}
		lastBlockNum = resp.BlockHeader.RawData.Number
		fmt.Printf("new block:%d\n", lastBlockNum)

		txIDs := make([]string, 0, len(resp.Transaction))
		for _, tx := range resp.Transaction {
			txIDs = append(txIDs, tx.TxID)
		}
		if len(txIDs) > 0 {
			CheckBlockTrans(ctx, txIDs)
		}
	}
}
