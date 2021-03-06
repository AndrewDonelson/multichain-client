package multichain

import (
	"strings"

	"github.com/mitchellh/mapstructure"
)

// Permissions represents an actual permission
type Permissions struct {
	Address    string        `json:"address"`
	For        interface{}   `json:"for"`
	Type       string        `json:"type"`
	Startblock int           `json:"startblock"`
	Endblock   int64         `json:"endblock"`
	Admins     []string      `json:"admins"`
	Pending    []interface{} `json:"pending"`
}

// GetPermissionInfo returns permissions either all, address specific or type specific
type GetPermissionInfo struct {
	Result []Permissions `json:"result"`
	Error  interface{}   `json:"error"`
	ID     string        `json:"id"`
}

// ParseResponse takes a valid response and parses it into the model
func (m *GetPermissionInfo) ParseResponse(r Response) {
	err := mapstructure.Decode(r, &m)
	//fmt.Println(m)
	if err != nil {
		panic(err)
	}
}

// ListPermissions returns a list of all permissions which have been explicitly granted to addresses.
// To list information about specific global permissions, set permissions to one of connect, send, receive, issue, mine, activate, admin, or a list thereof.
// Omit or pass * or all to list all global permissions.
// For per-asset or per-stream permissions, use the form entity.issue, entity.write,admin or entity.* where entity is an asset or stream name, ref or creation txid.
// Provide a list in addresses to list the permissions for particular addresses or omit for all addresses.
// If verbose is true, the admins output field lists the administrator/s who assigned the corresponding permission,
// and the pending field lists permission changes which are waiting to reach consensus.
func (client *Client) ListPermissions(permissions []string, addresses []string, verbose bool) (Response, error) {

	// Omit or pass * or all to list all global permissions.
	p := "*"
	if len(permissions) > 0 {
		p = strings.Join(permissions[:], ",")
	}

	// Same for addresses
	a := "*"
	if len(addresses) > 0 {
		a = strings.Join(addresses[:], ",")
	}

	params := []interface{}{
		p,
		a,
		verbose,
	}

	msg := client.Command(
		"listpermissions",
		params,
	)

	return client.Post(msg)
}
