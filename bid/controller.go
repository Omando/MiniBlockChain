/* Define all API methods for each available route. These methods are defined on the
Controller struct. Recall that the Controller struct holds the current node's URL
and own copy of the blockchain */
package bid

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
	"bidder_name": "YD",
	"auction_id": 100,
	"bid_value": "123.45"
}
*/
func (c *Controller) RegisterAndBroadcastBid(writer http.ResponseWriter, request *http.Request) {
	// Read body from request and check for errors
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("RegisterAndBroadcastBid error: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Parse the bid (in json) and convert to Bid object
	var bid Bid
	err = json.Unmarshal(body, &bid)
	if err != nil {
		log.Printf("RegisterAndBroadcastBid error: %s", err)
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// We have a Bid object. Register it in the blockchain
	c.blockChain.RegisterBid(bid)

	// Broadcast to all other available nodes
	for _, node := range c.blockChain.NetworkNodes {
		if node != c.currentNodeUrl {
			// Call RegisterBid on this node's controller
			DoPostCall(node+"/bid", body)
		}
	}

	// Return success to caller
	sendResponse(writer, http.StatusCreated, "RegisterAndBroadcastBid", "Bid created and broadcast successfully")
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

// sendResponse sends a standard response to all controller api method
func sendResponse(writer http.ResponseWriter, statusCode int, methodName string, message string) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(statusCode)
	var apiResponse ApiResponse = ApiResponse{
		Name:   methodName,
		Status: message,
		Time:   time.Now(),
	}
	data, _ := json.Marshal(apiResponse)
	writer.Write(data)
}

func DoPostCall(url string, body []byte) error {
	contentType := "application/json;charset=UTF-8"

	/* A Buffer is a variable-sized buffer of bytes with Read and Write methods.
	Recall the definition of io.Reader:
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
	Buffer implements the Reader interface as follows:
		func (b *Buffer) Read(p []byte) (n int, err error) {...}
	A Buffer instance can therefore be used as an io.Reader in http.Post
	*/
	var buffer *bytes.Buffer = bytes.NewBuffer(body)
	response, err := http.Post(url, contentType, buffer)
	if err != nil {
		log.Printf("Failed to POST call to %s: %s", url,  err)
		return err
	}

	response.Body.Close()
	return nil
}
