package lib

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"

	"github.com/mvpestretsov/go-eth-samples/insurance/util"
	"math/big"
)

var node_address string = "http://93.88.76.56:8545"
var key string = `{"address":"be70bc43cf77c53b25ac47cda34a7189e11f9dd5","crypto":{"cipher":"aes-128-ctr","ciphertext":"c234e54d6f7715b2d8b2cb83d904ab33dbc18860b7a2b7b4d99c815f74fd7170","cipherparams":{"iv":"696856fef88f8dc7144e15e4b89f1cd7"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"0a95535624fb402158f4f5b3fe3ec195d878b24a28f01042aeea210eda70e188"},"mac":"d19c874527186694034b1e90178d773fcec6645714f0dbd983bc7a5ee3b5d4a6"},"id":"dc9d65fe-e3e0-4114-b102-539c9bd7643f","version":3}`
var token_contract string = "0x920f0F1Db4fea63178e337FfB9f738607076a6C7"
var owner_address string = "0xbe70bc43cf77c53b25ac47cda34a7189e11f9dd5"

func GetNodeAddress() string {
	return node_address
}

func GetKey() string {
	return key
}

func GetTokenContract() string {
	return token_contract
}

func GetOwnerAddress() string {
	return owner_address
}

func GetStubAddress() common.Address {
	return common.HexToAddress("0x0000000000000000000000000000000000000000")
}

func MineTransaction(conn *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	ctx, stop := context.WithTimeout(context.Background(), 30000*time.Millisecond)
	defer stop()
	receipt, err := bind.WaitMined(ctx, conn, tx)
	if err != nil {
		log.Printf("Failed to get receipt for transaction %+v: %v", tx, err)
	}
	if receipt != nil {
		log.Printf("Receipt: %v", receipt)
	} else {
		log.Printf("Receipt is null")
	}
	return receipt, err
}

/*func DeployContract(conn *ethclient.Client, tx *types.Transaction) error {
	ctx, stop := context.WithTimeout(context.Background(), 60000*time.Millisecond)
	defer stop()
	address, err := bind.WaitDeployed(ctx, conn, tx)
	if err != nil {
		log.Printf("Failed to deployed: %v", err)
		return err
	}
	log.Printf("Address: %v", address)

	return nil
}*/

func DeployBuyInfoContract(auth *bind.TransactOpts, conn *ethclient.Client, tran *BuyInfoContract) (common.Address, *types.Transaction, *insurance.BuyInfo, error) {
	auth.Nonce = big.NewInt(int64(tran.Nonce))
	auth.GasLimit = big.NewInt(5500000)
	auth.GasPrice = big.NewInt(38000000000)
	address, tx, buyInfoC, err := insurance.DeployBuyInfo(auth, conn, common.HexToAddress(GetOwnerAddress()), "pk", tran.DataH, "idH", GetStubAddress(), big.NewInt(1000), big.NewInt(1000))
	if err != nil {
		log.Printf("Failed to deploy new BuyInfo contract: %v\n", err)
		return GetStubAddress(), nil, nil, err
	}
	log.Printf("Contract pending deploy: 0x%x\n", address)
	log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	return address, tx, buyInfoC, nil
}

func GetActualNonce(client *ethclient.Client, address common.Address) (uint64, error) {
	ctx, stop := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer stop()
	nonce, err := client.NonceAt(ctx, address, nil)
	return nonce, err
}

type BuyInfoContract struct {
	Nonce uint64
	DataH string
}
