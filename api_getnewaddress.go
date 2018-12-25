package multichain

import "github.com/mitchellh/mapstructure"

// GetNewAddress is a struct representing the result from the multichain.GetNewAddress() RPC Command
type GetNewAddress struct {
	Result struct {
		Version string `json:"version"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

// ParseResponse takesa valid response and parses it into the model
func (m *GetNewAddress) ParseResponse(r Response) {
	err := mapstructure.Decode(r, &m)
	if err != nil {
		panic(err)
	}
}

func (client *Client) GetNewAddress() (Response, error) {

	msg := client.Command(
		"getnewaddress",
		[]interface{}{},
	)

	return client.Post(msg)
}
