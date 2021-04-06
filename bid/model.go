package bid

// Bid bid information
type Bid struct {
	BidderName string `json:"bidder_name"`
	AuctionId  int    `json:"auction_id"`
	BidValue   int    `json:"bid_value"`
}

type Bids []Bid

// Block Basic structure of a blockchain block
type  Block struct {
	Index 				int 	`json:"index"`
	Timestamp 			int64	`json:"timestamp"`
	Bids 				Bids	`json:"bids"`
	Nonce 				int		`json:"nonce"`
	Hash				string	`json:"hash"`
	PreviousBlockHash	string 	`json:"previous_block_hash"`
}

type Blocks []Block

type Nodes []string

// Blockchain Basic structure of a blockchain consists of three collections:
// blocks, pending bids, and available network nodes
type BlockChain struct {
	Chain        Blocks   `json:"chain"`
	PendingBids  Bids     `json:"pending_bids"`
	NetworkNodes Nodes `json:"network_nodes"`
}

// BlockData is used in hash calculations
type BlockData struct {
	Index string
	Bids Bids
}

type Controller struct {
	blockChain *BlockChain
	currentNodeUrl string
}