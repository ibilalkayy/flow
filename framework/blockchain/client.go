package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	coingecko "github.com/superoo7/go-gecko/v3"
)

type Client struct {
	ethClient  *ethclient.Client
	privateKey *ecdsa.PrivateKey
}

func NewClient() (*Client, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Connect to Ethereum client
	alchemyURL := os.Getenv("ALCHEMY_API_URL")
	if alchemyURL == "" {
		return nil, fmt.Errorf("ALCHEMY_API_URL is not set in .env file")
	}

	ethClient, err := ethclient.Dial(alchemyURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	// Load private key
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}

	return &Client{ethClient: ethClient, privateKey: privateKey}, nil
}

func (c *Client) getExchangeRate() (float64, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	cg := coingecko.NewClient(httpClient)
	simpleSinglePrice, err := cg.SimpleSinglePrice("ethereum", "usd")
	if err != nil {
		return 0, fmt.Errorf("failed to get exchange rate: %v", err)
	}
	return float64(simpleSinglePrice.MarketPrice), nil
}

func (c *Client) ConvertUSDtoETH(usdAmount float64) (*big.Int, error) {
	rate, err := c.getExchangeRate()
	if err != nil {
		return nil, err
	}
	ethAmount := usdAmount / rate
	ethAmountWei := new(big.Float).Mul(big.NewFloat(ethAmount), big.NewFloat(1e18))
	ethAmountBigInt := new(big.Int)
	ethAmountWei.Int(ethAmountBigInt)
	return ethAmountBigInt, nil
}

func (c *Client) AllocateFunds(toAddress string, amount *big.Int) (string, error) {
	publicKey := c.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasLimit := uint64(21000)
	gasPrice, err := c.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	to := common.HexToAddress(toAddress)
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, nil)

	chainID, err := c.ethClient.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), c.privateKey)
	if err != nil {
		return "", err
	}

	err = c.ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	return signedTx.Hash().Hex(), nil
}

func SpendFunds(usdAmount float64, recipientAddress string) {
	client, err := NewClient()
	if err != nil {
		log.Fatalf("Failed to create blockchain client: %v", err)
	}

	ethAmount, err := client.ConvertUSDtoETH(usdAmount)
	if err != nil {
		log.Fatalf("Failed to convert USD to ETH: %v", err)
	}

	txHash, err := client.AllocateFunds(recipientAddress, ethAmount)
	if err != nil {
		log.Fatalf("Failed to allocate funds: %v", err)
	}
	fmt.Printf("Funds allocated with tx hash: %s\n", txHash)
}
