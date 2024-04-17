package main

import (
	"context"
	"testing"
)

func TestCheckBlockTrans(t *testing.T) {
	ctx := context.Background()
	CheckBlockTrans(ctx, []string{"a557403d8f3ef634e6db37906ef4b81ef9fee1f871c119bfcec0b982e9f6abb1"})
}
