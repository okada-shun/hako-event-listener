package hkfinance

import (
	"hako-event-logs/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

type HkfinanceTx struct {
	Config    *config.Config
	Ethclient *ethclient.Client
	Hkfinance *Hkfinance
}

// return eth client
func GetEthclient(config *config.Config) (*ethclient.Client, error) {
	client, err := ethclient.Dial(config.Ethereum.NetworkURL)
	if err != nil {
		return nil, err
	}
	return client, nil
}
