package httpapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	_testTransactionID = "94219a5c674cc30290b1464a58e5c18511185676f14971cd5f81323263e63594"
)

func TestClient_GetTransactionByID(t *testing.T) {
	ctx := context.Background()
	client := getTestClient()
	resp, err := client.GetTransactionByID(ctx, GetTransactionByIDRequest{Value: _testTransactionID})
	assert.Nil(t, err)
	t.Log(resp)
}
