package cmd

import (
	"encoding/json"
	"ethbench/ethereum"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var erc721 = &cobra.Command{
	Use:   "erc721",
	Short: "benchmark erc721",
	Long:  `This command benchmarks an erc721 type of token`,
	Run: func(cmd *cobra.Command, args []string) {
		erc721FireItUp()
	},
}

func init() {
	RootCmd.AddCommand(erc721)
}

func erc721FireItUp() {

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

	client1, err := ethereum.InitEthClient(os.Getenv("RPC_URL1"))
	if err != nil {
		log.Fatal(err)
	}
	client2, err := ethereum.InitEthClient(os.Getenv("RPC_URL2"))
	if err != nil {
		log.Fatal(err)
	}

	totalTestAccounts := 1000 //len(testAccounts)-2
	go func() {
		for i := 0; i < 100000; i++ {
			for i := 0; i < totalTestAccounts; i++ {
				time.Sleep(999 * time.Microsecond)
				go func(idx int) {
					balance, err := ethereum.GetWeiBalance(testAccounts[idx][0], client1)
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("balance wei %d", balance)

					tokenContract := common.HexToAddress("0x5fca33Ed0eDFEea852ace8129d0816876bD6A3D4")
					instance, err := ethereum.NewErc721(tokenContract, client1)
					if err != nil {
						log.Fatal(err)
					}
					tokenIDBig := new(big.Int)
					tokenIDBig, ok := tokenIDBig.SetString(fmt.Sprintf("%d", randInt(1, 10000)), 10)
					if !ok {
						log.Fatal("can't convert this tokenID to bigInt")
					}
					exists, err := instance.Exists(&bind.CallOpts{}, tokenIDBig)
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("Token %d exists %t", tokenIDBig, exists)
				}(i)
			}
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			for i := totalTestAccounts; i > 0; i-- {
				time.Sleep(999 * time.Microsecond)
				go func(idx int) {
					balance, err := ethereum.GetWeiBalance(testAccounts[idx][0], client2)
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("balance wei %d", balance)

					tokenContract := common.HexToAddress("0x5fca33Ed0eDFEea852ace8129d0816876bD6A3D4")
					instance, err := ethereum.NewErc721(tokenContract, client2)
					if err != nil {
						log.Fatal(err)
					}
					tokenIDBig := new(big.Int)
					tokenIDBig, ok := tokenIDBig.SetString(fmt.Sprintf("%d", randInt(1, 10000)), 10)
					if !ok {
						log.Fatal("can't convert this tokenID to bigInt")
					}
					exists, err := instance.Exists(&bind.CallOpts{}, tokenIDBig)
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("Token %d exists %t", tokenIDBig, exists)
				}(i)
			}
			time.Sleep(1 * time.Second)
		}
	}()


	time.Sleep(20 * time.Second)
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
