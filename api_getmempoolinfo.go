package multichain

import (
	"github.com/mitchellh/mapstructure"
)

// GetMemPoolInfo is a struct representing the result from the multichain.GetMemPoolInfo() RPC Command
type GetMemPoolInfo struct {
	Result struct {
		Size  int `json:"size"`
		Bytes int `json:"bytes"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

// ParseResponse takesa valid response and parses it into the model
func (m *GetMemPoolInfo) ParseResponse(r Response) {
	err := mapstructure.Decode(r, &m)
	if err != nil {
		panic(err)
	}
}

// GetMemPoolInfo execute the multichain RPC Command `getinfo` and returns the response
func (client *Client) GetMemPoolInfo() (Response, error) {

	msg := client.Command(
		"getmempoolinfo",
		[]interface{}{},
	)
	return client.Post(msg)
}
