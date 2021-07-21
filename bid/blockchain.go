package bid

import (
	"crypto/sha256"
	"encoding/base64"
	"hash"
	"strconv"
	"time"
)

// RegisterBid registers a bid in the blockchain
func (b *BlockChain) RegisterBid(bid Bid) {
	b.PendingBids = append(b.PendingBids, bid)
}

// RegisterNode registers a node in the blockchain if it does not already exist
func (b *BlockChain) RegisterNode(node string) bool {
	// Add node if it does not exist, else do nothing
	if !b.NetworkNodes[node] {
		b.NetworkNodes[node] = true
		return true		// node added
	}
	return false		// node already exists
}

// GetLastBlock gets last block in the chain
func (b *BlockChain) GetLastBlock() Block {
	return b.Chain[len(b.Chain)-1]
}

// CreateNewBlock create new block and appends it to the blockchain
func (b *BlockChain) CreateNewBlock(nonce int, previousBlockHash string, hash string) Block {
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

// HashBlock calculates hash value for the given parameters
func (b *BlockChain) HashBlock(previousBlockHash string, currentBlockData string, nonce int) string {
	//  Construct the string to hash from input data
	var stringToHash string = previousBlockHash + currentBlockData + strconv.Itoa(nonce)

	// Perform a sha256 hash. Standard logic to perform a hash:
	// 	h := sha256.New()
	//	h.Write([]byte("hello world\n"))  --> The array of bytes represents a very log sequence of 1s and 0s
	//	fmt.Printf("%x", h.Sum(nil))      --> prints: 948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a44
	var hash hash.Hash = sha256.New()
	hash.Write([]byte(stringToHash))

	// The output of sha256 is a binary value (expressed as an array of bytes). Encoding
	// is used to represent the hash as a string of characters, without increasing the size too much.
	// Note the use of hash.Sum(nil) to convert the hash into a byte slice as required by EncodeToString
	var base64Hash string = base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return base64Hash
}

// ProofOfWork increments a nonce until the hash value starts with a specific string value
func (b *BlockChain) ProofOfWork (previousBlockHash string, currentBlockData string) int {
	// Starting value for nonce
	nonce := -1
	inputFormat := ""

	// Increment the nonce until the SHA256 hash of the block data returns a string starting with “0000”
	for inputFormat != "0000" {
		nonce = nonce + 1
		var hashed string = b.HashBlock(previousBlockHash, currentBlockData, nonce)
		inputFormat = hashed[0:4]
	}

	return nonce
}

// CheckNewBlockHash
// A new candidate block is validated by checking its PreviousBlockHash and Index fields
// with our copy of the blockchain
func (b *BlockChain) CheckNewBlockHash(newBlock Block) bool {
	var lastBlock Block = b.GetLastBlock()
	return 	lastBlock.Hash == newBlock.PreviousBlockHash &&
			lastBlock.Index == newBlock.Index - 1
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


