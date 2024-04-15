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
