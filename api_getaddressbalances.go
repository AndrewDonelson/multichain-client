package multichain

import "github.com/mitchellh/mapstructure"

// GetInfo is a struct representing the result from the multichain.GetInfo() RPC Command
type GetAddressBalances struct {
	Result []struct {
		Assetref string  `json:"assetref"`
		Qty      float64 `json:"qty"`
		Raw      int64   `json:"raw"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

// ParseResponse takesa valid response and parses it into the model
func (m *GetAddressBalances) ParseResponse(r Response) {
	err := mapstructure.Decode(r, &m)
	if err != nil {
		panic(err)
	}
}

func (client *Client) GetAddressBalances(address string) (Response, error) {

	msg := client.Command(
		"getaddressbalances",
		[]interface{}{
			address,
		},
	)

	return client.Post(msg)
}
