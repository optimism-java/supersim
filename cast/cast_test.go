package cast

import "testing"

func TestBalanceCommand(t *testing.T) {
	args := []string{"balance", "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", "--rpc-url", "http://127.0.0.1:9545"}
	RunCast(args...)
}

func TestOpPortal(t *testing.T) {
	args := []string{
		"send",
		"0x37a418800d0c812A9dE83Bc80e993A6b76511B57",
		"--value",
		"0.1ether",
		"--rpc-url",
		"http://127.0.0.1:8545",
		"--private-key",
		"0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	}
	RunCast(args...)
}
