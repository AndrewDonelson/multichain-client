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

// GetBlock returns information about the block with hash (retrievable
// from getblockhash) or at the given height in the active chain. Set
// verbose to 0 or false for the block in raw hexadecimal form. Set to
// 1 or true for a block summary including the miner address and a list
// of txids. Set to 2 to 3 to include more information about each transaction
// and its raw hexadecimal. Set to 4 to include a full description of each
// transaction, formatted like the output of decoderawtransaction.
//
// Parameters:
// hash|height
// (verbose=1)
func (client *Client) GetBlock(heightOrHash string) (Response, error) {

	msg := client.Command(
		"getblock",
		[]interface{}{
			heightOrHash,
		},
	)

	return client.Post(msg)
}
