package bid

import "time"

// CreateNewBlock create new block and appends it to the blockchain
func (b *BlockChain) CreateNewBlock(nonce int, previousBlockHash string, hash string) Block {

	// Create and initialize a new block
	newBlock := Block{
		Index:             len(b.Chain) + 1,	// Length of chain + 1
		Timestamp:         time.Now().UnixNano(),
		Bids:              b.PendingBids,
		Nonce:             nonce	,
		Hash:              hash,
		PreviousBlockHash: previousBlockHash,
	}

	// There are no pending bids when a new block is created
	b.PendingBids = Bids{}

	// Add this new block to the chain
	b.Chain = append(b.Chain, newBlock )
	
	return  newBlock
}

// RegisterBid registers a bid in the blockchain
func (b *BlockChain) RegisterBid(bid Bid) bool {
	panic("not implemented")
}

// RegisterNode registers a node in the blockchain
func (b *BlockChain) RegisterNode(node string) bool {
	panic("not implemented")
}

// GetLastBlock gets last block in the chain
func (b *BlockChain) GetLastBlock() Block {
	panic("not implemented")
}

// HashBlock calculates hash value for the given parameters
func (b *BlockChain) HashBlock(previousBlockHash string, currentBlockData string, nonce int) string {
	panic("not implemented")
}

// ProofOfWork
func (b *BlockChain) ProofOfWork (previousBlockHash string, currentBlockData string) int {
	panic("not implemented")
}

// CheckNewBlockHash
func (b *BlockChain) CheckNewBlockHash(newBlock Block) bool {
	panic("not implemented")
}

// ChainIsValid checks if th/e entire block chain is valid
func (b *BlockChain) ChainIsValid() bool {
	panic("not implemented")
}

// GetBidsForMatch gets all bids for a specific auction
func (b* BlockChain) GetBidsForAuction(autionId string) Bids {
	panic("not implemented")
}

// GetBidsForPlayer gets all bids for a specific player id
func (b *BlockChain) GetBidsForPlayer(playerId string) Bids {
	panic("not implemented")
}


