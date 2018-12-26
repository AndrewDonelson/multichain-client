package multichain

import "github.com/mitchellh/mapstructure"

// GetInfo is a struct representing the result from the multichain.GetInfo() RPC Command
type GetAddresses struct {
	Result []struct {
		Address      string `json:"address"`
		Ismine       bool   `json:"ismine"`
		Iswatchonly  bool   `json:"iswatchonly"`
		Isscript     bool   `json:"isscript"`
		Pubkey       string `json:"pubkey"`
		Iscompressed bool   `json:"iscompressed"`
		Account      string `json:"account"`
		Synchronized bool   `json:"synchronized"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

// ParseResponse takesa valid response and parses it into the model
func (m *GetAddresses) ParseResponse(r Response) {
	err := mapstructure.Decode(r, &m)
	if err != nil {
		panic(err)
	}
}

func (client *Client) GetAddresses(verbose bool) (Response, error) {

	msg := client.Command(
		"getaddresses",
		[]interface{}{
			verbose,
		},
	)

	return client.Post(msg)
}
