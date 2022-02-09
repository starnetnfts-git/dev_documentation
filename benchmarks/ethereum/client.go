package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
)

// Client the eth client
var Client *ethclient.Client

// GasPrice stores the suggested gas price
var GasPrice string

// InitEthClient initialises the client
//noinspection GoNilness
func InitEthClient(url string) (*ethclient.Client, error) {

	ethClient, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	log.Printf("connected to the ETH provider %s", url)

	// is Syncing ?
	isSyncing, err := IsSyncying(ethClient)
	if err != nil {
		log.Warnf("can't get suggested gas price %s", err.Error())
	}
	if isSyncing {
		log.Error("!! The ETH client is Syncing !!")
	}

	blockNumber, err := GetBlockNumber(ethClient)
	if err != nil {
		log.Warnf("can't get suggested gas price %s", err.Error())
	}
	log.Printf("current block number %s", blockNumber)

	gPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Warnf("can't get suggested gas price %s", err.Error())
	}
	if gPrice == nil {
		gPrice, _ = new(big.Int).SetString("50000000000", 10)
	}

	wei := new(big.Int)
	wei.SetString(gPrice.String(), 10)
	GasPrice = ToDecimal(wei, 18).String()

	log.Printf("suggested gas price %s ETH", GasPrice)

	Client = ethClient

	return ethClient, nil
}
