package multichain

// CreateRawTransaction Creates a transaction spending the specified inputs,
// sending to the given addresses. Each amounts field can specify a quantity of
// the native blockchain currency, some asset(s) and/or inline metadata, using one
// of these forms. The optional data array adds one or more metadata outputs to
// the transaction, where each element is formatted as passed to appendrawdata.
// The optional action parameter can be lock (locks the given inputs in the wallet),
// sign (signs the transaction using wallet keys), lock,sign (does both) or send
// (signs and sends the transaction). If action is send the txid is returned. If
// action contains sign, an object with hex and complete fields is returned, as
// for signrawtransaction. Otherwise, the raw transaction hexadecimal is returned.
// See raw transactions for more details on building raw transactions.
//
// Parameters:
//	[{"txid":"id","vout":n},...]
//	{"address":amounts,...}
//	(data=[]) (action="")
func (client *Client) CreateRawTransaction(destinationAddress string, assets map[string]float64, unspentOutputs ...*Unspent) (Response, error) {

	msg := client.Command(
		"createrawtransaction",
		[]interface{}{
			unspentOutputs,
			map[string]interface{}{
				destinationAddress: assets,
			},
		},
	)

	return client.Post(msg)
}
