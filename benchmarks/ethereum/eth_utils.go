package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"math/big"
	"regexp"
)

// SignAndSendTx - all the info you need for a transaction of ETH is here
func SignAndSendTx(value *big.Int, toAddress string, senderPrivateKey string) (string, error) {
	fmt.Printf("Private key detected %s", senderPrivateKey)
	privateKey, err := crypto.HexToECDSA(senderPrivateKey)
	if err != nil {
		return "", fmt.Errorf("error reading from private key")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error getting publicKey")
	}

	gasLimit := uint64(55723) // in units

	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("failed to estimate gas price. setting it to 10 gwei!")
		gasPrice = big.NewInt(10000000000) // 10 gwei
	}
	log.Printf("Gas price suggested %s", gasPrice)

	nonce, err := Client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(*publicKeyECDSA))
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), value, gasLimit, gasPrice, nil)

	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	err = Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	return signedTx.Hash().String(), nil
}

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

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// CalcGasCost calculate gas cost given gas limit (units) and gas price (wei)
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
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

// ToDecimal wei to decimals
func WeiToEth(ivalue interface{}) decimal.Decimal {
	return ToDecimal(ivalue, 18)
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}
