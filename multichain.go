package multichain

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	//
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	//
	"github.com/dghubble/sling"
)

const (
	//ClientID is the ID of this Multichain RPC Client
	ClientID = "MultiChain-RPC-Client"
)

// Response defines the Response type to used with gettign the reponse from the Multichain RPC server
type Response map[string]interface{}

//Result returns the the result property of a Multichain RPC Response
func (r Response) Result() interface{} {
	return r["result"]
}

// Client is the model for the Multichain Client
type Client struct {
	httpClient  *http.Client
	chain       string
	host        string
	port        int
	credentials string
	debug       bool
	Connected   bool
}

// NewClient returns a non debug (verbose) Multichain RPC Client object for a given chain
func NewClient(chain, username, password string, port int) *Client {

	credentials := username + ":" + password

	return &Client{
		httpClient:  &http.Client{},
		chain:       chain,
		port:        port,
		credentials: base64.StdEncoding.EncodeToString([]byte(credentials)),
		debug:       false,
		Connected:   true,
	}
}

// NewClient returns a debug (verbose) Multichain RPC Client object for a given chain
func NewDebugClient(chain, username, password string, port int) *Client {
	credentials := username + ":" + password

	return &Client{
		httpClient:  &http.Client{},
		chain:       chain,
		port:        port,
		credentials: base64.StdEncoding.EncodeToString([]byte(credentials)),
		debug:       true,
		Connected:   true,
	}

	return NewClient(chain, username, password, port)
}

// ViaNode sets the desired node IP and port for a Multichain client
func (client *Client) ViaNode(ipv4 string, port int) *Client {
	c := *client
	c.host = fmt.Sprintf(
		"http://%s:%v",
		ipv4,
		port,
	)
	return &c
}

// IsDebugMode returns the clients debug mode state
func (client *Client) IsDebugMode() bool {
	return client.debug
}

// DebugMode sets the clients debug mode state
func (client *Client) DebugMode() *Client {
	client.debug = true
	return client
}

//Urlfetch Fetches the given HTTP URL, blocking until the result is returned.
func (client *Client) Urlfetch(ctx context.Context, seconds ...int) {

	if len(seconds) > 0 {
		ctx, _ = context.WithDeadline(
			ctx,
			time.Now().Add(time.Duration(1000000000*seconds[0])*time.Second),
		)
	}

	client.httpClient = urlfetch.Client(ctx)
}

//msg returns a prepared JSON-RPC request for communicating with the Multichain RPC Server
func (client *Client) msg(params []interface{}) map[string]interface{} {
	return map[string]interface{}{
		"jsonrpc": "1.0",
		"id":      ClientID,
		"params":  params,
	}
}

//Command sets the RPC Command to invoke in the Multichain JSON-RPC Request
func (client *Client) Command(method string, params []interface{}) map[string]interface{} {

	msg := client.msg(params)
	msg["method"] = fmt.Sprintf("%s", method)

	if client.debug {
		fmt.Println(msg)
	}

	return msg
}

//Post given a properly prepared JSON-RPC request (msg) will post to the Multichain RPC server and return a Response and Error if applicable.
func (client *Client) Post(msg interface{}) (Response, error) {

	if client.debug {
		fmt.Println("DEBUG MODE ON...")
		fmt.Println(client)
		b, _ := json.Marshal(msg)
		fmt.Println(string(b))
	}

	request, err := sling.New().Post(client.host).BodyJSON(msg).Request()
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Basic "+client.credentials)

	resp, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	} else {
		if resp.StatusCode != 200 {
			return nil, errors.New(resp.Status)
		}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if client.debug {
		fmt.Println(string(b))
	}

	obj := make(Response)

	err = json.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}

	if obj["error"] != nil {
		e := obj["error"].(map[string]interface{})
		var s string
		m, ok := msg.(map[string]interface{})
		if ok {
			s = fmt.Sprintf("multichaind - '%s': %s", m["method"], e["message"].(string))
		} else {
			s = fmt.Sprintf("multichaind - %s", e["message"].(string))
		}
		return nil, errors.New(s)
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("INVALID RESPONSE STATUS CODE: " + strconv.Itoa(resp.StatusCode))
	}

	return obj, nil
}
