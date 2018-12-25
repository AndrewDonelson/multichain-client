package multichain

import (
	"github.com/mitchellh/mapstructure"
)

// GetBlockchainInfo is a struct representing the result from the multichain.GetBlockchainInfo() RPC Command
type GetBlockchainInfo struct {
	Result struct {
		Chain                string  `json:"chain"`
		Chainname            string  `json:"chainname"`
		Description          string  `json:"description"`
		Protocol             string  `json:"protocol"`
		Setupblocks          int     `json:"setupblocks"`
		Reindex              bool    `json:"reindex"`
		Blocks               int     `json:"blocks"`
		Headers              int     `json:"headers"`
		Bestblockhash        string  `json:"bestblockhash"`
		Difficulty           float64 `json:"difficulty"`
		Verificationprogress float64 `json:"verificationprogress"`
		Chainwork            string  `json:"chainwork"`
		Chainrewards         float64 `json:"chainrewards"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

// ParseResponse takesa valid response and parses it into the model
func (m *GetBlockchainInfo) ParseResponse(r Response) {
	err := mapstructure.Decode(r, &m)
	if err != nil {
		panic(err)
	}
}

// GetBlockchainInfo execute the multichain RPC Command `getinfo` and returns the response
func (client *Client) GetBlockchainInfo() (Response, error) {

	msg := client.Command(
		"getblockchaininfo",
		[]interface{}{},
	)
	return client.Post(msg)
}
