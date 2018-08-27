// A simple setup package to cut down on boilerplate RMI/RPC setup in the examples
package simplermi

import (
	"errors"
	"net"
	"net/rpc"
	"strings"
)

// Just pass in os.Args, get back a client
func SimpleClient(args []string) (*rpc.Client, error) {
	if len(args) < 3 {
		return nil, errors.New("Insufficient arguments: Requires at least Host and Port")
	}
	host := args[1]
	port := args[2]

	client, err := rpc.Dial("tcp4", strings.Join([]string{host, port}, ":"))
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Pass in a port, get back a server
func SimpleServer(port, rpcname string, rpctype interface{}) (*rpc.Server, net.Listener, error) {
	server := rpc.NewServer()
	if err := server.RegisterName(rpcname, rpctype); err != nil {
		return nil, nil, err
	}

	listener, err := net.Listen("tcp4", ":"+port)
	if err != nil {
		return nil, nil, err
	}

	return server, listener, nil
}
