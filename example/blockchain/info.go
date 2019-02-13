package blockchain

import (
	"github.com/AndrewDonelson/multichain-client"
	"github.com/NlaakStudiosLLC/GoWAF/framework/logger"
)

// Info holds various information about the connected blockchain node
type Info struct {
	//Network
	Params  *multichain.GetBlockchainParams
	Chain   *multichain.GetInfo
	Wallet  *multichain.GetWalletInfo
	MemPool *multichain.GetMemPoolInfo
	Peers   *multichain.GetPeerInfo
	verbose bool
	failed  bool
}

// GetNewInfo is a help function that querys the node via JSON-RPC to obtain
// and return information about the connected blockchain node. This includes
// Blockchain Parameters, Generl Blockchain Data, Connected Peers, Wallet and
// Memory Pool. Currently this is synced every 5 minutes (300 seconds)
//
// You can access this information as well as make RPC Client API calls by
// using the global object AppMultiChain.[Info||Client]
func GetNewInfo(m *Client, verbose bool) *Info {
	i := &Info{}
	i.verbose = verbose

	if i.failed != true {
		i.GetChain(m)
	}

	if i.failed != true {
		i.GetParams(m)
	}

	if i.failed != true {
		i.GetPeers(m)
	}

	if i.failed != true {
		i.GetWallet(m)
	}

	if i.failed != true {
		i.GetMemPool(m)
	}

	if i.failed == true {
		logger.LogThis.Info("Unable to get data from Blockchain RPC Server")
	}

	return i
}

// GetParams returns populates the Blockchain Paramters via multichain API call
func (i *Info) GetParams(m *Client) {
	var info multichain.GetBlockchainParams

	obj, err1 := m.Client.GetBlockchainParams()
	if err1 != nil {
		i.failed = true
		if i.verbose {
			logger.LogThis.Error("RPC Request failed: ", err1)
		}
	} else {
		info.ParseResponse(obj)
		if info.Result.ChainName != "" {
			i.Params = &info
		}
		if i.verbose {
			logger.LogThis.Info("getblockchainparams: [", i.Params, "]")
		}
	}
}

// GetChain returns populates the Blockchain Information via multichain API call
func (i *Info) GetChain(m *Client) {
	var info multichain.GetInfo

	obj, err1 := m.Client.GetInfo()

	if err1 != nil {
		i.failed = true
		if i.verbose {
			logger.LogThis.Error("RPC Request failed: ", err1)
		}
	} else {
		info.ParseResponse(obj)
		if info.Result.Version != "" {
			i.Chain = &info
		}
		if i.verbose {
			logger.LogThis.Info("getinfo: [", i.Chain, "]")
		}
	}
}

// GetWallet returns populates the Wallet via multichain API call
func (i *Info) GetWallet(m *Client) {
	var info multichain.GetWalletInfo
	obj, err1 := m.Client.GetWalletInfo()
	if err1 != nil {
		i.failed = true
		if i.verbose {
			logger.LogThis.Error("RPC Request failed: ", err1)
		}
	} else {
		info.ParseResponse(obj)
		if info.Result.Walletversion != 0 {
			i.Wallet = &info
		}
		if i.verbose {
			logger.LogThis.Info("getwalletinfo: [", i.Wallet, "]")
		}
	}
}

// GetMemPool returns populates the Blockchain Memory Pool via multichain API call
func (i *Info) GetMemPool(m *Client) {
	var info multichain.GetMemPoolInfo
	obj, err1 := m.Client.GetMemPoolInfo()
	if err1 != nil {
		i.failed = true
		if i.verbose {
			logger.LogThis.Error("RPC Request failed: ", err1)
		}
	} else {
		info.ParseResponse(obj)
		i.MemPool = &info
		if i.verbose {
			logger.LogThis.Info("getmempoolinfo: [", i.MemPool, "]")
		}
	}
}

// GetPeers returns populates the Blockchain connect peers via multichain API call
func (i *Info) GetPeers(m *Client) {
	var info multichain.GetPeerInfo
	obj, err1 := m.Client.GetPeerInfo()
	if err1 != nil {
		i.failed = true
		if i.verbose {
			logger.LogThis.Error("RPC Request failed: ", err1)
		}
	} else {
		//fmt.Println("obj = ", obj)
		info.ParseResponse(obj)
		//fmt.Println("info = ", info)

		i.Peers = &info
		if i.verbose {
			logger.LogThis.Info("getpeerinfo: [", i.Peers, "]")
		}
	}
}
