package bid

import (
	"encoding/json"
	"net/http"
)

// Index GET/
func (c *Controller) Index(writer http.ResponseWriter, request *http.Request) {

}

// GetBlockChain GET /blockchain
func (c *Controller) GetBlockChain(writer http.ResponseWriter, request *http.Request) {
	data, _ := json.Marshal(c.blockChain)
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	writer.Write(data)
	return
}

// RegisterAndBroadcastNode POST /register-and-broadcast-node
func (c *Controller) RegisterAndBroadcastNode(writer http.ResponseWriter, request *http.Request) {

}

// RegisterNode POST /register-node
func (c *Controller) RegisterNode(writer http.ResponseWriter, request *http.Request) {

}

// RegisterNodesBulk POST /register-nodes-bulk
func (c *Controller) RegisterNodesBulk(writer http.ResponseWriter, request *http.Request) {

}

// RegisterBid POST/bid
func (c *Controller) RegisterBid(writer http.ResponseWriter, request *http.Request) {

}

// RegisterAndBroadcastBid POST /bid/broadcast
func (c *Controller) RegisterAndBroadcastBid(writer http.ResponseWriter, request *http.Request) {

}

// Mine GET /mine
func (c *Controller) Mine(writer http.ResponseWriter, request *http.Request) {

}

// ReceiveNewBlock POST /review-new-block
func (c *Controller) ReceiveNewBlock(writer http.ResponseWriter, request *http.Request) {

}
// Consensus GET /consensus
func (c *Controller) Consensus(writer http.ResponseWriter, request *http.Request) {

}

// GetBidsForAuction GET /auction/{auctionId} retrieves all bids for an auction
func (c *Controller) GetBidsForAuction(writer http.ResponseWriter, request *http.Request) {

}

// GetBidsForPlayer GET /player/{playerId}
func (c *Controller) GetBidsForPlayer(writer http.ResponseWriter, request * http.Request) {

}


