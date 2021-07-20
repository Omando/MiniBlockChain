package bid

import (
	"net/http"
	"time"
)

// Bid bid information
type Bid struct {
	BidderName string 		`json:"bidder_name"`
	AuctionId  int    		`json:"auction_id"`
	BidValue   float32    	`json:"bid_value,string"`	//Note of use string
}
type Bids []Bid

// Nodes is an alias for an array of strings where each string is the address of a node
type Nodes []string

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

// BlockData is used in mining to identify
type BlockData struct {
	Index string
	Bids Bids
}

// BlockChain basic structure of a blockchain consists of three collections:
// blocks, pending bids, and available network nodes
type BlockChain struct {
	Chain        Blocks   			`json:"chain"`
	PendingBids  Bids     			`json:"pending_bids"`
	NetworkNodes map[string]bool 	`json:"network_nodes"`
}

// Controller corresponds to a web api controller with methods to handle all available routes
type Controller struct {
	blockChain *BlockChain
	currentNodeUrl string
}

// Route struct models the concept of route by specifying route name, http method,
// path, and controller api method
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type ApiResponse struct {
	Name string
	Status string
	Time time.Time
}