package transaction

import "context"

// Subscriber 交易订阅
type Subscriber interface {
}

// AddressChecker 地址条件检查
type AddressChecker interface {
	Support(ctx context.Context, fromAddress, toAddress string) bool
}

type FilterCond struct {
	Type        string   // 交易类型
	ToAddresses []string // 目标账户

}

type Filter interface {
}
