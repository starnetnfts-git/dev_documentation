package cmd

import (
	"encoding/json"
	"ethbench/ethereum"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var eth = &cobra.Command{
	Use:   "eth",
	Short: "ETH transactions benchmarks",
	Long:  `this command benchmarks eth transactions`,
	Run: func(cmd *cobra.Command, args []string) {
		fireItUp()
	},
}

func init() {
	RootCmd.AddCommand(eth)
}

func fireItUp() {

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

	client1, err := ethereum.InitEthClient(os.Getenv("RPC_URL1"))
	if err != nil {
		log.Fatal(err)
	}
	client1b, err := ethereum.InitEthClient(os.Getenv("RPC_URL2"))
	if err != nil {
		log.Fatal(err)
	}
	client2, err := ethereum.InitEthClient(os.Getenv("RPC_URL3"))
	if err != nil {
		log.Fatal(err)
	}
	client2b, err := ethereum.InitEthClient(os.Getenv("RPC_URL4"))
	if err != nil {
		log.Fatal(err)
	}


	go func() {
		for i := 0; i < 10; i++ {
			for i := 0; i < len(testAccounts)-2; i++ {
				go func(idx int) {
					fmt.Printf("starting index %d\n", idx)
					err := sendEthToAddress(client1, testAccounts[idx][0], "100000", testAccounts[idx+1][1])
					if err != nil {
						fmt.Println(err)
					}
				}(i)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			for i := 0; i < len(testAccounts)-2; i++ {
				go func(idx int) {
					fmt.Printf("starting index %d\n", idx)
					err := sendEthToAddress(client1b, testAccounts[idx][0], "100000", testAccounts[idx+1][1])
					if err != nil {
						fmt.Println(err)
					}
				}(i)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			for i := 0; i < len(testAccounts)-2; i++ {
				go func(idx int) {
					fmt.Printf("starting index %d\n", idx)
					err := sendEthToAddress(client2, testAccounts[idx][0], "100000", testAccounts[idx+1][1])
					if err != nil {
						fmt.Println(err)
					}
				}(i)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			for i := 0; i < len(testAccounts)-2; i++ {
				go func(idx int) {
					fmt.Printf("starting index %d\n", idx)
					err := sendEthToAddress(client2b, testAccounts[idx][0], "100000", testAccounts[idx+1][1])
					if err != nil {
						fmt.Println(err)
					}
				}(i)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(20*time.Second)
}
