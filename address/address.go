package address

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/btcsuite/btcutil/base58"
	"github.com/pkg/errors"
)

type AccountAddress struct {
	networkAddr string // 网路上账户地址表示
	base58Addr  string // Base58 编码表示，用户钱包表示
}

// NetworkAddr 获取网络表示地址
func (a AccountAddress) NetworkAddr() string {
	return a.networkAddr
}

// Base58Addr 获取 Base58 地址
func (a AccountAddress) Base58Addr() string {
	return a.base58Addr
}

func NewAccountAddressFromNetWorkAddr(networkAddr string) (AccountAddress, error) {
	decodeAddress, err := hex.DecodeString(networkAddr)
	if err != nil {
		return AccountAddress{}, errors.Errorf("hex decode err:%s", err)
	}
	sha2560 := sha256.Sum256(decodeAddress)
	sha2561 := sha256.Sum256(sha2560[:])
	checkSum := sha2561[0:4]
	addCheckSum := append(decodeAddress, checkSum...)
	return AccountAddress{
		networkAddr: networkAddr,
		base58Addr:  base58.Encode(addCheckSum),
	}, nil
}

func NewAccountAddressFromBase58Addr(base58Addr string) (AccountAddress, error) {
	networkAddrAndCheckSumHex := base58.Decode(base58Addr)
	if len(networkAddrAndCheckSumHex) < 4 {
		return AccountAddress{}, errors.Errorf("invalid base58 str:%s", base58Addr)
	}
	networkAddrHex := networkAddrAndCheckSumHex[0 : len(networkAddrAndCheckSumHex)-4]
	return AccountAddress{
		networkAddr: hex.EncodeToString(networkAddrHex),
		base58Addr:  base58Addr,
	}, nil
}
