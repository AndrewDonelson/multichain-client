package main

import "github.com/AndrewDonelson/multichain-client/blockchain"

func main() {
	blockchain.NewClient(
		"GlobalWealthAndFreedom",
		"73.55.167.87",
		5001,
		"gwfrpc",
		"rpcuserpassword",
		true,
	)
	blockchain.BlockchainClient.Run()
}
