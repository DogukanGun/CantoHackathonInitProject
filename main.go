package main

import (
	cantoDex "canto/canto/dex"
	"canto/canto/lending"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("RPC"))
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	dexList := cantoDex.CalculatePrices(client)
	for i, pool := range dexList {
		fmt.Printf("Index %d the pool is %v \n", i, pool)
	}
	lendingList := lending.GetLendingPoolsApy(client)
	for i, pool := range lendingList {
		fmt.Printf("Index %d the pool is %v \n", i, pool)
	}
}
