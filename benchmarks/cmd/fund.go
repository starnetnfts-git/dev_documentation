package cmd

//normally a smart contract multisend would make sending MUCH faster....

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"ethbench/ethereum"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var fund = &cobra.Command{
	Use:   "fund",
	Short: "fund with ETH & ERC20 the accounts",
	Long:  `this command funds the test accounts with enough $`,
	Run: func(cmd *cobra.Command, args []string) {
		fundTestAddresses()
	},
}

var parentAddress string
var parentAddressPrivateKey string

func init() {
	RootCmd.AddCommand(fund)
}

func fundTestAddresses() {
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
	log.Printf("loaded %d test accounts", len(testAccounts))

	parentAddress = os.Getenv("PARENT_ADDRESS")
	parentAddressPrivateKey = os.Getenv("PARENT_PRIVATE_KEY")

	fmt.Printf("Parent Address %s\n", parentAddress)
	//fmt.Printf("Got %d test addresses\n", len(test_addresses))

	client1, err := ethereum.InitEthClient(os.Getenv("RPC_URL1"))
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
	if ethValue.String() == "0" {
		log.Fatal("please fund the parent account with lots of eth")
	}

	totalAccounts := 1000
	weiToSend := "1000000000000000" // 0.001

	for i := 0; i < totalAccounts; i++ {
		tAddress := testAccounts[i]

		fmt.Printf("Current Index %d / %d\n", i, len(testAccounts))
		err := sendEthToAddress(client1, tAddress[0], weiToSend, parentAddressPrivateKey) // 0.0001
		if err != nil {
			log.Fatal(err)
		}
	}

	// verifying some the balances
	for i := 0; i < 10; i = i + 10 {
		tAddress := testAccounts[i]
		balance, err := ethereum.GetWeiBalance(tAddress[0], ethereum.Client)
		if err != nil {
			log.Fatal(err)
		}
		fbalance := new(big.Float)
		fbalance.SetString(balance.String())
		ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
		fmt.Printf("Account %s has %s ETH\n", tAddress[0], ethValue.String())
		if ethValue.String() == "0" {
			fmt.Println("this should never happen but we got an address with 0 ETH")
		}
	}

}

func sendEthToAddress(client *ethclient.Client, toAddress string, amountInWei string, senderPrivateKey string) error {
	privateKey, err := crypto.HexToECDSA(senderPrivateKey[2:])
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("can't cast to ecds.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	value := new(big.Int)
	value, ok = value.SetString(amountInWei, 10)
	if !ok {
		return fmt.Errorf("wrong wei amount")
	}

	gasLimit := uint64(21000) // in units
	gasPriceWei, _ := strconv.Atoi(os.Getenv("GAS_PRICE_WEI"))
	gasPrice := big.NewInt(int64(gasPriceWei))

	var data []byte //nil
	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), value, gasLimit, gasPrice, data)

	chainID, _ := strconv.Atoi(os.Getenv("CHAIN_ID"))
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(chainID))), privateKey)
	if err != nil {
		return err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("failed to send eth: %s", err)
	}

	log.Printf("sent %s wei from %s to %s -> %s\n", amountInWei, fromAddress.String(), toAddress, tx.Hash().String())
	return nil
}
