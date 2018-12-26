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

// GetAddressBalances returns a list of all the asset balances for address
// in this node’s wallet, with at least minconf confirmations. Use includeLocked
// to include unspent outputs which have been locked, e.g. by a call to
// preparelockunspent.
//
// Paramters:
//	address (minconf=1)
//	(includeLocked=false)
func (client *Client) GetAddressBalances(address string) (Response, error) {

	msg := client.Command(
		"getaddressbalances",
		[]interface{}{
			address,
		},
	)

	return client.Post(msg)
}
