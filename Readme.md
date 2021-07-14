# Setup and running

## Links
[Article](https://medium.com/coinmonks/my-blockchain-in-go-8e2d1a853a84#06ce)  
[GitHub](https://github.com/morris-ribs/BlockchainBetSportGolang)
## Packages
[Gorilla](https://www.gorillatoolkit.org/)  
[Gorilla GitHub](https://github.com/gorilla/mux)  
To install gorilla/mux:  
```
go get github.com/gorilla/mux
```

## Running
- [x] ```go run main.go 9000``` (or Shift+F9 to run in debugger)  
- [x] Run postman and invoke API Methods

# Code Notes

---
## Type aliases
Note how the use of _type aliases_ simplifies and clarifies code:
```go
type Block struct { ... }
type Blocks []Block
type Nodes []string
type BlockChain struct {
	Chain        Blocks     // Chain is a []Block 
	// NetworkNodes Nodes      // NetworkNodes ia a []string
	...
}
```
## Command line args
```go
port := os.Args[1]      // getting command line arguments
```

## Convert to json
```go
/* Use this code when sending data as JSON over http. Note this signature:
func Marshal(v interface{}) ([]byte, error) */

// Define a struct to hold some data 
type ColorGroup struct {
    ID     int
    Name   string
    Colors []string
}

// Create and initialize a ColorGroup object
group := ColorGroup {
    ID:     1,
    Name:   "Reds",
    Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

// Convert the group object into an array of bytes holding JSON data
b, err := json.Marshal(group)
os.Stdout.Write(b)      // Output: {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
```

## Convert from json
```go
/* Use this code when receiving data as JSON from http. Note this signature:
func Unmarshal(data []byte, v interface{}) error */

// Some JSON data. Note the syntax below to create an array of bytes
// containing JSON data
var jsonBlob = []byte(`[
    {"Name": "Platypus", "Order": "Monotremata"},
    {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)

// JSON will be converted into this object
type Animal struct {
    Name  string
    Order string
}

// Convert the array of bytes holding JSON data into objects
var animals []Animal
err := json.Unmarshal(jsonBlob, &animals)
fmt.Printf("%+v", animals)  // Output: [{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
```

## Convert object to []byte then string
```go
/* Use this code when getting binary data in the form of []bye (i.e., from hashing)
to convert it to a human-readable string */

// Get some object
var myObject BlockData = BlockData{strconv.Itoa(lastBlock.Index), c.blockChain.PendingBids}

// Convert the object into an array of bytes
var myByteArray , _ = json.Marshal(myObject)

// Convert the array of bytes into a string
var myString  = base64.URLEncoding.EncodeToString(myByteArray)
```

## Reading data from an incoming HTTP request
```go
/* Use ioutil.ReadAll then json.Unmarshal. Note this signature
func ReadAll(r io.Reader) ([]byte, error) */

defer request.Body.Close()

// Read the incoming data into an array of bytes
body, err := ioutil.ReadAll(request.Body)

// Convert the array of bytes to an object
var bid Bid
err = json.Unmarshal(body, &bid)
```

## Replying to an incoming HTTP request
```go
// Write any requested headers
writer.Header().Set("Content-Type", "application/json; charset=UTF-8")

// Write the status code
writer.WriteHeader(some-status-code)

// Create and initialize an object containing the data to return to the client
var apiResponse ApiResponse = ApiResponse {
	Name:   methodName,
	Status: message,
	Time:   time.Now(),
}

// Convert the object to an array of bytes
data, _ := json.Marshal(apiResponse)

// Send the array of bytes back to the sender 
writer.Write(data)
```

## Making a POST call
```go
// Convert some object to an array of bytes
body, _ := json.Marshal(newBlock)                   // body is []byte

// A POST call requires data in the form of a pointer to bytes.Buffer
var buffer *bytes.Buffer = bytes.NewBuffer(body)        

// Make the POST  call
response, err := http.Post(url, "application/json;charset=UTF-8", buffer)
response.Body.Close()
```

## Making a GET call
```go

```

Markdown tutorial: <https://www.markdownguide.org/basic-syntax/>