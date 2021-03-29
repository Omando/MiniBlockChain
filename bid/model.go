package bid

type Bid struct {
	BidderName string `json:"bidder_name"`
	AuctionId  int    `json:"auction_id"`
	BidValue   int    `json:"bid_value"`
}

type Bids []Bid

// Block Basic structure of the blockchain
type  Block struct {
	Index 				int 	`json:"index"`
	Timestamp 			int64	`json:"timestamp"`
	Bids 				Bids	`json:"bids"`
	Nonce 				int		`json:"nonce"`
	Hash				string	`json:"hash"`
	PreviousBlockHash	string 	`json:"previous_block_hash"`
}

type Blocks []Block

type BlockChain struct {
	Chain        Blocks   `json:"chain"`
	PendingBids  Bids     `json:"pending_bids"`
	NetworkNodes []string `json:"network_nodes"`
}

// BlockData is used in hash calculations
type BlockData struct {
	Index string
	Bids Bids
}