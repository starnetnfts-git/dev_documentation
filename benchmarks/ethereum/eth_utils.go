package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"math/big"
)

//GetWeiBalance the balance in wei for a given address
func GetWeiBalance(address string, client *ethclient.Client) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// GetBlockNumber gets the block number
func GetBlockNumber(client *ethclient.Client) (string, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return "", err
	}
	return header.Number.String(), nil
}

// IsSyncying returns false if it's not syncing
func IsSyncying(client *ethclient.Client) (bool, error) {
	sync, err := client.SyncProgress(context.Background())
	if err != nil {
		return false, err
	}
	if sync == nil {
		return false, nil
	}
	return true, nil
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, err := decimal.NewFromString(value.String())
	if err != nil {
		log.Error(err)
		return decimal.NewFromFloat(0)
	}
	result := num.Div(mul)

	return result
}
