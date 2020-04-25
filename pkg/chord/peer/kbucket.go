package peer

import (
	"container/list"
	"github/frisbee-cdn/frisbee-daemon/pkg/util"
)

// KBucket is a list of routing addresses of other nodes in
// the network.
type KBucket struct {
	*list.List
}

// NewBucket is used to create an empty KBucket
func NewBucket() *KBucket {
	return &KBucket{
		list.New(),
	}
}

//Add
func (kb *KBucket) Add(node Node){

	kb.PushBack(node)

}

func (kb *KBucket) isFull() bool {

	return kb.Len() >= util.IDLENGTH
}
