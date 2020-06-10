package peer

import (
	"container/list"
	"time"
)


// Contact holds all related information for a peer in the K-Bucket
type Contact struct{

	// Id of the peer
	dhtId ID

	LastUsefulAt time.Time

	// Added At is the time this peer was added to the routing table
	AddedAt time.Time

	// if a bucket is full, this peer can be replaced to make space for a new peer.
	replaceable bool
}


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

func (kb *KBucket) GetAllPeers() []Contact {

	var peers []Contact

	for elem := kb.List.Front(); elem != nil; elem = elem.Next(){
		p := elem.Value.(*Contact)
		peers = append(peers, *p)
	}
	return peers
}

// Update is used to update the k-Bucket based on the
// original Kademplia algorithm
func (kb *KBucket) Update(id ID){

	elem := kb.find(id)
	if elem != nil{
		//kb.List.MoveToBack(elem)
	}else{
		if !kb.isFull(){
			kb.List.PushBack(&Contact{})
		}else{
			/* TODO: Ping the node from the front and see if it responds , if contact fails add the node to the tail
			and drop the element from the front, element from the from is the oldest contact node in the k-bucket.
			*/
		}
	}

}

// find returns the kbucket with the given Id if it exists
// returns nil if the peerId does not exist in the kBucket
func (kb *KBucket) find(id ID) *Contact {
	for elem := kb.List.Front(); elem != nil; elem = elem.Next(){
		if id.Equals(elem.Value.(*Contact).dhtId) {
			return elem.Value.(*Contact)
		}
	}
	return nil
}

// remove deletes the kbucket with the given Id it exists and returns true
// returns false if the kbucket does not exist in the kBucket
func (kb *KBucket) remove(id ID) bool {

	for elem := kb.List.Front(); elem != nil; elem = elem.Next(){
		if id.Equals(elem.Value.(*Contact).dhtId){
			kb.List.Remove(elem)
			return true
		}
	}
	return false
}

func (kb *KBucket) isFull() bool {

	//return kb.List.Len() >= kademlia.IDLENGTH
	return true
}

// Split splits a bucket peers into two buckets
// TODO: Further examination to understand how it works for Kademlia
func (kb *KBucket)Split(cpl int, target ID) * KBucket{

	newBucket := NewBucket()

	elem := kb.List.Front()

	for elem != nil {

		pId := elem.Value.(*Contact).dhtId
		peerCpl := CommonPrefixLen(pId, target)
		if peerCpl > cpl {

			cur := elem
			newBucket.List.PushBack(elem.Value)
			elem = elem.Next()
			kb.List.Remove(cur)
			continue
		}

		elem = elem.Next()
	}

	return newBucket
}

func (kb *KBucket) maxCommonPrefix(target ID) uint{

	maxCpl := uint(0)
	for elem := kb.List.Front(); elem != nil; elem = elem.Next(){
		cpl := uint(CommonPrefixLen(elem.Value.(*Contact).dhtId, target))
		if cpl > maxCpl {
			maxCpl = cpl
		}
	}
	return maxCpl
}