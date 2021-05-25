/* Define all API methods for each available route. These methods are defined on the
Controller struct. Recall that the Controller struct holds the current node's URL
and own copy of the blockchain */
package bid

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetBlockChain GET /blockchain
/* Retrieves the blockchain in JSON format. Typical output looks like this:
{
	"chain": [
		{
			"index": 1,
			"timestamp": 1627171722582903400,
			"bids": [],
			"nonce": 100,
			"hash": "0",
			"previous_block_hash": "0"
		}
	],
	"pending_bids": [],
	"network_nodes": []
}
*/
func (c *Controller) GetBlockChain(writer http.ResponseWriter, request *http.Request) {
	// Setup headers
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusOK)

	// Get the blockchain from the controller and convert it into JSON
	data, _ := json.Marshal(c.blockChain)
	
	// Send the blockchain back to the client
	writer.Write(data)
	return
}

// RegisterAndBroadcastBid POST /bid/broadcast
/* Register a bid in current blockchain and transmit to all nodes in the network. Typical body input
{
	"bidder_name": "YD"
	"auction_id": 100
	"bid_value": 1.99
}
*/
func (c *Controller) RegisterAndBroadcastBid(writer http.ResponseWriter, request *http.Request) {
	// Read body from request and check for errors
	defer request.Body.Close()
	jsonBid, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// We have bid as json. Parse the bid (in json) and convert into a Bid object
	var bid Bid
	err = json.Unmarshal(jsonBid, &bid)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.blockChain.PendingBids = append(c.blockChain.PendingBids, bid)

}

// RegisterBid POST/bid
/* This API method is called by RegisterAndBroadcastBid which registers an API bid locally
to transmit this bid to  other nodes of the network */
func (c *Controller) RegisterBid(writer http.ResponseWriter, request *http.Request) {

}

// Index GET/
func (c *Controller) Index(writer http.ResponseWriter, request *http.Request) {

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


