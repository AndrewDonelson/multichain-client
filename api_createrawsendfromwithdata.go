package multichain

// CreateRawSendFromWithData This works like createrawtransaction, except it
// automatically selects the transaction inputs from those belonging to
// from-address, to cover the appropriate amounts. One or more change outputs
// going back to from-address will also be added to the end of the transaction.
//
// Parameters:
//	from-address
//	{"to-address":amounts,...}
//	(data=[]) (action="")
func (client *Client) CreateRawSendFromWithData(watchAddress, destinationAddress string, assets map[string]float64, data []string) (Response, error) {

	msg := client.Command(
		"createrawsendfrom",
		[]interface{}{
			watchAddress,
			map[string]interface{}{
				destinationAddress: assets,
			},
			data,
		},
	)

	return client.Post(msg)
}
