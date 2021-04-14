package bid

import "net/http"

// Index GET/
func (c *Controller) Index(writer http.ResponseWriter, request *http.Request) {

}

// GetBlockChain GET /blockchain
func (c *Controller) GetBlockChain(writer http.ResponseWriter, request *http.Request) {

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

// RegisterNode POST /register-node
func (c *Controller) RegisterNode(writer http.ResponseWriter, request *http.Request) {

}

// RegisterNodesBulk POST /register-nodes-bulk
func (c *Controller) RegisterNodesBulk(writer http.ResponseWriter, request *http.Request) {

}

// RegisterAndBroadcastNode POST /register-and-broadcast-node
func (c *Controller) RegisterAndBroadcastNode(writer http.ResponseWriter, request *http.Request) {

}

// ReceiveNewBlock POST /review-new-block
func (c *Controller) ReceiveNewBlock(writer http.ResponseWriter, request *http.Request) {

}

// Consensus GET /consensus
func (c *Controller) Consensus(writer http.ResponseWriter, request *http.Request) {

}


