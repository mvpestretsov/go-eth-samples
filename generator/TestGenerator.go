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

	"time"
	//"context"
	//"context"

	"math/big"

	"github.com/mvpestretsov/go-eth-samples/generator/contract"
)

func main() {

	key1 := `{"address":"d889637e948b737b335e8b4585d29b1132987b36","crypto":{"cipher":"aes-128-ctr","ciphertext":"ec50e452cac44aafbd678dcb9796c6a7ae7d1b36f04100c4a0fa6d9644dcae9b","cipherparams":{"iv":"b7c326a53d9be43280913b58ec5ae59c"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"c9db835948ce03276b10c74fe345f70decf0a0a78aae706ef411c2fc12d9e4bf"},"mac":"112d7257bb4e3ad09b442061d5e41e46aa469c88541d07be0a31355d38939887"},"id":"835484af-622e-4aaf-938f-e5d86d2aa526","version":3}`
	// Create RPC connection to a remote node
	conn, err := ethclient.Dial(u.GetNodeAddress())
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v\n", err)
	}
	log.Printf("conn:%+v", conn)
	// read key
	auth, err := bind.NewTransactor(strings.NewReader(key1), "qwe123")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v\n", err)
	}
	log.Printf("auth:%+v", auth)

	nonce, err := u.GetActualNonce(conn, common.HexToAddress(u.GetOwnerAddress()))
	if err != nil {
		log.Fatalf("Failed to get nonce: %v\n", err)
	}
	log.Printf("base nonce:%d ", nonce)

	start := time.Now()

	contractA, tx, contract, err := deployParentContract(copyAuth(auth), copyClient(conn), 1)
	if err == nil {
		log.Printf("address:0x%x\n", contractA)
		log.Printf("contract:%+v", contract)
		log.Printf("transaction:%v", tx.Hash().Hex())
		//mine deploy transaction
		dReceipt, _ := u.MineTransaction(copyClient(conn), tx)
		if dReceipt != nil {
			//upload data to contract
			auth.GasLimit = big.NewInt(2551488)
			upTx, err := contract.GenerateChild(auth, big.NewInt(2));
			if err == nil {
				//mine upload transaction
				uReceipt, err := u.MineTransaction(copyClient(conn), upTx)
				if uReceipt != nil {
					//TODO
					log.Println("transaction", upTx.String());
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

	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)

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

func deployParentContract(auth *bind.TransactOpts, conn *ethclient.Client, nonce uint64) (common.Address, *types.Transaction, *generator.Parent, error) {
	auth.GasLimit = big.NewInt(5500000)
	auth.GasPrice = big.NewInt(38000000000)
	address, tx, buyInfoC, err := generator.DeployParent(auth, conn);
	if err != nil {
		log.Printf("Failed to deploy new BuyInfo contract: %v\n", err)
		return u.GetStubAddress(), nil, nil, err
	}
	log.Printf("Contract pending deploy: 0x%x\n", address)
	log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	return address, tx, buyInfoC, nil
}
