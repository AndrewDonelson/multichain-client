package multichain

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// GetInfo is a struct representing the result from the multichain.GetInfo() RPC Command
type ListAddresses struct {
	Result []struct {
		Address string `json:"address"`
		Ismine  bool   `json:"ismine"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

// ParseResponse takesa valid response and parses it into the model
func (m *ListAddresses) ParseResponse(r Response) {
	err := mapstructure.Decode(r, &m)
	if err != nil {
		panic(err)
	}
}

func (client *Client) ListAddresses(verbose bool, addresses ...string) (Response, error) {

	v := fmt.Sprintf("verbose=%v", verbose)

	var params []interface{}

	if len(addresses) > 0 {
		params = []interface{}{
			addresses,
			v,
		}
	} else {
		if verbose {
			params = []interface{}{
				[]string{},
				v,
			}
		}
	}

	msg := client.Command(
		"listaddresses",
		params,
	)

	return client.Post(msg)
}
