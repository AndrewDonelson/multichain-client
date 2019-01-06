package multichain

import (
	"github.com/mitchellh/mapstructure"
)

type Peer struct {
	ID             int           `json:"id"`
	Addr           string        `json:"addr"`
	Addrlocal      string        `json:"addrlocal"`
	Services       string        `json:"services"`
	Lastsend       int           `json:"lastsend"`
	Lastrecv       int           `json:"lastrecv"`
	Bytessent      int           `json:"bytessent"`
	Bytesrecv      int           `json:"bytesrecv"`
	Conntime       int           `json:"conntime"`
	Pingtime       float64       `json:"pingtime"`
	Version        int           `json:"version"`
	Subver         string        `json:"subver"`
	Inbound        bool          `json:"inbound"`
	Startingheight int           `json:"startingheight"`
	Banscore       int           `json:"banscore"`
	SyncedHeaders  int           `json:"synced_headers"`
	SyncedBlocks   int           `json:"synced_blocks"`
	Inflight       []interface{} `json:"inflight"`
	Whitelisted    bool          `json:"whitelisted"`
}

// GetPeerInfo is a struct representing the result from the multichain.GetInfo() RPC Command
type GetPeerInfo struct {
	Result []Peer      `json:"result"`
	error  interface{} `json:"error"`
	id     string      `json:"id"`
}

// ParseResponse takes a valid response and parses it into the model
func (m *GetPeerInfo) ParseResponse(r Response) {
	err := mapstructure.Decode(r, &m)
	//fmt.Println(m)
	if err != nil {
		panic(err)
	}
}

// GetPeerInfo returns information about the other nodes to which this node is
// connected. If this is a MultiChain blockchain, includes handshake and
// handshakelocal fields showing the remote and local address used during the
// handshaking for that connection.
func (client *Client) GetPeerInfo() (Response, error) {

	msg := client.Command(
		"getpeerinfo",
		[]interface{}{},
	)

	return client.Post(msg)
}
