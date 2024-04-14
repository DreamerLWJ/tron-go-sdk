package api

type GetNowBlockRequest struct {
	BlockID     string      `json:"blockID"`
	BlockHeader BlockHeader `json:"block_header"`
}

type RawData struct {
	Number         int    `json:"number"`
	TxTrieRoot     string `json:"txTrieRoot"`
	WitnessAddress string `json:"witness_address"`
	ParentHash     string `json:"parentHash"`
	Version        int    `json:"version"`
	Timestamp      int64  `json:"timestamp"`
}

type BlockHeader struct {
	RawData          RawData `json:"raw_data"`
	WitnessSignature string  `json:"witness_signature"`
}
