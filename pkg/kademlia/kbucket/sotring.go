package kbucket

import (
	"container/list"
	"sort"
)

type peerDistance struct {
	c        *Contact
	distance ID
}

type peerDistanceSorter struct {
	peers  []peerDistance
	target ID
}

func (pds *peerDistanceSorter) Less(i, j int) bool {
	return pds.peers[i].distance.Less(pds.peers[j].distance)
}

func (pds *peerDistanceSorter) Len() int { return len(pds.peers) }
func (pds *peerDistanceSorter) Swap(a, b int) {
	pds.peers[a], pds.peers[b] = pds.peers[b], pds.peers[a]
}

func (pds *peerDistanceSorter) appendPeer(contact *Contact) {
	pds.peers = append(pds.peers, peerDistance{
		c:        contact,
		distance: XOR(pds.target, contact.Node.SelfKey),
	})
}

func (pds *peerDistanceSorter) appendPeersFromList(l *list.List) {
	for elem := l.Front(); elem != nil; elem = elem.Next() {
		pds.appendPeer(elem.Value.(*Contact))
	}
}

func (pds *peerDistanceSorter) sort() {
	sort.Sort(pds)
}

// func SortClosestPeers(peers []peer.ID, target ID) []peer.ID {
// 	sorter := peerDistanceSorter{
// 		peers:  make([]peerDistance, 0, len(peers)),
// 		target: target,
// 	}

// 	for _, p := range peers {
// 		hashId, _ := HashKey(p)
// 		sorter.appendPeer(p)
// 	}
// 	sorter.sort()
// 	out := make([]peer.ID, 0, sorter.Len())
// 	for _, p := range sorter.peers {
// 		out = append(out, p.p)
// 	}
// 	return out
// }
