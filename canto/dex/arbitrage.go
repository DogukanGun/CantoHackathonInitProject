package cantoDex

import (
	"canto/canto/contracts"
	"canto/canto/contracts/erc20"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
)

func CalculatePrices(client *ethclient.Client) []Pool {
	pools := make([]Pool, len(DexPools))
	for i, poolAddress := range DexPools {
		pool, err := contracts.NewContractsCaller(common.HexToAddress(poolAddress), client)
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
		token0Address, err := pool.Token0(nil)
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
		token0 := getTokenEquality(client, pool, token0Address)
		token1Address, err := pool.Token1(nil)
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
		token1 := getTokenEquality(client, pool, token1Address)
		pools[i] = Pool{
			PoolAddress: poolAddress,
			Token0:      token0,
			Token1:      token1,
		}
	}
	return pools
}

func getTokenEquality(client *ethclient.Client, pool *contracts.ContractsCaller, tokenAddress common.Address) (tokenResponse Token) {
	erc20Contract, err := erc20.NewErc20Caller(tokenAddress, client)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	symbol, err := erc20Contract.Symbol(nil)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	price := getPriceFromPool(pool, tokenAddress)
	tokenResponse = Token{
		Symbol:            symbol,
		Address:           tokenAddress.Hex(),
		PriceAsOtherToken: price,
	}
	return
}

func getPriceFromPool(pool *contracts.ContractsCaller, tokenAddress common.Address) int64 {
	price, err := pool.GetAmountOut(nil, big.NewInt(1), tokenAddress)
	if err != nil {
		fmt.Printf("Error is %v", err)
		os.Exit(0)
	}
	return price.Int64()
}
