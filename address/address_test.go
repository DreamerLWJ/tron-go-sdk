package address

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	_testBase58Addr  = "TDKXrDK51xHQ33m1HSptJ2V3Xo1mcoay2j"
	_testNetworkAddr = "4124c06dff283762d481fddb3ee693765c88ee1266"
)

func TestNewAccountAddressFromBase58Addr(t *testing.T) {
	addr, err := NewAccountAddressFromBase58Addr(_testBase58Addr)
	assert.Nil(t, err)
	t.Logf("networkAddr: %s", addr.NetworkAddr())
	assert.Equal(t, _testBase58Addr, addr.Base58Addr())
	assert.Equal(t, _testNetworkAddr, addr.NetworkAddr())
}

func TestNewAccountAddressFromNetWorkAddr(t *testing.T) {
	addr, err := NewAccountAddressFromNetWorkAddr(_testNetworkAddr)
	assert.Nil(t, err)
	t.Logf("base58Addr: %s", addr.Base58Addr())
	assert.Equal(t, _testBase58Addr, addr.Base58Addr())
	assert.Equal(t, _testNetworkAddr, addr.NetworkAddr())
}
