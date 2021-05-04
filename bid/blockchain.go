package bid

// RegisterBid registers a bid in the blockchain
func (b *BlockChain) RegisterBid(bid Bid) bool {

}

// RegisterNode registers a node in the blockchain
func (b *BlockChain) RegisterNode(node string) bool {

}

// CreateNewBlock create new block and appends it to the blockchain
func (b *BlockChain) CreateNewBlock(nonce int, previousBlockHash string, hash string) Block {

}

// GetLastBlock gets last block in the chain
func (b *BlockChain) GetLastBlock() Block {

}

// HashBlock calculates hash value for the given parameters
func (b *BlockChain) HashBlock(previousBlockHash string, currentBlockData string, nonce int) string {

}

// ProofOfWork
func (b *BlockChain) ProofOfWork (previousBlockHash string, currentBlockData string) int {

}

// CheckNewBlockHash
func (b *BlockChain) CheckNewBlockHash(newBlock Block) bool {
	return
}

// ChainIsValid checks if the entire block chain is valid
func (b *BlockChain) ChainIsValid() bool {

}



