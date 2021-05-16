[1. Gorilla]
https://www.gorillatoolkit.org/
https://github.com/gorilla/mux
to install gorilla/mux:
$ go get github.com/gorilla/mux

[Running]
go build main.go
go run main.go 9000 (or Shift+F9 to run in debugger)
Run postman and invoke API Methods

[3. Code Notes]
[3.1 Type aliases]
Note how the use of 'type aliases' simplifies and clarifies code:

type  Block struct { ... }
type Blocks []Block
type Nodes []string
type BlockChain struct {
	Chain        Blocks     // Chain is a []Block
	NetworkNodes Nodes      // NetworkNodes ia a []string
	...
}

[3.2 Useful code blocks]
port := os.Args[1]      // getting command line arguments
