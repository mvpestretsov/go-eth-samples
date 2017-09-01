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
	//setNonce(0);

	//start to measure execution time
	start := time.Now()
	tryAmount := 2

	results := make(chan Result)

	for i := 0; i < tryAmount; i++ {
		req := Request{i, getNextNonce(), "HER", results, copyAuth(auth), copyClient(conn)}
		go runRequest(&req)
	}

	overall := waitDeployResult(tryAmount, results)
	log.Printf("%+v", overall)

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

type Result struct {
	requestNumber int
	requestNonce  uint64
	status        string
}

type Request struct {
	requestNumber int
	requestNonce  uint64
	data          string
	response      chan Result
	auth          *bind.TransactOpts
	conn          *ethclient.Client
}

func getNextNonce() uint64 {
	mutex.Lock()
	defer mutex.Unlock()
	gNonce = gNonce + 1
	return gNonce
}

func runRequest(request *Request) {
	var res Result
	tran := lib.BuyInfoContract{request.requestNonce, request.data}
	log.Printf("tran %+v", tran)
	log.Printf("request %+v", request)
	_, tx, _, err := lib.DeployBuyInfoContract(request.auth, request.conn, &tran)
	if err != nil {
		res = Result{requestNonce: request.requestNonce, requestNumber: request.requestNumber, status: err.Error()}
	} else {
		receipt, err := lib.MineTransaction(request.conn, tx)
		if receipt != nil {
			res = Result{requestNonce: request.requestNonce, requestNumber: request.requestNumber, status: "well done"}

		} else {
			res = Result{requestNonce: request.requestNonce, requestNumber: request.requestNumber, status: err.Error()}
		}
	}
	request.response <- res
	return
}

func waitDeployResult(tryAmount int, results chan Result) []Result {
	timeout := time.After(120 * time.Second)
	var deployed []Result
	for i := 0; i < tryAmount; i++ {
		log.Printf("%d: in waiting loop of deployed contracts", i)
		select {
		case result := <-results:
			deployed = append(deployed, result)
		case <-timeout:
			log.Println("return as is")
			return deployed
		}
	}
	return deployed
}

func runUpload() {

}
