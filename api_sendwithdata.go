package multichain

func (client *Client) SendWithData(accountAddress, assets map[string]float64, data string) (Response, error) {

	msg := client.Command(
		"sendassettoaddress",
		[]interface{}{
			accountAddress,
			assets,
			data,
		},
	)

	return client.Post(msg)
}
