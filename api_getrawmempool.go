package multichain

import (
	"github.com/mitchellh/mapstructure"
)

// GetRawMemPool is a struct representing the result from the multichain.GetRawMemPool() RPC Command
type GetRawMemPool struct {
	Result struct {
		//TODO: Need result to define structure
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

// ParseResponse takes a valid response and parses it into the model
func (m *GetRawMemPool) ParseResponse(r Response) {
	err := mapstructure.Decode(r, &m)
	if err != nil {
		//panic(err)
		// No Panic, it can return no results if no transactions are pending
	}
}

// GetRawMemPool Returns a list of transaction IDs which are in the node’s
// memory pool (see getmempoolinfo).
func (client *Client) GetRawMemPool() (Response, error) {

	msg := client.Command(
		"getrawmempool",
		[]interface{}{},
	)
	return client.Post(msg)
}
