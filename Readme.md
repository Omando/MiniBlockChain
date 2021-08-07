# Setup and running

---
## Packages
[Gorilla](https://www.gorillatoolkit.org/)  
[Gorilla GitHub](https://github.com/gorilla/mux)  
To install gorilla/mux:  
```$ go get github.com/gorilla/mux```

## Running
- [x] ```go run main.go 9000``` (or Shift+F9 to run in debugger)  
- [x] Run postman and invoke API Methods

# Code Notes

---
## Type aliases
Note how the use of _type aliases_ simplifies and clarifies code:
```go
type  Block struct { ... }
type Blocks []Block
type Nodes []string
type BlockChain struct {
	Chain        Blocks     // Chain is a []Block
	NetworkNodes Nodes      // NetworkNodes ia a []string
	...
}
```
## Command line args
```go
port := os.Args[1]      // getting command line arguments
```

## Convert to json
```go
// Typically used  to send data over http. Note this signature:
//  func Marshal(v interface{}) ([]byte, error)
type ColorGroup struct {
    ID     int
    Name   string
    Colors []string
}

group := ColorGroup{
    ID:     1,
    Name:   "Reds",
    Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

b, err := json.Marshal(group)
os.Stdout.Write(b)      // Output: {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
```

## Convert from json
```go
// Use this code to receive json data from http. Note this signature:
//  func Unmarshal(data []byte, v interface{}) error
var jsonBlob = []byte(`[
    {"Name": "Platypus", "Order": "Monotremata"},
    {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)

type Animal struct {
    Name  string
    Order string
}

var animals []Animal
err := json.Unmarshal(jsonBlob, &animals)
fmt.Printf("%+v", animals)  // Output: [{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
```

## Convert object to []byte then string
```go
// Use this when getting binary data in the form of []bye (i.e., from hashing), then
// convert it to human-readable string
var newBlockData BlockData = BlockData{strconv.Itoa(lastBlock.Index), c.blockChain.PendingBids}
var newBlockDataAsBinary , _ = json.Marshal(newBlockData)
var newBlockDataAsString  = base64.URLEncoding.EncodeToString(newBlockDataAsBinary)
```

## Reading data from an incoming request
```go
// Use ioutil.RealAll then json.Unmarshal
defer request.Body.Close()
body, err := ioutil.ReadAll(request.Body)
var bid Bid
err = json.Unmarshal(body, &bid)
```

## Replying to incoming request
```go
writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
writer.WriteHeader(some-status-code)
var apiResponse ApiResponse = ApiResponse{      // Some other struct to send back to client
		Name:   methodName,
		Status: message,
		Time:   time.Now(),
}
data, _ := json.Marshal(apiResponse)
writer.Write(data)
```

## Making a post call
```go
contentType := "application/json;charset=UTF-8"
body, _ := json.Marshal(newBlock)                       // newBlock is some object. blockToBroadCast is []byte
var buffer *bytes.Buffer = bytes.NewBuffer(body)        // body is []byte
response, err := http.Post(url, contentType, buffer)
response.Body.Close()
```