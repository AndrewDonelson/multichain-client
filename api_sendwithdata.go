package multichain

func (client *Client) SendWithData(accountAddress string, assets map[string]float64, data string) (Response, error) {

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
