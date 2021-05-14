package bid

import "net/http"

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

// Route struct models the concept of route by specifying  route name, http method,
// path, and controller api method
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Controller  corresponds to a web api controller with methods to handle all available routes
type Controller struct {
	blockChain *BlockChain
	currentNodeUrl string
}
