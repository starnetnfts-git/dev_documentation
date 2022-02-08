package cmd

import (
	"encoding/json"
	"ethbench/ethereum"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"os"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "unspent eth/erc20 transfer to base account",
	Long:  `transfers all unspent eth/erc20 from test accounts to base account`,
	Run: func(cmd *cobra.Command, args []string) {
		CleanUp()
	},
}

func init() {
	RootCmd.AddCommand(cleanupCmd)
}

// Cleanup returns the unspent eth/erc20
func CleanUp() {

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

	parentAddress = os.Getenv("PARENT_ADDRESS")
	parentAddressPrivateKey = os.Getenv("PARENT_PRIVATE_KEY")

	fmt.Printf("Parent Address %s\n", parentAddress)

	client1, err := ethereum.InitEthClient(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatal(err)
	}

	balance, err := ethereum.GetWeiBalance(parentAddress, ethereum.Client)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Printf("Parent Account %s has %s ETH\n", parentAddress, ethValue.String())

	for k, tAddress := range testAccounts {
		// 1000000000000000000000 = 1000 ETH
		fmt.Printf("Current Index %d / %d\n", k, len(testAccounts))

		balance, err := ethereum.GetWeiBalance(tAddress[0], ethereum.Client)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Balance of account %s is %s\n", tAddress[0], balance.String())

		if balance.String() != "0" {

			err = sendEthToAddress(client1, parentAddress, balance.Sub(balance, big.NewInt(10000000000000000)).String(), tAddress[1])
			if err != nil {
				log.Println(err)
			}

		} else {
			fmt.Printf("account %s has 0 balance", tAddress[0])
		}
	}

	// verifying the balances
	for _, tAddress := range testAccounts {
		balance, err := ethereum.GetWeiBalance(tAddress[0], ethereum.Client)
		if err != nil {
			log.Fatal(err)
		}
		fbalance := new(big.Float)
		fbalance.SetString(balance.String())
		ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
		fmt.Printf("Account %s has %s ETH\n", tAddress[0], ethValue.String())
	}
}
