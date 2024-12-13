package admin

type ChainInfo struct {
	Name    string `json:"name"`
	RPC     string `json:"rpc"`
	ChainID uint64 `json:"chainID"`
	Type    string `json:"type"`
}
