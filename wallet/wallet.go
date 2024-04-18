package wallet

import (
	"crypto/ecdsa"
	"strings"
	"tron-go-sdk/address"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

type Wallet struct {
	address.AccountAddress
	privateKey *ecdsa.PrivateKey
}

// PrivateKeyHex 获取私钥的十六进制字符串表示
func (w *Wallet) PrivateKeyHex() string {
	privateKeyBytes := crypto.FromECDSA(w.privateKey)
	return hexutil.Encode(privateKeyBytes)
}

func (w *Wallet) Address() address.AccountAddress {
	return w.AccountAddress
}

// GenerateKeyPair 生成账户地址和私钥
func GenerateKeyPair() (address, pk string, err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()[2:]
	return address, hexutil.Encode(privateKeyBytes)[2:], nil
}

// CreateNewWallet 创建全新的钱包
func CreateNewWallet() (Wallet, error) {
	_, pk, err := GenerateKeyPair()
	if err != nil {
		return Wallet{}, err
	}
	return ImportWalletFromPrivateKeyStr(pk)
}

// ImportWalletFromPrivateKeyStr 从私钥生成钱包
func ImportWalletFromPrivateKeyStr(privateKey string) (Wallet, error) {
	if !strings.HasPrefix(privateKey, "0x") {
		privateKey = "0x" + privateKey
	}
	privateKeyBytes, err := hexutil.Decode(privateKey)
	if err != nil {
		return Wallet{}, errors.Errorf("decode private to hex bytes err:%s", err)
	}
	pk, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return Wallet{}, errors.Errorf("convert to ecdsa failed: %s", err)
	}
	return ImportWalletFromPrivateKey(pk)
}

// ImportWalletFromPrivateKey 从 ecdsa 创建钱包
func ImportWalletFromPrivateKey(privateKey *ecdsa.PrivateKey) (Wallet, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return Wallet{}, errors.Errorf("invalid privateKey, error casting public key to ECDSA")
	}
	rawNetworkAddr := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()[2:]
	networkAddr := "41" + rawNetworkAddr
	addr, err := address.NewAccountAddressFromNetWorkAddr(networkAddr)
	if err != nil {
		return Wallet{}, err
	}
	return Wallet{
		AccountAddress: addr,
		privateKey:     privateKey,
	}, nil
}
