/* Define all API methods for each available route. These methods are defined on the
Controller struct. Recall that the Controller struct holds the current node's URL
and own copy of the blockchain */
package bid

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	c.registerBidImp(writer, request, true)		// Broadcast
}

// RegisterBid POST /bid
/* This method registers an API bid locally but does not transmit it. This happens when a user registers
a bid and broadcasts the bid to all other nodes (users) by calling RegisterAndBroadcastBid. Typical body input:
{
	"bidder_name": "YD",
	"auction_id": 100,
	"bid_value": "123.45"
}
*/
func (c *Controller) RegisterBid(writer http.ResponseWriter, request *http.Request) {
	c.registerBidImp(writer, request, false)		// Do not broadcast
}

// Mine GET /mine
// Mining works by getting the last block and calling ProofOfWork to find the nonce
// of the new block to be added. Once the nonce is found, a new block is created
// and added to the chain. Lastly the block is transmitted to all other nodes by
// calling ReceiveNewBlock on each available node
func (c *Controller) Mine(writer http.ResponseWriter, request *http.Request) {
	// Get hash of the last block in the blockchain
	var lastBlock Block = c.blockChain.GetLastBlock()
	var lastBlockHash string = lastBlock.Hash

	// To calculate proof of work, we need two items: the hash of the last block
	//(we have it above),  and data for the new block in the form of a string.
	// To collect data for the new block, we create a BlockData struct and then
	// convert this struct value to a string.
	// To convert a BlockData struct value to a string, we first convert it a []byte
	// using json.Marshal and then we use base64 encoding to get a string representation
	// of the []byte (recall, base64 only contains A–Z, a–z, 0–9, +, / and =)
	var newBlockData BlockData = BlockData{strconv.Itoa(lastBlock.Index), c.blockChain.PendingBids}
	var newBlockDataAsBinary , _ = json.Marshal(newBlockData)
	var newBlockDataAsString  = base64.URLEncoding.EncodeToString(newBlockDataAsBinary)

	// We now have both items required for proof of work. Run proof of work to get nonce
	var nonce int =  c.blockChain.ProofOfWork(lastBlockHash, newBlockDataAsString)

	// Now that we have the nonce, to create a new block we also need a hash for the new block
	var hash =  c.blockChain.HashBlock(lastBlockHash, newBlockDataAsString, nonce)
	var newBlock Block =  c.blockChain.CreateNewBlock(nonce, lastBlockHash, hash)

	// We have a new block! Broadcast it to all nodes (call ReceiveNewBlock on all nodes)
	blockToBroadcast, _ := json.Marshal(newBlock)
	c.broadcastToAllNodes("/receive-new-block", blockToBroadcast);

	// Let caller know that we've completed mining and broadcasting
	sendStandardResponse(writer, http.StatusOK, "Mine", "New block mined and broadcast")
}

// ReceiveNewBlock POST /receive-new-block
/* Receive and validate a new block. If validated, the new block is accepted, otherwise it is rejected */
func (c *Controller) ReceiveNewBlock(writer http.ResponseWriter, request *http.Request) {
	// Receive the new block (note the pattern: ioUtil.ReadAll followed by json.Unmarshal)
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("Failed to receive new block: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	var newBlock Block
	json.Unmarshal(body, &newBlock)

	// Process new block: if validated, add to the blockchain
	var message string = "New block has been rejected"
	var statusCode int = http.StatusInternalServerError
	if c.blockChain.CheckNewBlockHash(newBlock) {
		c.blockChain.PendingBids = Bids{}
		c.blockChain.Chain = append(c.blockChain.Chain, newBlock)
		message = "New block received and accepted"
		statusCode = http.StatusOK
	}
	// Send response back with the result of receiving this block
	sendStandardResponse(writer, statusCode, "ReceiveNewBlock", message)
}

// RegisterAndBroadcastNode POST /register-and-broadcast-node
/* When a node comes online, it finds the list of available nodes (how?), and for each node
calls its RegisterAndBroadcastNode passing itself as the new node. This function:
1. Will add the incoming node to its list of known nodes
2. For each node in its list of known nodes, calls RegisterNode passing this incoming node.
3. Calls RegisterNodesBulk passing all the current nodes of the network to the new node
Typical input looks like this
	{
		"newnodeurl":"http://address:port"
	}
*/
func (c *Controller) RegisterAndBroadcastNode(writer http.ResponseWriter, request *http.Request) {

}

// RegisterNode POST /register-node
/* When a new node comes online it calls RegisterAndBroadcastNode passing itself as the new node.
This function will add the incoming node to its list of known nodes, and then for each node in
its list of known nodes, calls RegisterNode passing this incoming node  */
func (c *Controller) RegisterNode(writer http.ResponseWriter, request *http.Request) {

}

// RegisterNodesBulk POST /register-nodes-bulk
/* When a new node comes online it calls RegisterAndBroadcastNode passing itself as the new node.
RegisterAndBroadcastNode will process the request and then it will call RegisterNodesBulk on the
new node passing all the current nodes of the network to the new node. This ensure that the new
node knows about all other nodes in the network */
func (c *Controller) RegisterNodesBulk(writer http.ResponseWriter, request *http.Request) {

}

// Index GET/
func (c *Controller) Index(writer http.ResponseWriter, request *http.Request) {

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

/* Helpers */
func (c *Controller) broadcastToAllNodes(api string, body []byte) {
	for _, node := range c.blockChain.NetworkNodes {
		if node != c.currentNodeUrl {
			doPostCall(node + api, body)
		}
	}
}


// Creates a Bid object from the body and adds the bid to the blockchain. The bid is conditionally
// broadcast to all other registered nodes
func (c *Controller) registerBidImp(writer http.ResponseWriter, request *http.Request, shouldBroadCast bool) {
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
	if shouldBroadCast {
		c.broadcastToAllNodes("/bid", body)
	}

	// Return success to caller
	sendStandardResponse(writer, http.StatusCreated, "RegisterAndBroadcastBid", "Bid created and broadcast successfully")
}

// sendStandardResponse sends a standard response from all controller api methods: send a content type,
// a status, and a ApiResponse object with additional data
func sendStandardResponse(writer http.ResponseWriter, statusCode int, methodName string, message string) {
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

// Do a post call to the given url. Typically used to inform other nodes of interesting changes
// such as a new block or a new node
func doPostCall(url string, body []byte) error {
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
