package wallet

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

const (
	_testPrivateKey  = "0xe9688b4b54916d43817ddb0de36ef226c5dba60f39d6efc990e24fbe9f166588"
	_testPrivateKey2 = "e9688b4b54916d43817ddb0de36ef226c5dba60f39d6efc990e24fbe9f166588"
)

func TestGenerateKeyPair(t *testing.T) {
	address, pk, err := GenerateKeyPair()
	assert.Nil(t, err)
	t.Logf("address: %s, pk: %s", address, pk)
}

func TestCreateNewWallet(t *testing.T) {
	wallet, err := CreateNewWallet()
	assert.Nil(t, err)
	assert.NotEmpty(t, wallet.AccountAddress.NetworkAddr())
	assert.NotEmpty(t, wallet.AccountAddress.Base58Addr())
	assert.NotEmpty(t, wallet.PrivateKeyHex())
	t.Logf("address: %s, pk :%s", wallet.AccountAddress.NetworkAddr(), wallet.PrivateKeyHex())
}

func TestImportWalletFromPrivateKey(t *testing.T) {
	privateKeyBytes, err := hexutil.Decode(_testPrivateKey)
	assert.Nil(t, err)
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	assert.Nil(t, err)
	wallet, err := ImportWalletFromPrivateKey(privateKey)
	assert.Nil(t, err)
	assert.NotEmpty(t, wallet.AccountAddress.NetworkAddr())
	assert.NotEmpty(t, wallet.AccountAddress.Base58Addr())
	assert.NotEmpty(t, wallet.PrivateKeyHex())
	assert.Equal(t, _testPrivateKey, wallet.PrivateKeyHex())
	t.Logf("address: %s, pk :%s", wallet.AccountAddress.NetworkAddr(), wallet.PrivateKeyHex())
}

func TestImportWalletFromPrivateKeyStr(t *testing.T) {
	_, err := ImportWalletFromPrivateKeyStr(_testPrivateKey2)
	assert.Nil(t, err)
	_, err = ImportWalletFromPrivateKeyStr(_testPrivateKey)
	assert.Nil(t, err)
}
