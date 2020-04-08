package mock

import (
	"github.com/ElrondNetwork/elrond-go/data/block"
	"github.com/ElrondNetwork/elrond-go/p2p"
)

// MiniBlocksResolverStub -
type MiniBlocksResolverStub struct {
	RequestDataFromHashCalled      func(hash []byte, epoch uint32) error
	RequestDataFromHashArrayCalled func(hashes [][]byte, epoch uint32) error
	ProcessReceivedMessageCalled   func(message p2p.MessageP2P) error
	GetMiniBlocksCalled            func(hashes [][]byte) (block.MiniBlockSlice, [][]byte)
	GetMiniBlocksFromPoolCalled    func(hashes [][]byte) (block.MiniBlockSlice, [][]byte)
	SetNumPeersToQueryCalled       func(intra int, cross int)
	GetNumPeersToQueryCalled       func() (int, int)
}

// SetNumPeersToQuery -
func (mbrs *MiniBlocksResolverStub) SetNumPeersToQuery(intra int, cross int) {
	if mbrs.SetNumPeersToQueryCalled != nil {
		mbrs.SetNumPeersToQueryCalled(intra, cross)
	}
}

// GetNumPeersToQuery -
func (mbrs *MiniBlocksResolverStub) GetNumPeersToQuery() (int, int) {
	if mbrs.GetNumPeersToQueryCalled != nil {
		return mbrs.GetNumPeersToQueryCalled()
	}

	return 2, 2
}

// RequestDataFromHash -
func (mbrs *MiniBlocksResolverStub) RequestDataFromHash(hash []byte, epoch uint32) error {
	return mbrs.RequestDataFromHashCalled(hash, epoch)
}

// RequestDataFromHashArray -
func (mbrs *MiniBlocksResolverStub) RequestDataFromHashArray(hashes [][]byte, epoch uint32) error {
	return mbrs.RequestDataFromHashArrayCalled(hashes, epoch)
}

// ProcessReceivedMessage -
func (mbrs *MiniBlocksResolverStub) ProcessReceivedMessage(message p2p.MessageP2P, _ p2p.PeerID) error {
	return mbrs.ProcessReceivedMessageCalled(message)
}

// GetMiniBlocks -
func (mbrs *MiniBlocksResolverStub) GetMiniBlocks(hashes [][]byte) (block.MiniBlockSlice, [][]byte) {
	return mbrs.GetMiniBlocksCalled(hashes)
}

// GetMiniBlocksFromPool -
func (mbrs *MiniBlocksResolverStub) GetMiniBlocksFromPool(hashes [][]byte) (block.MiniBlockSlice, [][]byte) {
	return mbrs.GetMiniBlocksFromPoolCalled(hashes)
}

// IsInterfaceNil returns true if there is no value under the interface
func (mbrs *MiniBlocksResolverStub) IsInterfaceNil() bool {
	return mbrs == nil
}
