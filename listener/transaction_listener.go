package listener

import (
	"sync"
	"tron-go-sdk/httpapi"

	"github.com/thoas/go-funk"
)

const (
	_defaultSaveRecentBlockCount = 200 // 默认保存近 200 个区块号
)

type TransactionListener struct {
	c *httpapi.Client
}

func NewTransactionListener() *TransactionListener {

}

type Setting struct {
	SaveRecentBlockCount int // 最近的区块号数，用于避免 Node 共识差异时导致的区块重复
}

type BlockListener struct {
	rcbnMutex             sync.RWMutex
	tail                  int
	recentCommitBlockNums []int // 近期提交过的区块号
	latestCommitBlockNum  int   // 最新提交的区块号

	errChan chan error
}

func (b *BlockListener) Bootstrap() error {
	// TODO
}

func (b *BlockListener) isRecentBlock(blockNum int) bool {
	b.rcbnMutex.RLock()
	defer func() {
		b.rcbnMutex.RUnlock()
	}()
	if funk.InInts(b.recentCommitBlockNums, blockNum) {
		return true
	}
	return false
}

// 提交区块号
func (b *BlockListener) commitBlock(blockNum int) {
	b.rcbnMutex.Lock()
	defer b.rcbnMutex.Unlock()
	if funk.InInts(b.recentCommitBlockNums, blockNum) {
		return // ignore commited block
	}
	b.recentCommitBlockNums[b.tail%len(b.recentCommitBlockNums)] = blockNum
	b.latestCommitBlockNum = blockNum
}

func NewBlockListener(c *httpapi.Client, setting Setting) *BlockListener {
	recentBlockCount := _defaultSaveRecentBlockCount
	if setting.SaveRecentBlockCount > 0 {
		recentBlockCount = setting.SaveRecentBlockCount
	}

	return &BlockListener{
		recentCommitBlockNums: make([]int, recentBlockCount),
	}
}
