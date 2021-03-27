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

