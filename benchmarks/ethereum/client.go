package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

// Client the eth client
var Client *ethclient.Client

// InitEthClient initialises the the client
//noinspection GoNilness
func InitEthClient(url string) (*ethclient.Client, error) {

	ethClient, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	log.Printf("connected to the ETH provider %s", url)

	Client = ethClient

	return ethClient, nil
}
