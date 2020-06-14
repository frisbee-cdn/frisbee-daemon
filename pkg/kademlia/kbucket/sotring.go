package kbucket

import (
	"container/list"
	"sort"
)

type peerDistance struct {
	p        ID
	distance ID
}

type peerDistanceSorter struct {
	peers  []peerDistance
	target ID
}

func (pds *peerDistanceSorter) Less(i, j int) bool {
	return pds.peers[i].distance.less(pds.peers[j].distance)
}

func (pds *peerDistanceSorter) Len() int { return len(pds.peers) }
func (pds *peerDistanceSorter) Swap(a, b int) {
	pds.peers[a], pds.peers[b] = pds.peers[b], pds.peers[a]
}

func (pds *peerDistanceSorter) appendPeer(p ID, pDhtId ID) {
	pds.peers = append(pds.peers, peerDistance{
		p:        p,
		distance: XOR(pds.target, pDhtId),
	})
}

func (pds *peerDistanceSorter) appendPeersFromList(l *list.List) {
	for elem := l.Front(); elem != nil; elem = elem.Next() {
		pds.appendPeer(elem.Value.(*Contact).Id, elem.Value.(*PeerInfo).dhtId)
	}
}

func (pds *peerDistanceSorter) sort() {
	sort.Sort(pds)
}

func SortClosestPeers(peers []ID, target ID) []ID {
	sorter := peerDistanceSorter{
		peers:  make([]peerDistance, 0, len(peers)),
		target: target,
	}
	for _, p := range peers {
		sorter.appendPeer(p, ConvertPeerID(p))
	}
	sorter.sort()
	out := make([]ID, 0, sorter.Len())
	for _, p := range sorter.peers {
		out = append(out, p.p)
	}
	return out
}