package cmd

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"ethbench/ethereum"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var simplest = &cobra.Command{
	Use:   "simplest",
	Short: "benchmark simplest",
	Long:  `This command benchmarks an simplest type of contract`,
	Run: func(cmd *cobra.Command, args []string) {
		simplestFireItUp()
	},
}

func init() {
	RootCmd.AddCommand(simplest)
}

func simplestFireItUp() {

	client1, err := ethereum.InitEthClient(os.Getenv("RPC_URL1"))
	if err != nil {
		log.Fatal(err)
	}
	client2, err := ethereum.InitEthClient(os.Getenv("RPC_URL2"))
	if err != nil {
		log.Fatal(err)
	}

	jsonFile, err := os.Open("./test_accounts.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	type TestAccounts [][]string

	var testAccounts TestAccounts
	err = json.Unmarshal(byteValue, &testAccounts)
	if err != nil {
		log.Fatal(err)
	}

	fundedAccounts := 1000

	go func() {
		for i := 0; i < fundedAccounts; i++ {
			time.Sleep(50 * time.Microsecond)
			go func(idx int) {

				tokenContract := common.HexToAddress("0x5F5C51DED5469c0f7e4B459C48E13C48aa2FDB63")
				instance, err := ethereum.NewSimplest(tokenContract, client1)
				if err != nil {
					log.Fatal(err)
				}
				tokenIDBig := new(big.Int)
				tokenIDBig, ok := tokenIDBig.SetString(fmt.Sprintf("%d", randInt(1, 10000)), 10)
				if !ok {
					log.Fatal("can't convert this tokenID to bigInt")
				}

				gasPrice := new(big.Int)
				gasPrice, _ = gasPrice.SetString("1000000000", 10)

				privateKey, err := crypto.HexToECDSA(testAccounts[idx+1][1][2:])
				if err != nil {
					log.Fatal(err)
				}
				publicKey := privateKey.Public()
				publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
				if !ok {
					log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
				}

				fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

				nonce, err := client1.PendingNonceAt(context.Background(), fromAddress)
				if err != nil {
					log.Fatal(err)
				}
				auth := bind.NewKeyedTransactor(privateKey)
				auth.Nonce = big.NewInt(int64(nonce))
				auth.Value = big.NewInt(0)    // in wei
				auth.GasLimit = uint64(22000) // in units
				auth.GasPrice = gasPrice

				tx, err := instance.SetVersion(auth, tokenIDBig)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("hash %s", tx.Hash().Hex())

			}(i)
			time.Sleep(1 * time.Millisecond)
		}
	}()

	//go func() {
	//	for i := fundedAccounts; i > 0; i-- {
	//		time.Sleep(50 * time.Millisecond)
	//		go func(idx int) {
	//
	//			tokenContract := common.HexToAddress("0x5F5C51DED5469c0f7e4B459C48E13C48aa2FDB63")
	//			instance, err := ethereum.NewSimplest(tokenContract, client2)
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//			tokenIDBig := new(big.Int)
	//			tokenIDBig, ok := tokenIDBig.SetString(fmt.Sprintf("%d", randInt(1, 10000)), 10)
	//			if !ok {
	//				log.Fatal("can't convert this tokenID to bigInt")
	//			}
	//
	//			gasPrice := new(big.Int)
	//			gasPrice, _ = gasPrice.SetString("1000000000", 10)
	//
	//			privateKey, err := crypto.HexToECDSA(testAccounts[idx+1][1][2:])
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//			publicKey := privateKey.Public()
	//			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//			if !ok {
	//				log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//			}
	//
	//			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//
	//			nonce, err := client2.PendingNonceAt(context.Background(), fromAddress)
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//			auth := bind.NewKeyedTransactor(privateKey)
	//			auth.Nonce = big.NewInt(int64(nonce))
	//			auth.Value = big.NewInt(0)    // in wei
	//			auth.GasLimit = uint64(22000) // in units
	//			auth.GasPrice = gasPrice
	//
	//			tx, err := instance.SetVersion(auth, tokenIDBig)
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//			log.Printf("(2) hash %s", tx.Hash().Hex())
	//
	//		}(i)
	//		time.Sleep(1 * time.Millisecond)
	//	}
	//}()

	// READS ...........
	go func() {
		for i := fundedAccounts; i > 0; i-- {
			time.Sleep(500 * time.Microsecond)
			go func(idx int) {
				tokenContract := common.HexToAddress("0x5F5C51DED5469c0f7e4B459C48E13C48aa2FDB63")
				instance, err := ethereum.NewSimplest(tokenContract, client1)
				if err != nil {
					log.Fatal(err)
				}
				version, err := instance.Version(nil)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("(1) got read version %d", version.Int64())

			}(i)
			time.Sleep(1 * time.Millisecond)
		}
	}()
	//
	go func() {
		for i := fundedAccounts; i > 0; i-- {
			time.Sleep(500 * time.Microsecond)
			go func(idx int) {
				tokenContract := common.HexToAddress("0x5F5C51DED5469c0f7e4B459C48E13C48aa2FDB63")
				instance, err := ethereum.NewSimplest(tokenContract, client2)
				if err != nil {
					log.Fatal(err)
				}
				version, err := instance.Version(nil)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("(2) got read version %d", version.Int64())

			}(i)
			time.Sleep(1 * time.Millisecond)
		}
	}()


	time.Sleep(20 * time.Second)
}
