package multichain

import (
	"math/rand"
	"testing"
)

const (
	ChainName   = "gwfchain"
	RPCUser     = "rpcadmin"
	RPCPassword = "2T22j'8@5Z!}K+KGt)']PQf_"
	RPCHost     = "73.55.167.87"
	RPCPort     = 5001

	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	client        *Client
	BlockHash     string
	TXID          string
	WalletAddress string
	StreamName    string
)

func Init(debug bool) {

	if debug == true {
		// Change these properties for your chain
		client = NewDebugClient(
			ChainName,
			RPCUser,
			RPCPassword,
			8071,
		).ViaNode(
			RPCHost,
			RPCPort,
		)
	} else {
		// Change these properties for your chain
		client = NewClient(
			ChainName,
			RPCUser,
			RPCPassword,
			8071,
		).ViaNode(
			RPCHost,
			RPCPort,
		)
	}
}

/***************** DONT EDIT BELOW THIS LINE *********************/

//RandString - generate random string using masking with source
func RandString(n int) string {
	b := make([]byte, n)
	l := len(letterBytes)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < l {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func testGetInfo(t *testing.T) {
	t.Helper()

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
	t.Helper()

	fName := "GetBlockchainInfo"
	obj, err := client.GetBlockchainInfo()
	if err != nil {
		t.Error(fName, err)
	} else {
		var info GetBlockchainInfo
		info.ParseResponse(obj)
		BlockHash = info.Result.Bestblockhash
		t.Log(fName, ": Passed!")
	}
}

func testGetBlockchainParams(t *testing.T) {
	t.Helper()

	fName := "GetBlockchainParams"
	obj, err := client.GetBlockchainParams()
	if err != nil {
		t.Error(fName, err)
	} else {
		var info GetBlockchainParams
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

func testGetRawMemPool(t *testing.T) {

	fName := "GetRawMemPool"
	obj, err := client.GetRawMemPool()
	if err != nil {
		t.Error(fName, err)
	} else {
		var info GetRawMemPool
		info.ParseResponse(obj)
		t.Log(fName, ": Passed!")
	}
}

func testGetWalletInfo(t *testing.T) {

	fName := "GetWalletInfo"
	obj, err := client.GetWalletInfo()
	if err != nil {
		t.Error(fName, err)
	} else {
		var info GetWalletInfo
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
		TXID = info.Result.Tx[0]
		t.Log(fName, ": Passed!")
	}
}

func testGetTransaction(t *testing.T) {

	fName := "GetTransaction"
	obj, err := client.GetTransaction(TXID)
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
		WalletAddress = info.Result[0].Address
		t.Log(fName, ": Passed!")
	}
}

func testGetNewAddress(t *testing.T) {

	fName := "GetNewAddress"
	obj, err := client.GetNewAddress()
	if err != nil {
		t.Error(fName, err)
	} else {
		var info GetNewAddress
		info.ParseResponse(obj)
		t.Log(fName, ": Passed!")
	}
}

func testGetAddressBalances(t *testing.T) {

	fName := "GetAddressBalances"
	obj, err := client.GetAddressBalances(WalletAddress)
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

func testCreateKeyPair(t *testing.T) {

	fName := "CreateKeyPairs"
	obj, err := client.CreateKeypair()
	if err != nil {
		t.Error(fName, err)
	} else {
		t.Log(obj)
		t.Log(fName, ": Passed!")
	}
}

func testListPermissions(t *testing.T) {

	fName := "ListPermissions"
	//obj, err := client.ListPermissions([]string{"receive", "send"}, []string{}, false)
	obj, err := client.ListPermissions(nil, nil, true)
	if err != nil {
		t.Error(fName, " (Global)", err)
	} else {
		var info GetPermissionInfo
		info.ParseResponse(obj)
		t.Log(fName, ": Passed!")
	}

	obj2, err2 := client.ListPermissions([]string{"receive", "send"}, nil, true)
	if err2 != nil {
		t.Error(fName, " (Type)", err)
	} else {
		var info GetPermissionInfo
		info.ParseResponse(obj2)
		t.Log(fName, ": Passed!")
	}

	obj3, err3 := client.ListPermissions(nil, []string{WalletAddress}, true)
	if err3 != nil {
		t.Error(fName, " (Type)", err)
	} else {
		var info GetPermissionInfo
		info.ParseResponse(obj3)
		t.Log(fName, ": Passed!")
	}

}

func testCreateStream(t *testing.T) {
	fName := "CreateStream"
	StreamName := RandString(32)
	obj, err := client.CreateStream(StreamName, false)
	if err != nil {
		t.Error(fName, err)
	} else {
		var info ListAddresses
		info.ParseResponse(obj)
		t.Log(fName, ": Passed!")
	}
}
func TestAll(t *testing.T) {
	Init(true)                 //set to true for debug client
	testGetInfo(t)             // Passed
	testGetBlockchainInfo(t)   // Passed
	testGetBlockchainParams(t) // Passed
	testGetPeerInfo(t)         // Passed
	testGetMemPoolInfo(t)      // Passed
	testGetRawMemPool(t)       // Passed
	testGetWalletInfo(t)       // Passed
	testGetNewAddress(t)       // Passed
	testGetAddresses(t)        // Passed
	testGetBlock(t)            // Passed
	testGetTransaction(t)      // Passed
	testGetAddressBalances(t)  // Passed
	testListAddresses(t)       // Passed
	testCreateKeyPair(t)       // Passed
	testListPermissions(t)     // Passed

	testCreateStream(t) // Passed

}
