package httpapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestClient() *Client {
	return NewClient(NetworkShasta)
}

func TestClient_GetNowBlock(t *testing.T) {
	ctx := context.Background()
	client := getTestClient()
	resp, err := client.GetNowBlock(ctx)
	assert.Nil(t, err)
	t.Log(resp)
}

func TestClient_GetBlockByNum(t *testing.T) {
	ctx := context.Background()
	client := getTestClient()
	resp, err := client.GetBlockByNum(ctx, GetBlockByNumRequest{Num: 43284519})
	assert.Nil(t, err)
	t.Log(resp)
}
