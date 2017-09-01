package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mvpestretsov/go-eth-samples/lib"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"sync"
)

var gNonce uint64
var mutex sync.Mutex

func main() {

	// Create RPC connection to a remote node
	conn, err := ethclient.Dial(lib.GetNodeAddress())
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v\n", err)
	}
	log.Printf("conn:%+v", conn)
	// read key
	auth, err := bind.NewTransactor(strings.NewReader(lib.GetKey()), "qwe123")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v\n", err)
	}
	log.Printf("auth:%+v", auth)

	nonce, err := lib.GetActualNonce(conn, common.HexToAddress(lib.GetOwnerAddress()))
	if err != nil {
		log.Fatalf("Failed to get nonce: %v\n", err)
	}
	log.Printf("base nonce:%d ", nonce)

	setNonce(nonce - 1)
	start := time.Now()
	tryAmount := 2

	//TODO here
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func setNonce(i uint64) {
	mutex.Lock()
	gNonce = i
	mutex.Unlock()
}

func copyAuth(s *bind.TransactOpts) *bind.TransactOpts {
	clone := *s   // This is where we essentially make a new struct
	return &clone // Return reference to a new struct
}
func copyClient(s *ethclient.Client) *ethclient.Client {
	clone := *s   // This is where we essentially make a new struct
	return &clone // Return reference to a new struct
}

func copyTransaction(s *types.Transaction) *types.Transaction {
	clone := *s   // This is where we essentially make a new struct
	return &clone // Return reference to a new struct
}
