package multichain

import (
	"github.com/mitchellh/mapstructure"
)

// GetInfo is a struct representing the result from the multichain.GetInfo() RPC Command
type GetInfo struct {
	Result struct {
		Version         string  `json:"version"`
		Nodeversion     int     `json:"nodeversion"`
		Protocolversion int     `json:"protocolversion"`
		Chainname       string  `json:"chainname"`
		Description     string  `json:"description"`
		Protocol        string  `json:"protocol"`
		Port            int     `json:"port"`
		Setupblocks     int     `json:"setupblocks"`
		Nodeaddress     string  `json:"nodeaddress"`
		Burnaddress     string  `json:"burnaddress"`
		Incomingpaused  bool    `json:"incomingpaused"`
		Miningpaused    bool    `json:"miningpaused"`
		Walletversion   int     `json:"walletversion"`
		Balance         float64 `json:"balance"`
		Walletdbversion int     `json:"walletdbversion"`
		Reindex         bool    `json:"reindex"`
		Blocks          int     `json:"blocks"`
		Timeoffset      int     `json:"timeoffset"`
		Connections     int     `json:"connections"`
		Proxy           string  `json:"proxy"`
		Difficulty      float64 `json:"difficulty"`
		Testnet         bool    `json:"testnet"`
		Keypoololdest   int     `json:"keypoololdest"`
		Keypoolsize     int     `json:"keypoolsize"`
		Paytxfee        float64 `json:"paytxfee"`
		Relayfee        float64 `json:"relayfee"`
		Errors          string  `json:"errors"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

// ParseResponse takesa valid response and parses it into the model
func (m *GetInfo) ParseResponse(r Response) {
	err := mapstructure.Decode(r, &m)
	if err != nil {
		panic(err)
	}
}

// GetInfo execute the multichain RPC Command `getinfo` and returns the response
func (client *Client) GetInfo() (Response, error) {

	msg := client.Command(
		"getinfo",
		[]interface{}{},
	)
	return client.Post(msg)
}
