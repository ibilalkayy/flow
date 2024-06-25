package blockchain

import (
	"crypto/ecdsa"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
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
