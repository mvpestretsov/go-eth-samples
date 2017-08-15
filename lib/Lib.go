package lib

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

var node_address string = "http://93.88.76.52:8545"
var key string = `{"address":"792832a7151cb334f3ff9f7880f68c676afc9b01","crypto":{"cipher":"aes-128-ctr","ciphertext":"f49a0fd038c3c1c5d65a307df2c6cedc7a04ac9a7ea8d964019dd0c857e517f7","cipherparams":{"iv":"2978e5e8f842afd8efc60ce9153848f1"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"52b861e95af7ca5eceb0ec02e3556872e2b905802f273de5397874dfbe1d116d"},"mac":"dc8b4ce35e732ebff234f1ef7476943772d503e93a8cf33f78586fbba244935b"},"id":"893275b6-1611-4962-a744-5fd0fb026278","version":3}`
var token_contract string = "0x920f0F1Db4fea63178e337FfB9f738607076a6C7"
var owner_address string = "0x792832a7151cb334f3ff9f7880f68c676afc9b01"

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

func MineTransaction(conn *ethclient.Client, tx *types.Transaction) error {
	ctx, stop := context.WithTimeout(context.Background(), 60000*time.Millisecond)
	defer stop()
	receipt, err := bind.WaitMined(ctx, conn, tx)
	if err != nil {
		log.Printf("Failed to get receipt: %v", err)
		return err
	}
	if receipt != nil {
		log.Printf("Receipt: %v", receipt)
		return fmt.Errorf("receipt is null")
	}
	return nil
}
