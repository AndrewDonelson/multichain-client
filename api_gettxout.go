package multichain

//GetTxOut Returns details about an unspent transaction output vout of txid.
// For a MultiChain blockchain, includes assets and permissions fields listing
// any assets or permission changes encoded within the output. Set unconfirmed
// to true to include unconfirmed transaction outputs.
//
// Parameters:
// txid vout
// (unconfirmed=false)
func (client *Client) GetTxOut(txid string, vout int) (Response, error) {

	msg := client.Command(
		"gettxout",
		[]interface{}{
			txid,
			vout,
		},
	)

	return client.Post(msg)
}
