package rpc

// Transport abstraction of RPC service functionalities
type Transport interface {

	// Check Status

	Ping(string) string
}
