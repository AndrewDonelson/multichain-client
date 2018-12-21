package multichain

import (
	"github.com/NlaakStudiosLLC/GoWAF/framework/logger"
	"github.com/mitchellh/mapstructure"
)

// GetInfo is a struct representing the result from the multichain.GetInfo() RPC Command
/*
{
	"result": {
		"version": "1.0.7",
		"nodeversion": 10006901,
		"protocolversion": 10011,
		"chainname": "chain1",
		"description": "some text",
		"protocol": "bitcoin",
		"port": 5002,
		"setupblocks": 50,
		"nodeaddress": "chain1@192.168.1.134:5002",
		"burnaddress": "1XXXXXXXXuXXXXXXLjXXXXXXSTXXXXXXYk8dDb",
		"incomingpaused": false,
		"miningpaused": false,
		"walletversion": 60000,
		"balance": 17788000000.00000000,
		"walletdbversion": 2,
		"reindex": false,
		"blocks": 30153,
		"timeoffset": 0,
		"connections": 3,
		"proxy": "",
		"difficulty": 0.01386980,
		"testnet": false,
		"keypoololdest": 1544843363,
		"keypoolsize": 101,
		"paytxfee": 0.00000000,
		"relayfee": 0.00000000,
		"errors": ""
	},
	"error": null,
	"id": "MultiChain-RPC-Client"
}
*/
type GetInfo struct {
	result struct {
		hash          string      `json:"hash"`
		miner         interface{} `json:"miner"`
		confirmations int         `json:"confirmations"`
		size          int         `json:"size"`
		height        int         `json:"height"`
		version       int         `json:"version"`
		merkleroot    string      `json:"merkleroot"`
		tx            []string    `json:"tx"`
		time          int         `json:"time"`
		nonce         int         `json:"nonce"`
		bits          string      `json:"bits"`
		difficulty    float64     `json:"difficulty"`
		chainwork     string      `json:"chainwork"`
		nextblockhash string      `json:"nextblockhash"`
	} `json:"result"`
	error interface{} `json:"error"`
	id    string      `json:"id"`
}

// ParseResponse takesa valid response and parses it into the model
func (m *GetInfo) ParseResponse(r Response) {
	err := mapstructure.Decode(m, &r)
	if err != nil {
		logger.LogThis.Error(err)
	}
}

// GetInfo Returns an object containing various state info.
//
//Result:
//```
//{
//  "version": xxxxx,           (numeric) the server version
//  "protocolversion": xxxxx,   (numeric) the protocol version
//  "walletversion": xxxxx,     (numeric) the wallet version
//  "balance": xxxxxxx,         (numeric) the total bitcoin balance of the wallet
//  "blocks": xxxxxx,           (numeric) the current number of blocks processed in the server
//  "timeoffset": xxxxx,        (numeric) the time offset
//  "connections": xxxxx,       (numeric) the number of connections
//  "proxy": "host:port",     	(string, optional) the proxy used by the server
//  "difficulty": xxxxxx,       (numeric) the current difficulty
//  "testnet": true|false,      (boolean) if the server is using testnet or not
//  "keypoololdest": xxxxxx,    (numeric) the timestamp (seconds since Unix epoch) of the oldest pre-generated key in the key pool
//  "keypoolsize": xxxx,        (numeric) how many new keys are pre-generated
//  "unlocked_until": ttt,      (numeric) the timestamp in seconds since epoch (midnight Jan 1 1970 GMT) that the wallet is unlocked for transfers, or 0 if the wallet is locked
//  "paytxfee": x.xxxx,         (numeric) the transaction fee set in BTC/kB
//  "relayfee": x.xxxx,         (numeric) minimum relay fee for non-free transactions in BTC/kB
//  "errors": "..."           (string) any error messages
//}
//```
func (client *Client) GetInfo() (Response, error) {

	msg := client.Command(
		"getinfo",
		[]interface{}{},
	)

	return client.Post(msg)
}
