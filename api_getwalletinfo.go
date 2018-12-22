package multichain

import (
	"github.com/NlaakStudiosLLC/GoWAF/framework/logger"
	"github.com/mitchellh/mapstructure"
)

// GetWalletInfo is a struct representing the result from the multichain.GetInfo() RPC Command
type GetWalletInfo struct {
	Result struct {
		Walletversion   int     `json:"walletversion"`
		Balance         float64 `json:"balance"`
		Walletdbversion int     `json:"walletdbversion"`
		Txcount         int     `json:"txcount"`
		Utxocount       int     `json:"utxocount"`
		Keypoololdest   int     `json:"keypoololdest"`
		Keypoolsize     int     `json:"keypoolsize"`
	} `json:"result"`
	error interface{} `json:"error"`
	id    string      `json:"id"`
}

// ParseResponse takesa valid response and parses it into the model
func (m *GetWalletInfo) ParseResponse(r Response) {
	err := mapstructure.Decode(m, &r)
	if err != nil {
		logger.LogThis.Error(err)
	}
}

// GetWalletInfo Returns an object containing various wallet state info.
//```
//{
//	"walletversion": xxxxx,       (numeric) the wallet version
//	"balance": xxxxxxx,           (numeric) the total confirmed balance of the wallet in BTC
//	"unconfirmed_balance": xxx,   (numeric) the total unconfirmed balance of the wallet in BTC
//	"immature_balance": xxxxxx,   (numeric) the total immature balance of the wallet in BTC
//	"txcount": xxxxxxx,           (numeric) the total number of transactions in the wallet
//	"keypoololdest": xxxxxx,      (numeric) the timestamp (seconds since Unix epoch) of the oldest pre-generated key in the key pool
//	"keypoolsize": xxxx,          (numeric) how many new keys are pre-generated
//	"unlocked_until": ttt,        (numeric) the timestamp in seconds since epoch (midnight Jan 1 1970 GMT) that the wallet is unlocked for transfers, or 0 if the wallet is locked
//	"paytxfee": x.xxxx,           (numeric) the transaction fee configuration, set in BTC/kB
//	"hdmasterkeyid": "<hash160>", (string) the Hash160 of the HD master pubkey
//}
//```
// JSON-RPC Request:
//```
//{
//	"id": "MultiChain-RPC-Client",
//	"jsonrpc": "1.0",
//	"method": "getwalletinfo",
//	"params": []
//}
//
// JSON Response:
//```
//{
//	"result": {
//		"walletversion": 60000,
//		"balance": 20279000000.00000000,
//		"walletdbversion": 2,
//		"txcount": 20690,
//		"utxocount": 20279,
//		"keypoololdest": 1544843363,
//		"keypoolsize": 101
//	},
//	"error": null,
//	"id": "MultiChain-RPC-Client"
//}
//```
func (client *Client) GetWalletInfo() (Response, error) {

	msg := client.Command(
		"getwalletinfo",
		[]interface{}{},
	)

	return client.Post(msg)
}
