package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	u "github.com/mvpestretsov/go-eth-samples/lib"
	"math/big"
	"strings"
	"time"
)

func printTokenBalance(tkn *MyToken, ownerAddress string) {
	balance, err := tkn.BalanceOf(nil, common.HexToAddress(ownerAddress))
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	log.Println("Balance:", balance)
}

func deployTokenContract(auth *bind.TransactOpts, conn bind.ContractBackend) (common.Address, *types.Transaction, *MyToken) {
	address, tx, tkn, err := DeployMyToken(auth, conn, big.NewInt(1000), "MT", 0, "MT")
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	log.Printf("Contract pending deploy: 0x%x\n", address)
	log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	return address, tx, tkn
}

func main() {
	// Create RPC connection to a remote node
	conn, err := ethclient.Dial(u.GetNodeAddress())
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// read key
	auth, err := bind.NewTransactor(strings.NewReader(u.GetKey()), "qwe123")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	token, err := NewMyToken(common.HexToAddress(u.GetTokenContract()), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	printTokenBalance(token, u.GetOwnerAddress())
	// Create an authorized transactor and spend 1 unicorn

	tx, err := token.Transfer(auth, common.HexToAddress("0x0000000000000000000000000000000000000000"), big.NewInt(1))
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
	}
	log.Printf("Transfer pending: 0x%x\n", tx.Hash())
	ctx, stop := context.WithTimeout(context.Background(), 15000*time.Millisecond)
	defer stop()
	receipt, err := bind.WaitMined(ctx, conn, tx)

	if err != nil {
		log.Fatalf("Failed to get receipt: %v", err)
	}

	if receipt != nil {
		log.Printf("Receipt: %v", receipt)
	}

	printTokenBalance(token, u.GetOwnerAddress())
	// Instantiate the contract and display its name
	/*
		token, err := mytoken.NewMyToken(common.HexToAddress("0x8cf72a7473fd402b0bd46a29e83fa040e7486e7f"), conn)
		if err != nil {
			log.Fatalf("Failed to instantiate a Token contract: %v", err)
		}
		name, err := token.Name(nil)
		if err != nil {
			log.Fatalf("Failed to retrieve token name: %v", err)
		}
		fmt.Println("Token name:", name)

		balance, err=token.BalanceOf(nil,common.HexToAddress("0x44674ad384bc12ee853c3918ea96bbff7bbf3775"));
		fmt.Println("Balance one more time:", balance)
	*/
}
