package blockchain

import (
	"time"

	"github.com/AndrewDonelson/multichain-client"

	"github.com/NlaakStudiosLLC/GoWAF/framework/logger"
)

// BlockchainClient provides global access to the Applications MultiChain
var BlockchainClient Client

// Client
type Client struct {
	apptag         string
	beats          int
	updateInterval *time.Ticker
	Client         *multichain.Client
	Info           *Info
	Verbose        bool
}

//Struct to JSON
type Payload struct {
	ProtocolID uint
	Asset      interface{}
}

func (p *Payload) ToString() string {
	return ""
}

// NewClient creates a new instance of Client and saves it to BlockchainClient
func NewClient(name string, host string, port int, user string, password string, verbose bool) {
	m := &Client{}

	m.Verbose = verbose

	//Update Every 5 minutes (TODO: Move this to config)
	m.updateInterval = time.NewTicker(300 * time.Second)

	logger.LogThis.Info("Connecting to ", name, "...")

	m.Client = multichain.NewClient(
		name,
		user,
		password,
		8071,
	).ViaNode(
		host,
		port,
	)

	if m.Client.Connected {
		logger.LogThis.Info("Connected to via JSON-RPC node ", host, ":", port)
		m.Info = GetNewInfo(m, m.Verbose)
		m.apptag = name
	} else {
		logger.LogThis.Error("Connection to blockchain failed - Blockchain JSON-RPC not available")
		m = nil
	}

	BlockchainClient = *m
}

// Returns the number of Clients since app started
func (m *Client) UpdateNodeInfo() {
	m.beats++
	logger.LogThis.Info("Blockchain ", m.apptag, " Update #", m.beats)
	m.Info = GetNewInfo(m, m.Verbose)
}

// Run the process.
func (m *Client) Run() {
	logger.LogThis.Info("Blockchain ", m.apptag, " Sync started.")
	// Dispatch a process into the background.
	go func() {
		// Now run it forever.
		for {
			// Watch for stuff to happen.
			select {
			// When the Client ticker is fired, execute this.
			case <-m.updateInterval.C:
				m.UpdateNodeInfo()
			} // select
		} // for
	}() // go func

} // func
