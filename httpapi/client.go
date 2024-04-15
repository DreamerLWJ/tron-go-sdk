package httpapi

type Network string

const (
	NetworkMainNet Network = "mainnet" // 主网
	NetworkShasta  Network = "shasta"  // 测试网 shasta
)

type NetworkConfig struct {
	HTTPApi string
	JSONApi string
}

var NetworkDefaultCfg map[Network]NetworkConfig = map[Network]NetworkConfig{
	NetworkMainNet: {
		HTTPApi: "",
		JSONApi: "",
	},
	NetworkShasta: {
		HTTPApi: "https://api.shasta.trongrid.io",
		JSONApi: "https://api.shasta.trongrid.io/jsonrpc",
	},
}

type Client struct {
	HTTPApiHost string
	JSONApiHost string
	Network     Network
}

func NewClient(network Network) *Client {
	return &Client{
		HTTPApiHost: NetworkDefaultCfg[network].HTTPApi,
		JSONApiHost: NetworkDefaultCfg[network].JSONApi,
		Network:     network,
	}
}
