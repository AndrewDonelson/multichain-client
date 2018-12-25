package multichain

import "github.com/mitchellh/mapstructure"

type GetBlock struct {
	Result struct {
		Hash          string      `json:"hash"`
		Miner         interface{} `json:"miner"`
		Confirmations int         `json:"confirmations"`
		Size          int         `json:"size"`
		Height        int         `json:"height"`
		Version       int         `json:"version"`
		Merkleroot    string      `json:"merkleroot"`
		Tx            []string    `json:"tx"`
		Time          int         `json:"time"`
		Nonce         int         `json:"nonce"`
		Bits          string      `json:"bits"`
		Difficulty    float64     `json:"difficulty"`
		Chainwork     string      `json:"chainwork"`
		Nextblockhash string      `json:"nextblockhash"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

// ParseResponse takesa valid response and parses it into the model
func (m *GetBlock) ParseResponse(r Response) {
	err := mapstructure.Decode(r, &m)
	if err != nil {
		panic(err)
	}
}

func (client *Client) GetBlock(heightOrHash string) (Response, error) {

	msg := client.Command(
		"getblock",
		[]interface{}{
			heightOrHash,
		},
	)

	return client.Post(msg)
}
