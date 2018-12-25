package multichain

import (
	"testing"
)

const (
	// Change this to a known BlockHash on your chain
	BlockHash = "0063243875ac046f964a5b27889d926baf9f50e097b457a5bf5e4f66e175d081"
	// Change this to a known Transacation ID on your chain
	TransactionHash = "4ca7e5e480f9e66fb6b8ceadca0c177e48bee8a7e8df4c685e5dcc03d82eb15e"
	NodeAddress     = "1B2snMEHWsAMk9p6AZEXEcWN5cDvmWNc85Hdhw"
)

var client *Client

func Init() {

	// Change these properties for your chain
	client = NewDebugClient(
		"nlaakstudioscryptobond",
		"nscb",
		"2T22j'8@5Z!}K+KGt)']PQf_",
		8071,
	).ViaNode(
		"73.55.167.87",
		5001,
	)
}

/***************** DONT EDIT BELOW THIS LINE *********************/

func testGetInfo(t *testing.T) {

	fName := "GetInfo"
	obj, err := client.GetInfo()
	if err != nil {
		t.Error(fName, err)
	} else {
		var info GetInfo
		info.ParseResponse(obj)
		t.Log(fName, ": Passed!")
	}
}

func testGetBlockchainInfo(t *testing.T) {

	fName := "GetBlockchainInfo"
	obj, err := client.GetBlockchainInfo()
	if err != nil {
		t.Error(fName, err)
	} else {
		var info GetBlockchainInfo
		info.ParseResponse(obj)
		t.Log(fName, ": Passed!")
	}
}

func testGetPeerInfo(t *testing.T) {

	fName := "GetPeerInfo"
	obj, err := client.GetPeerInfo()
	if err != nil {
		t.Error(fName, err)
	} else {
		var info GetPeerInfo
		info.ParseResponse(obj)
		t.Log(fName, ": Passed!")
	}
}

func testGetMemPoolInfo(t *testing.T) {

	fName := "GetMemPoolInfo"
	obj, err := client.GetMemPoolInfo()
	if err != nil {
		t.Error(fName, err)
	} else {
		var info GetMemPoolInfo
		info.ParseResponse(obj)
		t.Log(fName, ": Passed!")
	}
}

func testGetBlock(t *testing.T) {

	fName := "GetBlock"
	obj, err := client.GetBlock(BlockHash)
	if err != nil {
		t.Error(fName, err)
	} else {
		var info GetBlock
		info.ParseResponse(obj)
		t.Log(fName, ": Passed!")
	}
}

func testGetTransaction(t *testing.T) {

	fName := "GetTransaction"
	obj, err := client.GetTransaction(TransactionHash)
	if err != nil {
		t.Error(fName, err)
	} else {
		var info GetTransaction
		info.ParseResponse(obj)
		t.Log(fName, ": Passed!")
	}
}

func testGetAddresses(t *testing.T) {

	fName := "GetAddresses"
	obj, err := client.GetAddresses(true)
	if err != nil {
		t.Error(fName, err)
	} else {
		var info GetAddresses
		info.ParseResponse(obj)
		t.Log(fName, ": Passed!")
	}
}

func testGetAddressBalances(t *testing.T) {

	fName := "GetAddressBalances"
	obj, err := client.GetAddressBalances(NodeAddress)
	if err != nil {
		t.Error(fName, err)
	} else {
		var info GetAddressBalances
		info.ParseResponse(obj)
		t.Log(fName, ": Passed!")
	}
}

func testListAddresses(t *testing.T) {
	fName := "ListAddresses"
	obj, err := client.ListAddresses(false)
	if err != nil {
		t.Error(fName, err)
	} else {
		var info ListAddresses
		info.ParseResponse(obj)
		t.Log(fName, ": Passed!")
	}
}
func TestAll(t *testing.T) {
	Init()
	testGetInfo(t)            // Passed
	testGetBlockchainInfo(t)  // Passed
	testGetPeerInfo(t)        // Passed
	testGetMemPoolInfo(t)     // Passed
	testGetBlock(t)           // Passed
	testGetTransaction(t)     // Passed
	testGetAddresses(t)       // Passed
	testGetAddressBalances(t) // Passed
	testListAddresses(t)      // Passed
}
