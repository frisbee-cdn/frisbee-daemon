package rpc

import (
	models "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// Transport abstraction of RPC service functionalities
type Transport interface {
	GetSuccessor(*models.Node) (*models.Node, error)
	FindSuccessor(*models.Node) error

	//Storage

	GetKey(*models.Node, string) error
	SetKey(*models.Node, string, string) error
	DeleteKey(*models.Node, string) error
}
