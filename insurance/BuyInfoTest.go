package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mvpestretsov/go-eth-samples/insurance/util"
	u "github.com/mvpestretsov/go-eth-samples/lib"
	"math/big"
	"strings"
	"sync"
	"time"
	//"context"
	"context"
	"fmt"
)

func main() {

	/*
		log.Println(auth.Nonce)
		ctx := context.TODO();
		nonce,err:=conn.NonceAt(ctx,common.HexToAddress(u.GetOwnerAddress()),nil)
		log.Println(nonce)
	*/
	// Create RPC connection to a remote node
	conn, err := ethclient.Dial(u.GetNodeAddress())
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v\n", err)
	}
	// read key
	auth, err := bind.NewTransactor(strings.NewReader(u.GetKey()), "qwe123")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v\n", err)
	}
	start := time.Now()
	//compute nonce
	nonce, err := conn.PendingNonceAt(context.TODO(), common.HexToAddress(u.GetOwnerAddress()))
	//nonce=100
	log.Printf("base nonce:%d ", nonce)
	pendingTransactionCount, err := conn.PendingTransactionCount(context.TODO())
	log.Printf("base pendingTransactionCount:%d ", pendingTransactionCount)

	var tryCount int = 5
	var wg sync.WaitGroup
	wg.Add(tryCount)

	for i := 0; i < tryCount; i++ {
		DeployAndSaveIteration(&wg, i, conn, auth, nonce)
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)

}

func deployBuyInfoContract(auth *bind.TransactOpts, conn bind.ContractBackend, i int) (common.Address, *types.Transaction, *insurance.BuyInfo, error) {
	address, tx, buyInfoC, err := insurance.DeployBuyInfo(auth, conn, common.HexToAddress(u.GetOwnerAddress()), "pk", fmt.Sprintf("%s%d", "dataH", i), "idH", u.GetStubAddress(), big.NewInt(1000), big.NewInt(1000))
	if err != nil {
		log.Printf("Failed to deploy new BuyInfo contract: %v\n", err)
		return u.GetStubAddress(), nil, nil, err
	}
	log.Printf("Contract pending deploy: 0x%x\n", address)
	log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	return address, tx, buyInfoC, nil
}

func DeployAndSaveIteration(wg *sync.WaitGroup, i int, conn *ethclient.Client, auth *bind.TransactOpts, nonce uint64) error {

	log.Printf("%+v\n", auth)
	//mining contract
	//auth.GasPrice=big.NewInt(int64(2000000+i*20000+1))
	auth.Nonce = big.NewInt(int64(nonce + uint64(i+1)))
	//log.Printf("%d: gasPrice:%d",i,auth.GasPrice)
	//log.Printf("%d: nonce :%d",i,auth.Nonce)
	address, tx, contract, err := deployBuyInfoContract(auth, conn, i)
	if err != nil {
		log.Printf("%d:Failed to create new contract: %v\n", i, err)
		wg.Done()
		return err
	}
	log.Printf("%d: address of new BuyInfo contract:0x%x\n", i, address)
	log.Printf("%d:tx nonce=%v", i, tx.Nonce())
	go func() {
		u.MineTransaction(conn, tx)
		wg.Done()
	}()
	log.Println(contract)
	//dataHash,err:=buyInfoC.DataHash(nil)
	//if err != nil {
	//	log.Printf("%d:Failed to retrieve dataHash: %v\n",i, err)
	//	return err
	//}
	//log.Printf("encData:%v\n", dataHash)

	//upload data to contract
	/*
		auth.GasPrice=big.NewInt(int64(2000000+i*200000+2))
		auth.Nonce= big.NewInt(int64(nonce+uint64(i*2+2)))
		tx, err =buyInfoC.UploadData(auth,fmt.Sprintf("%s%d","test",i))
		if err != nil {
			log.Printf("%d:Failed to upload data: %v\n", i,err)
			return err
		}
		log.Printf("%d:upload pending: 0x%x\n", i,tx.Hash())
		u.MineTransaction(conn,tx)
		//read data
		encData,err:=buyInfoC.EncryptedData(nil)
		if err != nil {
			log.Printf("%d:Failed to retrieve encData: %v\n",i, err)
			return err
		}
		log.Printf("encData:%v\n", encData)
	*/
	return nil
}

func copyAuth(s *bind.TransactOpts) *bind.TransactOpts {
	clone := *s   // This is where we essentially make a new struct
	return &clone // Return reference to a new struct
}
func copyClient(s *ethclient.Client) *ethclient.Client {
	clone := *s   // This is where we essentially make a new struct
	return &clone // Return reference to a new struct
}
