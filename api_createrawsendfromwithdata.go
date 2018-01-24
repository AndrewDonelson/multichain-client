package multichain

// This works like createrawsendfrom but it adds the possiblity to send data
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
