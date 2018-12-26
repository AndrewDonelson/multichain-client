package multichain

import "github.com/mitchellh/mapstructure"

// GetTransaction is a struct representing the result from the multichain.GetTransaction() RPC Command
type GetTransaction struct {
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
func (m *GetTransaction) ParseResponse(r Response) {
	err := mapstructure.Decode(r, &m)
	if err != nil {
		panic(err)
	}
}

// GetTransaction If verbose is 1, returns a JSON object describing transaction
// txid. For a MultiChain blockchain, each transaction output includes assets
// and permissions fields listing any assets or permission changes encoded within
// that output. There will also be a data field listing the content of any
// OP_RETURN outputs in the transaction.
//
// Parameters:
// 	txid
// 	(verbose=0)
func (client *Client) GetTransaction(txid string) (Response, error) {

	msg := client.Command(
		"getrawtransaction",
		[]interface{}{
			txid,
			1,
		},
	)

	return client.Post(msg)
}
