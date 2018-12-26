package multichain

import (
	"github.com/mitchellh/mapstructure"
)

// GetPeerInfo is a struct representing the result from the multichain.GetInfo() RPC Command
type GetPeerInfo struct {
	Result []struct {
		ID             int           `json:"id"`
		Addr           string        `json:"addr"`
		Addrlocal      string        `json:"addrlocal"`
		Services       string        `json:"services"`
		Lastsend       int           `json:"lastsend"`
		Lastrecv       int           `json:"lastrecv"`
		Bytessent      int           `json:"bytessent"`
		Bytesrecv      int           `json:"bytesrecv"`
		Conntime       int           `json:"conntime"`
		Pingtime       float64       `json:"pingtime"`
		Version        int           `json:"version"`
		Subver         string        `json:"subver"`
		Inbound        bool          `json:"inbound"`
		Startingheight int           `json:"startingheight"`
		Banscore       int           `json:"banscore"`
		SyncedHeaders  int           `json:"synced_headers"`
		SyncedBlocks   int           `json:"synced_blocks"`
		Inflight       []interface{} `json:"inflight"`
		Whitelisted    bool          `json:"whitelisted"`
	} `json:"result"`
	error interface{} `json:"error"`
	id    string      `json:"id"`
}

// ParseResponse takes a valid response and parses it into the model
func (m *GetPeerInfo) ParseResponse(r Response) {
	err := mapstructure.Decode(r, &m)
	if err != nil {
		panic(err)
	}
}

// GetPeerInfo Returns data about each connected network node as a json array of objects.
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
//	"id": "GoWAF-MultiChain-RPC-Client",
//	"jsonrpc": "1.0",
//	"method": "getpeerinfo",
//	"params": []
//}
//
// JSON Response:
//```
//Result:
//[
//  {
//    "id": n,						(numeric) Peer index
//    "addr":"host:port",      		(string) The IP address and port of the peer
//    "addrbind":"ip:port",			(string) Bind address of the connection to the peer
//    "addrlocal":"ip:port",		(string) Local address as reported by the peer
//    "services":"xxxxxxxxxxxxxxxx",(string) The services offered
//    "relaytxes":true|false,		(boolean) Whether peer has asked us to relay transactions to it
//    "lastsend": ttt,				(numeric) The time in seconds since epoch (Jan 1 1970 GMT) of the last send
//    "lastrecv": ttt,				(numeric) The time in seconds since epoch (Jan 1 1970 GMT) of the last receive
//    "bytessent": n,				(numeric) The total bytes sent
//    "bytesrecv": n,				(numeric) The total bytes received
//    "conntime": ttt,				(numeric) The connection time in seconds since epoch (Jan 1 1970 GMT)
//    "timeoffset": ttt,			(numeric) The time offset in seconds
//    "pingtime": n,             (numeric) ping time (if available)
//    "minping": n,              (numeric) minimum observed ping time (if any at all)
//    "pingwait": n,             (numeric) ping wait (if non-zero)
//    "version": v,              (numeric) The peer version, such as 7001
//    "subver": "/Satoshi:0.8.5/",  (string) The string version
//    "inbound": true|false,     (boolean) Inbound (true) or Outbound (false)
//    "addnode": true|false,     (boolean) Whether connection was due to addnode/-connect or if it was an automatic/inbound connection
//    "startingheight": n,       (numeric) The starting height (block) of the peer
//    "banscore": n,             (numeric) The ban score
//    "synced_headers": n,       (numeric) The last header we have in common with this peer
//    "synced_blocks": n,        (numeric) The last block we have in common with this peer
//    "inflight": [
//       n,                        (numeric) The heights of blocks we're currently asking from this peer
//       ...
//    ],
//    "whitelisted": true|false, (boolean) Whether the peer is whitelisted
//    "bytessent_per_msg": {
//       "addr": n,              (numeric) The total bytes sent aggregated by message type
//       ...
//    },
//    "bytesrecv_per_msg": {
//       "addr": n,              (numeric) The total bytes received aggregated by message type
//       ...
//    }
//  }
//  ,...
//]
//```
func (client *Client) GetPeerInfo() (Response, error) {

	msg := client.Command(
		"getpeerinfo",
		[]interface{}{},
	)

	return client.Post(msg)
}
