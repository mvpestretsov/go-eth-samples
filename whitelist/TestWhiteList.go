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

	//"math/big"

	"github.com/mvpestretsov/go-eth-samples/whitelist/contract"
)

func main() {

	key1 := `{"address":"d889637e948b737b335e8b4585d29b1132987b36","crypto":{"cipher":"aes-128-ctr","ciphertext":"ec50e452cac44aafbd678dcb9796c6a7ae7d1b36f04100c4a0fa6d9644dcae9b","cipherparams":{"iv":"b7c326a53d9be43280913b58ec5ae59c"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"c9db835948ce03276b10c74fe345f70decf0a0a78aae706ef411c2fc12d9e4bf"},"mac":"112d7257bb4e3ad09b442061d5e41e46aa469c88541d07be0a31355d38939887"},"id":"835484af-622e-4aaf-938f-e5d86d2aa526","version":3}`;
	key2 := `{"address":"eab4297609dbca23dc0cfffe654a35b580c05104","crypto":{"cipher":"aes-128-ctr","ciphertext":"5705d82cab12fe164514abb8e6f31539472d4824fb4806e491917302eb3292b8","cipherparams":{"iv":"92b70ab8bc534715a9f8f3810603b07c"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"4bb4d85a5b8c0744236e9288fbcffef122cedcdbd0fcd2029e4f95716fcb2806"},"mac":"ee4fc7091697880c793dc4c2491cf72ddb3cb8b34edaddfe2b9e89a55f676de4"},"id":"28a865d6-fae1-4ce3-b940-4fc101c10362","version":3}`;
	key3 := `{"address":"e12bfd6f94b9b6517366972bd1591a22fdf8dd45","crypto":{"cipher":"aes-128-ctr","ciphertext":"f869f4f855b0f5672997b225ea6dd493b7bf1c8850f6180ea3a9f262d96e0bbd","cipherparams":{"iv":"83b9c791534229ca576c137372f215bf"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"ed3faf35ec916e347d5bb97b121399f474cb9579e44aefb6b007e9b3249893bd"},"mac":"d65e079a7e380deb7aa0a108ef434f9c4177965e2e97afc07ed0caca43ab1e18"},"id":"da1ef131-81bf-4713-8978-aabfb9192e25","version":3}`;
	account1 := `d889637e948b737b335e8b4585d29b1132987b36`;
	account2 := `eab4297609dbca23dc0cfffe654a35b580c05104`;
	//account3:=`e12bfd6f94b9b6517366972bd1591a22fdf8dd45`;

	conn, err := ethclient.Dial(u.GetNodeAddress())
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v\n", err)
	}
	log.Printf("conn:%+v", conn)
	// read key of 1st admin
	auth1, err := bind.NewTransactor(strings.NewReader(key1), "qwe123")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v\n", err)
	}
	log.Printf("auth1:%+v", auth1)

	// read key of 2st admin
	auth2, err := bind.NewTransactor(strings.NewReader(key2), "qwe123")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v\n", err)
	}
	log.Printf("auth2:%+v", auth2)

	// read key of 2st admin
	auth3, err := bind.NewTransactor(strings.NewReader(key3), "qwe123")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v\n", err)
	}
	log.Printf("auth3:%+v", auth3)

	//start to measure execution time
	start := time.Now()

	//deployAdministration
	admins := []common.Address{common.HexToAddress(account1)}
	administrationData := u.AdministrationData{1, admins}
	contractA, tx, contract, err := u.DeployAdministrationContract(copyAuth(auth1), copyClient(conn), &administrationData)
	if err == nil {
		log.Printf("administration:address:0x%x\n", contractA)
		log.Printf("admininstration:contract:%+v", contract)
		log.Printf("admininstration:transaction:%v", tx.Hash().Hex())
		//mine deploy transaction
		dReceipt, _ := u.MineTransaction(copyClient(conn), tx)
		if dReceipt != nil {
			//print trustees for account1
			isAcc1Trustee, _ := contract.TrueTrustee(&bind.CallOpts{}, common.HexToAddress(account1));
			log.Printf("trustee :%v", isAcc1Trustee);
			//print trustees for account2
			isAcc2Trustee, _ := contract.TrueTrustee(&bind.CallOpts{}, common.HexToAddress(account2));
			log.Printf("trustee :%v", isAcc2Trustee);
			//add account2 to trustee
			addTrusteeTx, err := contract.AddTrustee(auth1, common.HexToAddress(account2))
			if err == nil {
				//mine addTrustee transaction
				uReceipt, _ := u.MineTransaction(copyClient(conn), addTrusteeTx)
				if uReceipt != nil {
					//get voting contract
					actualVoting, _ := contract.Actual(&bind.CallOpts{});
					votingAddress := actualVoting.Voting;
					log.Printf("votingAddress:%v", votingAddress.Hex())
					votingContract, _ := whitelist.NewVoting(votingAddress, conn);
					//accept voting
					acceptTx, err := votingContract.Accept(copyAuth(auth1));
					if err != nil {
						log.Printf("Failed to accept voting : %v\n", err)
					}
					aReceipt, _ := u.MineTransaction(copyClient(conn), acceptTx)
					if aReceipt != nil {
						//count result
						countTx, err := votingContract.CountResult(copyAuth(auth1));
						if err != nil {
							log.Printf("Failed to count voting : %v\n", err)
						}
						cReceipt, _ := u.MineTransaction(copyClient(conn), countTx)
						if (cReceipt != nil) {
							//check trustee for account2
							isAcc2Trustee, _ := contract.TrueTrustee(&bind.CallOpts{}, common.HexToAddress(account2));
							log.Printf("trustee2 :%v", isAcc2Trustee);
							if isAcc2Trustee {
								//create voting for add account2 for admins
								addAdmTx, err := contract.AddAdmin(auth1, common.HexToAddress(account2));
								if err != nil {
									log.Printf("Failed to create voting : %v\n", err)
								}
								addAdmR, _ := u.MineTransaction(copyClient(conn), addAdmTx)
								if addAdmR != nil {
									//get voting contract
									actualVoting, _ := contract.Actual(&bind.CallOpts{});
									votingAddress := actualVoting.Voting;
									log.Printf("votingAddress:%v", votingAddress.Hex())
									votingContract, _ := whitelist.NewVoting(votingAddress, conn);
									//accept voting
									acceptTx, err := votingContract.Accept(copyAuth(auth1));
									if err != nil {
										log.Printf("Failed to accept voting : %v\n", err)
									}
									aReceipt, _ := u.MineTransaction(copyClient(conn), acceptTx)
									if aReceipt != nil {
										//count result
										countTx, err := votingContract.CountResult(copyAuth(auth1));
										if err != nil {
											log.Printf("Failed to count voting : %v\n", err)
										}
										cReceipt, _ := u.MineTransaction(copyClient(conn), countTx)
										if (cReceipt != nil) {
											//check trustee for account2
											isAcc2Admin, _ := contract.IsAdmin(&bind.CallOpts{}, common.HexToAddress(account2));
											log.Printf("isAcc2Admin :%v", isAcc2Admin);
										}
									}
								}
							}
						}
					}
				}
			} else {
				log.Printf("Failed to addTrustee : %v\n", err)
			}

			//upload data to contract
			/*
					auth1.GasLimit = big.NewInt(2551488)
					upTx, err := contract.AddTrustee(aut2,common.HexToAddress(""))
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
					*/

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
