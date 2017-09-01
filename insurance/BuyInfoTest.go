package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	u "github.com/mvpestretsov/go-eth-samples/lib"
	//"math/big"
	"strings"
	"sync"
	"time"
	//"context"
	//"context"
	"github.com/mvpestretsov/go-eth-samples/insurance/util"
	"math/big"

	"fmt"
)

var mutex = &sync.Mutex{}
var deployCounter int = 0

func main() {

	// Create RPC connection to a remote node
	conn, err := ethclient.Dial(u.GetNodeAddress())
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v\n", err)
	}
	log.Printf("conn:%+v", conn)
	// read key
	auth, err := bind.NewTransactor(strings.NewReader(u.GetKey()), "qwe123")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v\n", err)
	}
	log.Printf("auth:%+v", auth)

	nonce, err := u.GetActualNonce(conn, common.HexToAddress(u.GetOwnerAddress()))
	if err != nil {
		log.Fatalf("Failed to get nonce: %v\n", err)
	}
	log.Printf("base nonce:%d ", nonce)

	//start to measure execution time
	start := time.Now()

	/*
		doConcurrency(1,1, func(i int) {
			// Make request here
			deployBuyInfoContract(copyAuth(auth),copyClient(conn),i,nonce)
		})
	*/

	var tryAmount int = 1

	/*
		var wg sync.WaitGroup


		for i := 0; i < tryAmount; i++ {
			//make local deployment
			contractA,tx,contract,err:=deployBuyInfoContract(copyAuth(auth),copyClient(conn),nonce + uint64(i))
			if err==nil {
				log.Printf("%d:address:0x%x\n", i, contractA)
				log.Printf("%d:contract:%+v", i, contract)
				log.Printf("%d:transaction:%v", i, tx.Hash().Hex())
				wg.Add(1)
				//mine deployed contract
				go  MineAndGetReceipt(&wg,copyTransaction(tx),copyClient(conn))
			}
		}
		wg.Wait()
	*/

	//make needed amount of contracts
	deployRequests := make(chan uint64, tryAmount)
	for i := 1; i <= tryAmount; i++ {
		deployRequests <- nonce + uint64(i)
	}
	close(deployRequests)

	//for each contract count make deploy request
	for req := range deployRequests {
		contractA, tx, contract, err := deployBuyInfoContract(copyAuth(auth), copyClient(conn), req)
		if err == nil {
			log.Printf("%d:address:0x%x\n", req, contractA)
			log.Printf("%d:contract:%+v", req, contract)
			log.Printf("%d:transaction:%v", req, tx.Hash().Hex())
			//mine deploy transaction
			dReceipt, _ := u.MineTransaction(copyClient(conn), tx)
			if dReceipt != nil {
				incrementSuccessCounter()
				//upload data to contract
				auth.GasLimit = big.NewInt(2551488)
				upTx, err := contract.UploadData(auth, "test")
				if err == nil {
					//mine upload transaction
					uReceipt, err := u.MineTransaction(copyClient(conn), upTx)
					if uReceipt != nil {
						//read uploaded data from transaction
						data, err := contract.EncryptedData(&bind.CallOpts{})
						if err != nil {
							log.Fatalf("Failed to get data: %v\n", err)
						}
						log.Printf("%d:contract:%v", req, data)
					} else {
						log.Printf("Failed to mine upload data: %v\n", err)
					}
				} else {
					log.Printf("Failed to upload data: %v\n", err)
				}
			} else {
				log.Printf("Failed to mine contract: %v\n", err)
			}
		}
	}

	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
	log.Printf("deployCounter=%d", deployCounter)
}

func genNonce(nonce uint64, tryAmount int) <-chan uint64 {
	out := make(chan uint64)
	go func() {
		for i := 0; i < tryAmount; i++ {
			out <- nonce + uint64(i)
		}
		close(out)
	}()
	return out
}

func doConcurrency(total, concurrency int, fn func(int)) {
	workQueue := make(chan int, concurrency)

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			for i := range workQueue {
				fn(i)
			}
			wg.Done()
		}()
	}
	for i := 0; i < total; i++ {
		workQueue <- i
	}
	close(workQueue)
	wg.Wait()
}

func MineAndGetReceipt(wg *sync.WaitGroup, tx *types.Transaction, conn *ethclient.Client) {
	defer wg.Done()
	log.Printf("mine transaction :%v", tx.Hash().Hex())
	receipt, _ := u.MineTransaction(conn, tx)
	if receipt != nil {
		incrementSuccessCounter()
	}
}

func incrementSuccessCounter() {
	mutex.Lock()
	deployCounter = deployCounter + 1
	mutex.Unlock()
}

func deployBuyInfoContract(auth *bind.TransactOpts, conn *ethclient.Client, nonce uint64) (common.Address, *types.Transaction, *insurance.BuyInfo, error) {
	address, tx, buyInfoC, err := insurance.DeployBuyInfo(auth, conn, common.HexToAddress(u.GetOwnerAddress()), "pk", fmt.Sprintf("%s%d", "dataHer", nonce), "idH", u.GetStubAddress(), big.NewInt(1000), big.NewInt(1000))
	if err != nil {
		log.Printf("Failed to deploy new BuyInfo contract: %v\n", err)
		return u.GetStubAddress(), nil, nil, err
	}
	log.Printf("Contract pending deploy: 0x%x\n", address)
	log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	return address, tx, buyInfoC, nil
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
