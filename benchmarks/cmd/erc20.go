package cmd

import (
	"encoding/json"
	"ethbench/ethereum"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var erc20 = &cobra.Command{
	Use:   "erc20",
	Short: "benchmark erc20",
	Long:  `This command benchmarks an erc20 type of token`,
	Run: func(cmd *cobra.Command, args []string) {
		erc20FireItUp()
	},
}

func init() {
	RootCmd.AddCommand(erc20)
}

func erc20FireItUp() {

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
				time.Sleep(200 * time.Microsecond)
				go func(idx int) {
					balance, err := ethereum.GetWeiBalance(testAccounts[idx][0], client1)
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("balance wei %d", balance)

					//token balance
					tokenCaller, err := ethereum.NewErc20Caller(common.HexToAddress("0x3c1D955499887C35F4E6DFf3A0d8fA25983642F0"), client1)
					if err != nil {
						log.Fatal(err)
					}

					// gets the balance
					bal, err := tokenCaller.BalanceOf(nil, common.HexToAddress(testAccounts[idx][0]))
					if err != nil {
						log.Fatal(err)
					}
					fbalance := new(big.Float)
					fbalance.SetString(bal.String())
					decimalValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
					log.Printf("token balance %f", decimalValue)

				}(i)
			}
			time.Sleep(3 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			for i := 0; i < totalTestAccounts; i++ {
				time.Sleep(200 * time.Microsecond)
				go func(idx int) {
					balance, err := ethereum.GetWeiBalance(testAccounts[idx][0], client2)
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("balance wei %d", balance)

					//token balance
					tokenCaller, err := ethereum.NewErc20Caller(common.HexToAddress("0x3c1D955499887C35F4E6DFf3A0d8fA25983642F0"), client1)
					if err != nil {
						log.Fatal(err)
					}

					// gets the balance
					bal, err := tokenCaller.BalanceOf(nil, common.HexToAddress(testAccounts[idx][0]))
					if err != nil {
						log.Fatal(err)
					}
					fbalance := new(big.Float)
					fbalance.SetString(bal.String())
					decimalValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
					log.Printf("token balance %f", decimalValue)

				}(i)
			}
			time.Sleep(3 * time.Second)
		}
	}()

	time.Sleep(20 * time.Second)
}
