package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/sha3"
)

var (
	lastBlockNum = 0
)

//func main() {
//getTransTest()
//ctx := context.Background()
//client := httpapi.NewClient(httpapi.NetworkShasta)
//
//for {
//	resp, err := client.GetNowBlock(ctx)
//	if err != nil {
//		fmt.Printf("GetNowBlock err:%s\n", err)
//		continue
//	}
//	if lastBlockNum > 0 && resp.BlockHeader.RawData.Number == lastBlockNum {
//		time.Sleep(500 * time.Millisecond)
//		continue
//	}
//	lastBlockNum = resp.BlockHeader.RawData.Number
//	fmt.Printf("new block:%d\n", lastBlockNum)
//
//	txIDs := make([]string, 0, len(resp.Transaction))
//	for _, tx := range resp.Transaction {
//		txIDs = append(txIDs, tx.TxID)
//	}
//
//	for i := range txIDs {
//		transactionResp, err := client.GetTransactionByID(ctx, httpapi.GetTransactionByIDRequest{Value: txIDs[i]})
//		if err != nil {
//			fmt.Printf("GetTransactionByID err:%s\n", err)
//			continue
//		}
//		isTransfer := false
//		ownerAddress := ""
//		toAddress := ""
//		for j := range transactionResp.RawData.Contract {
//			if transactionResp.RawData.Contract[j].Type == "TransferContract" {
//				isTransfer = true
//				ownerAddress = transactionResp.RawData.Contract[j].Parameter.Value.OwnerAddress
//				toAddress = transactionResp.RawData.Contract[j].Parameter.Value.ToAddress
//			}
//		}
//		if isTransfer {
//			fmt.Printf("owner_address:%s,to_address:%s\n", ownerAddress, toAddress)
//		}
//	}
//}
//}

func privateToPublicDemo(privateKey []byte) []byte {
	privKey := new(ecdsa.PrivateKey)
	privKey.Curve = elliptic.P256()
	privKey.D = new(big.Int).SetBytes(privateKey)
	privKey.PublicKey.X, privKey.PublicKey.Y = privKey.Curve.ScalarBaseMult(privateKey)
	return append(privKey.PublicKey.X.Bytes(), privKey.PublicKey.Y.Bytes()...)
}

func publicToAddressDemo(publicKey []byte) []byte {
	hash := sha3.Sum256(publicKey[1:])
	address := make([]byte, 20)
	copy(address, hash[11:])
	address[0] = 0x41 // Assuming WalletApi.getAddressPreFixByte() is 0x41 in hex
	return address
}

func addressToEncode58CheckDemo(input []byte) string {
	hash0 := sha3.Sum256(input)
	hash1 := sha3.Sum256(hash0[:])
	checkSum := hash1[:4]
	inputCheck := append(input, checkSum...)
	return base58.Encode(inputCheck)
}

func privateToAddress(privateKey []byte) string {
	var ecdsaPrivKey *ecdsa.PrivateKey
	var err error

	if privateKey == nil {
		ecdsaPrivKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		ecdsaPrivKey = new(ecdsa.PrivateKey)
		ecdsaPrivKey.PublicKey.Curve = elliptic.P256()
		ecdsaPrivKey.D = new(big.Int).SetBytes(privateKey)
		ecdsaPrivKey.PublicKey.X, ecdsaPrivKey.PublicKey.Y = ecdsaPrivKey.Curve.ScalarBaseMult(privateKey)
	}

	fmt.Printf("Private Key: %x\n", ecdsaPrivKey.D.Bytes())

	publicKey0 := append(ecdsaPrivKey.PublicKey.X.Bytes(), ecdsaPrivKey.PublicKey.Y.Bytes()...)
	publicKey1 := privateToPublicDemo(privateKey)
	if !bytes.Equal(publicKey0, publicKey1) {
		log.Fatal("publickey error")
	}
	fmt.Printf("Public Key: %x\n", publicKey0)

	address0 := publicToAddressDemo(publicKey0)
	address1 := publicToAddressDemo(publicKey1)
	if !bytes.Equal(address0, address1) {
		log.Fatal("address error")
	}
	fmt.Printf("Address: %x\n", address0)

	base58checkAddress0 := base58.Encode(address0)
	base58checkAddress1 := addressToEncode58CheckDemo(address0)
	if base58checkAddress0 != base58checkAddress1 {
		log.Fatal("base58checkAddress error")
	}

	return base58checkAddress1
}

func main() {
	privateKey := "F43EBCC94E6C257EDBE559183D1A8778B2D5A08040902C0F0A77A3343A1D0EA5"
	address := privateToAddress(hexToBytes(privateKey))
	fmt.Println("base58Address:", address)

	fmt.Println("================================================================\r\n")

	address = privateToAddress(nil)
	fmt.Println("base58Address:", address)
}

func hexToBytes(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
