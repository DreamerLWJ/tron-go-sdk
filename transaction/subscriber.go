package transaction

import "context"

// Subscriber 交易订阅
type Subscriber interface {
}

// Checker 检查是否支持处理
type Checker interface {
	// SupportAddress 是否处理该地址
	SupportAddress(ctx context.Context, fromAddress, toAddress string) bool

	// SupportTransaction 是否处理该交易
	SupportTransaction(ctx context.Context, txID string) bool
}

type FilterCond struct {
	Type        string   // 交易类型
	ToAddresses []string // 目标账户
}

type Filter interface {
}
