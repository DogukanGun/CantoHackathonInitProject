package aave

import (
	aavePool "canto/ethereum/contracts/aave"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
)

func GetApyForAavePools(client *ethclient.Client) []PoolAPY {
	resultArr := make([]PoolAPY, len(TOKENS))
	aaveV3PoolInstance, err := aavePool.NewAavePoolCaller(common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"), client)
	if err != nil {
		fmt.Printf("Error occured: %v", err)
	}
	counter := 0
	for key := range TOKENS {
		data, err := aaveV3PoolInstance.GetReserveData(nil, common.HexToAddress(key))
		if err != nil {
			fmt.Printf("Error occured: %v", err)
			return nil
		}
		apy, apr := calculateAPYAndAPR(data.CurrentLiquidityRate)
		apyResult := PoolAPY{
			AssetName:          TOKENS[key],
			AssetAddress:       key,
			APY:                apy,
			APR:                apr,
			LendingPoolAddress: "0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2",
		}
		resultArr[counter] = apyResult
	}
	return resultArr
}

func calculateAPYAndAPR(currentLiquidityRate *big.Int) (float64, float64) {
	SecondsPerYear := 31536000
	base := big.NewInt(10)
	exponent := big.NewInt(27)
	ray := new(big.Int).Exp(base, exponent, nil)
	rayAsFloat := new(big.Float).SetInt(ray)
	rayFloat64, _ := rayAsFloat.Float64()
	currentLiquidityRateFloat := new(big.Float).SetInt(currentLiquidityRate)
	liquidity, _ := currentLiquidityRateFloat.Float64()
	depositAPR := liquidity / rayFloat64
	depositAPRForApy := (depositAPR / float64(SecondsPerYear)) + 1
	temp := math.Pow(depositAPRForApy, float64(SecondsPerYear))
	result := (temp - 1) * 100

	return result, depositAPR
}
