[Gorilla]
https://www.gorillatoolkit.org/
https://github.com/gorilla/mux
to install gorilla/mux:
$ go get github.com/gorilla/mux

[Code Notes]
    [Type aliases]
Note how the use of 'type aliases' simplifies and clarifies code:

type  Block struct { ... }
type Blocks []Block
type Nodes []string
type BlockChain struct {
	Chain        Blocks     // Chain is a []Block
	NetworkNodes Nodes      // NetworkNodes ia a []string
	...
}

    [Useful code blocks]
port := os.Args[1]      // getting command line arguments

